package orchestrator

import (
	"context"
	"errors"
	"fmt"
	"time"

	tcfg "github.com/RTradeLtd/config"
	"github.com/RTradeLtd/database"
	"github.com/RTradeLtd/database/models"

	"github.com/RTradeLtd/ipfs-orchestrator/config"
	"github.com/RTradeLtd/ipfs-orchestrator/ipfs"
	"github.com/RTradeLtd/ipfs-orchestrator/log"
	"github.com/RTradeLtd/ipfs-orchestrator/registry"
	"go.uber.org/zap"
)

// Orchestrator contains most primary application logic and manages node
// availability
type Orchestrator struct {
	l  *zap.SugaredLogger
	nm *models.IPFSNetworkManager

	client  ipfs.NodeClient
	reg     *registry.NodeRegistry
	address string
}

// New instantiates and bootstraps a new Orchestrator
func New(logger *zap.SugaredLogger, address string, c ipfs.NodeClient,
	ports config.Ports, pg tcfg.Database, dev bool) (*Orchestrator, error) {
	var l = logger.Named("orchestrator")
	if address == "" {
		l.Warn("host address not set")
	}

	// bootstrap registry
	l.Info("bootstrapping existing nodes")
	nodes, err := c.Nodes(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to fetch nodes: %s", err.Error())
	}
	reg := registry.New(l, ports, nodes...)

	// set up database connection
	l.Infow("intializing database connection",
		"db.host", pg.URL,
		"db.port", pg.Port,
		"db.name", pg.Name,
		"db.with_ssl", !dev,
		"db.with_migrations", dev)
	dbm, err := database.Initialize(&tcfg.TemporalConfig{
		Database: pg,
	}, database.Options{
		SSLModeDisable: dev,
		RunMigrations:  dev,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %s", err.Error())
	}
	l.Info("successfully connected to database")

	return &Orchestrator{
		l:       l,
		nm:      models.NewHostedIPFSNetworkManager(dbm.DB),
		client:  c,
		reg:     reg,
		address: address,
	}, nil
}

// Run initializes the orchestrator's background tasks. Cancelling the context
// will end the tasks and release the orchestrator's resources.
func (o *Orchestrator) Run(ctx context.Context) error {
	go o.client.Watch(ctx)
	go func() {
		select {
		case <-ctx.Done():
			o.l.Info("releasing orchestrator resources")

			// close registry
			o.reg.Close()

			// close database
			if err := o.nm.DB.Close(); err != nil {
				o.l.Warnw("error occurred closing database connection",
					"error", err)
			}
		}
	}()
	return nil
}

// NetworkDetails provides information about an instantiated network
type NetworkDetails struct {
	API      string
	SwarmKey string
}

// NetworkUp intializes a node for given network
func (o *Orchestrator) NetworkUp(ctx context.Context, network string) (NetworkDetails, error) {
	if network == "" {
		return NetworkDetails{}, errors.New("invalid network name provided")
	}

	var start = time.Now()
	var jobID = generateID()
	var l = log.NewProcessLogger(o.l, "network_up",
		"job_id", jobID,
		"network", network)
	l.Info("network up process started")

	// check if request is valid
	n, err := o.nm.GetNetworkByName(network)
	if err != nil {
		l.Infow("failed to fetch network from database",
			"error", err)
		return NetworkDetails{}, fmt.Errorf("no network with name '%s' found", network)
	}
	l = l.With("network.db_id", n.ID)
	l.Info("network retrieved from database")

	// set options based on database entry
	opts, err := getOptionsFromDatabaseEntry(n)
	if err != nil {
		l.Warnw("invalid database entry",
			"error", err)
		return NetworkDetails{}, fmt.Errorf("failed to configure network: %s", err.Error())
	}

	// register node for network
	newNode := getNodeFromDatabaseEntry(jobID, n)
	if err := o.reg.Register(newNode); err != nil {
		l.Errorw("no available ports",
			"error", err)
		return NetworkDetails{}, fmt.Errorf("failed to allocate resources for network '%s': %s", network, err)
	}

	// instantiate node
	l = l.With("node", newNode)
	l.Info("network registered, creating node")
	if err := o.client.CreateNode(ctx, newNode, opts); err != nil {
		l.Errorw("unable to create node",
			"error", err)
		return NetworkDetails{}, fmt.Errorf("failed to instantiate node for network '%s': %s", network, err)
	}
	l.Info("node created")

	// update network in database
	n.APIURL = o.address + ":" + newNode.Ports.API
	n.SwarmKey = string(opts.SwarmKey)
	n.Activated = time.Now()
	if check := o.nm.DB.Save(n); check != nil && check.Error != nil {
		l.Errorw("failed to update database",
			"error", err,
			"entry", n)
		return NetworkDetails{}, fmt.Errorf("failed to update network '%s': %s", network, check.Error)
	}

	l.Infow("network up process completed",
		"network_up.duration", time.Since(start))

	return NetworkDetails{
		API:      n.APIURL,
		SwarmKey: n.SwarmKey,
	}, nil
}

// NetworkUpdate updates given network's configuration from database
func (o *Orchestrator) NetworkUpdate(ctx context.Context, network string) error {
	if network == "" {
		return errors.New("invalid network name provided")
	}

	// check node exists
	node, err := o.reg.Get(network)
	if err != nil {
		return fmt.Errorf("failed to find node for network '%s': %s", network, err.Error())
	}

	var start = time.Now()
	var jobID = generateID()
	var l = log.NewProcessLogger(o.l, "network_update",
		"job_id", jobID,
		"network", network)
	l.Info("network up process started")

	// retrieve from database
	n, err := o.nm.GetNetworkByName(network)
	if err != nil {
		l.Infow("failed to fetch network from database even though node was found in registry",
			"node", node,
			"error", err)
		return fmt.Errorf("no network with name '%s' found", network)
	}
	l = l.With("network.db_id", n.ID)
	l.Info("network retrieved from database")

	// construct new node based on new config and old settings
	var new = getNodeFromDatabaseEntry(jobID, n)
	new.DockerID = node.DockerID
	new.Ports = node.Ports
	new.DataDir = node.DataDir

	// execute update
	l.Info("updating node",
		"node.config", new)
	if err = o.client.UpdateNode(ctx, new); err != nil {
		l.Errorw("failed to update network", "error", err)
		return fmt.Errorf("failed to update network '%s': %s", network, err.Error())
	}

	// update registry
	l.Info("updating registry")
	o.reg.Deregister(network)
	if err := o.reg.Register(new); err != nil {
		l.Errorw("failed to register updated network", "error", err)
		return fmt.Errorf("error updating registry: %s", err.Error())
	}

	l.Infow("network update process completed",
		"network_update.duration", time.Since(start))
	return nil
}

// NetworkDown brings a network offline
func (o *Orchestrator) NetworkDown(ctx context.Context, network string) error {
	if network == "" {
		return errors.New("invalid network name provided")
	}

	start := time.Now()
	jobID := generateID()
	l := log.NewProcessLogger(o.l, "network_down",
		"job_id", jobID,
		"network", network)
	l.Info("network up process started")

	// retrieve node from registry
	node, err := o.reg.Get(network)
	if err != nil {
		l.Info("could not find node in registry")
		return fmt.Errorf("failed to get node for network %s from registry: %s", network, err.Error())
	}

	// shut down node
	l = l.With("node", node)
	l.Info("network found, stopping node")
	if err := o.client.StopNode(ctx, &node); err != nil {
		l.Errorw("error occurred while stopping node",
			"error", err)
	}
	l.Info("node stopped")

	// deregister node
	if err := o.reg.Deregister(network); err != nil {
		l.Errorw("error occurred while deregistering node",
			"error", err)
	}

	// update network in database to indicate it is no longer active
	var t time.Time
	if err := o.nm.UpdateNetworkByName(network, map[string]interface{}{
		"activated": t,
		"api_url":   "",
	}); err != nil {
		l.Errorw("failed to update database entry for network",
			"err", err)
		return fmt.Errorf("failed to update network '%s': %s", network, err)
	}

	l.Infow("network down process completed",
		"network_down.duration", time.Since(start))

	return nil
}

// NetworkRemove removes network assets
func (o *Orchestrator) NetworkRemove(ctx context.Context, network string) error {
	if network == "" {
		return errors.New("invalid network name provided")
	}

	if _, err := o.reg.Get(network); err == nil {
		return errors.New("network is still online and in registry - must be offline for removal")
	}

	return o.client.RemoveNode(ctx, network)
}

// NetworkStatus denotes details about requested network
type NetworkStatus struct {
	Network   string
	API       string
	Uptime    time.Duration
	DiskUsage int64
	Stats     interface{}
}

// NetworkStatus retrieves the status of the node for the given status
func (o *Orchestrator) NetworkStatus(ctx context.Context, network string) (NetworkStatus, error) {
	n, err := o.reg.Get(network)
	if err != nil {
		return NetworkStatus{}, fmt.Errorf("failed to retrieve network details: %s", err.Error())
	}

	stats, err := o.client.NodeStats(ctx, &n)
	if err != nil {
		o.l.Errorw("error occurred while attempting to acess registered node",
			"error", err,
			"node", n)
		return NetworkStatus{}, err
	}

	return NetworkStatus{
		Network:   network,
		API:       o.address + ":" + n.Ports.API,
		Uptime:    stats.Uptime,
		DiskUsage: stats.DiskUsage,
		Stats:     stats.Stats,
	}, nil
}
