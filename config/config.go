package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// DefaultIPFSVersion declares the current default version of go-ipfs to use
const DefaultIPFSVersion = "v0.4.17"

// IPFSOrchestratorConfig configures the orchestration daemon
type IPFSOrchestratorConfig struct {
	IPFS     `json:"ipfs"`
	API      `json:"api"`
	Postgres `json:"postgres"`
}

// IPFS configures settings relevant to IPFS nodes
type IPFS struct {
	Version string `json:"version"`
	Ports   `json:"ports"`
}

// Ports declares port-range configuration for IPFS nodes. Elements of each
// array can be of the form "<PORT>" or "<LOWER>-<UPPER>"
type Ports struct {
	Swarm   []string `json:"swarm"`
	API     []string `json:"api"`
	Gateway []string `json:"gateway"`
}

// API declares configuration for the orchestrator daemon's gRPC API
type API struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	KeyPath string `json:"keypath"`
}

// Postgres declares configuration for the orchestrator's database
type Postgres struct {
	Database string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoadConfig loads a TemporalConfig from given filepath
func LoadConfig(configPath string) (IPFSOrchestratorConfig, error) {
	var cfg IPFSOrchestratorConfig

	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return cfg, fmt.Errorf("could not open config: %s", err.Error())
	}

	if err = json.Unmarshal(raw, &cfg); err != nil {
		return cfg, fmt.Errorf("could not read config: %s", err.Error())
	}

	cfg.setDefaults()

	return cfg, nil
}

// GenerateConfig writes a empty TemporalConfig template to given filepath
func GenerateConfig(configPath string) error {
	template := &IPFSOrchestratorConfig{}
	template.setDefaults()
	b, err := json.Marshal(template)
	if err != nil {
		return err
	}

	var pretty bytes.Buffer
	if err = json.Indent(&pretty, b, "", "\t"); err != nil {
		return err
	}
	return ioutil.WriteFile(configPath, pretty.Bytes(), os.ModePerm)
}

func (c *IPFSOrchestratorConfig) setDefaults() {
	if c.IPFS.Version == "" {
		c.IPFS.Version = DefaultIPFSVersion
	}
	if c.API.Host == "" {
		c.API.Host = "127.0.0.1"
	}
	if c.API.Port == "" {
		c.API.Port = "9111"
	}
	if c.Postgres.Host == "" {
		c.Postgres.Host = "127.0.0.1"
	}
	if c.Postgres.Port == "" {
		c.Postgres.Port = "5432"
	}
	if c.IPFS.Ports.Swarm == nil {
		c.IPFS.Ports.Swarm = []string{"4001-5000"}
	}
	if c.IPFS.Ports.API == nil {
		c.IPFS.Ports.API = []string{"5001-6000"}
	}
	if c.IPFS.Ports.Gateway == nil {
		c.IPFS.Ports.Gateway = []string{"8001-9000"}
	}
}
