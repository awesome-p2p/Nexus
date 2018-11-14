// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	context "context"
	sync "sync"

	ipfs "github.com/RTradeLtd/ipfs-orchestrator/ipfs"
)

type FakeNodeClient struct {
	CreateNodeStub        func(context.Context, *ipfs.NodeInfo, ipfs.NodeOpts) error
	createNodeMutex       sync.RWMutex
	createNodeArgsForCall []struct {
		arg1 context.Context
		arg2 *ipfs.NodeInfo
		arg3 ipfs.NodeOpts
	}
	createNodeReturns struct {
		result1 error
	}
	createNodeReturnsOnCall map[int]struct {
		result1 error
	}
	NodeStatsStub        func(context.Context, *ipfs.NodeInfo) (ipfs.NodeStats, error)
	nodeStatsMutex       sync.RWMutex
	nodeStatsArgsForCall []struct {
		arg1 context.Context
		arg2 *ipfs.NodeInfo
	}
	nodeStatsReturns struct {
		result1 ipfs.NodeStats
		result2 error
	}
	nodeStatsReturnsOnCall map[int]struct {
		result1 ipfs.NodeStats
		result2 error
	}
	NodesStub        func(context.Context) ([]*ipfs.NodeInfo, error)
	nodesMutex       sync.RWMutex
	nodesArgsForCall []struct {
		arg1 context.Context
	}
	nodesReturns struct {
		result1 []*ipfs.NodeInfo
		result2 error
	}
	nodesReturnsOnCall map[int]struct {
		result1 []*ipfs.NodeInfo
		result2 error
	}
	RemoveNodeStub        func(context.Context, string) error
	removeNodeMutex       sync.RWMutex
	removeNodeArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	removeNodeReturns struct {
		result1 error
	}
	removeNodeReturnsOnCall map[int]struct {
		result1 error
	}
	StopNodeStub        func(context.Context, *ipfs.NodeInfo) error
	stopNodeMutex       sync.RWMutex
	stopNodeArgsForCall []struct {
		arg1 context.Context
		arg2 *ipfs.NodeInfo
	}
	stopNodeReturns struct {
		result1 error
	}
	stopNodeReturnsOnCall map[int]struct {
		result1 error
	}
	WatchStub        func(context.Context) (<-chan ipfs.Event, <-chan error)
	watchMutex       sync.RWMutex
	watchArgsForCall []struct {
		arg1 context.Context
	}
	watchReturns struct {
		result1 <-chan ipfs.Event
		result2 <-chan error
	}
	watchReturnsOnCall map[int]struct {
		result1 <-chan ipfs.Event
		result2 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNodeClient) CreateNode(arg1 context.Context, arg2 *ipfs.NodeInfo, arg3 ipfs.NodeOpts) error {
	fake.createNodeMutex.Lock()
	ret, specificReturn := fake.createNodeReturnsOnCall[len(fake.createNodeArgsForCall)]
	fake.createNodeArgsForCall = append(fake.createNodeArgsForCall, struct {
		arg1 context.Context
		arg2 *ipfs.NodeInfo
		arg3 ipfs.NodeOpts
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateNode", []interface{}{arg1, arg2, arg3})
	fake.createNodeMutex.Unlock()
	if fake.CreateNodeStub != nil {
		return fake.CreateNodeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createNodeReturns
	return fakeReturns.result1
}

func (fake *FakeNodeClient) CreateNodeCallCount() int {
	fake.createNodeMutex.RLock()
	defer fake.createNodeMutex.RUnlock()
	return len(fake.createNodeArgsForCall)
}

func (fake *FakeNodeClient) CreateNodeCalls(stub func(context.Context, *ipfs.NodeInfo, ipfs.NodeOpts) error) {
	fake.createNodeMutex.Lock()
	defer fake.createNodeMutex.Unlock()
	fake.CreateNodeStub = stub
}

func (fake *FakeNodeClient) CreateNodeArgsForCall(i int) (context.Context, *ipfs.NodeInfo, ipfs.NodeOpts) {
	fake.createNodeMutex.RLock()
	defer fake.createNodeMutex.RUnlock()
	argsForCall := fake.createNodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeNodeClient) CreateNodeReturns(result1 error) {
	fake.createNodeMutex.Lock()
	defer fake.createNodeMutex.Unlock()
	fake.CreateNodeStub = nil
	fake.createNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeClient) CreateNodeReturnsOnCall(i int, result1 error) {
	fake.createNodeMutex.Lock()
	defer fake.createNodeMutex.Unlock()
	fake.CreateNodeStub = nil
	if fake.createNodeReturnsOnCall == nil {
		fake.createNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeClient) NodeStats(arg1 context.Context, arg2 *ipfs.NodeInfo) (ipfs.NodeStats, error) {
	fake.nodeStatsMutex.Lock()
	ret, specificReturn := fake.nodeStatsReturnsOnCall[len(fake.nodeStatsArgsForCall)]
	fake.nodeStatsArgsForCall = append(fake.nodeStatsArgsForCall, struct {
		arg1 context.Context
		arg2 *ipfs.NodeInfo
	}{arg1, arg2})
	fake.recordInvocation("NodeStats", []interface{}{arg1, arg2})
	fake.nodeStatsMutex.Unlock()
	if fake.NodeStatsStub != nil {
		return fake.NodeStatsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.nodeStatsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNodeClient) NodeStatsCallCount() int {
	fake.nodeStatsMutex.RLock()
	defer fake.nodeStatsMutex.RUnlock()
	return len(fake.nodeStatsArgsForCall)
}

func (fake *FakeNodeClient) NodeStatsCalls(stub func(context.Context, *ipfs.NodeInfo) (ipfs.NodeStats, error)) {
	fake.nodeStatsMutex.Lock()
	defer fake.nodeStatsMutex.Unlock()
	fake.NodeStatsStub = stub
}

func (fake *FakeNodeClient) NodeStatsArgsForCall(i int) (context.Context, *ipfs.NodeInfo) {
	fake.nodeStatsMutex.RLock()
	defer fake.nodeStatsMutex.RUnlock()
	argsForCall := fake.nodeStatsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNodeClient) NodeStatsReturns(result1 ipfs.NodeStats, result2 error) {
	fake.nodeStatsMutex.Lock()
	defer fake.nodeStatsMutex.Unlock()
	fake.NodeStatsStub = nil
	fake.nodeStatsReturns = struct {
		result1 ipfs.NodeStats
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) NodeStatsReturnsOnCall(i int, result1 ipfs.NodeStats, result2 error) {
	fake.nodeStatsMutex.Lock()
	defer fake.nodeStatsMutex.Unlock()
	fake.NodeStatsStub = nil
	if fake.nodeStatsReturnsOnCall == nil {
		fake.nodeStatsReturnsOnCall = make(map[int]struct {
			result1 ipfs.NodeStats
			result2 error
		})
	}
	fake.nodeStatsReturnsOnCall[i] = struct {
		result1 ipfs.NodeStats
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) Nodes(arg1 context.Context) ([]*ipfs.NodeInfo, error) {
	fake.nodesMutex.Lock()
	ret, specificReturn := fake.nodesReturnsOnCall[len(fake.nodesArgsForCall)]
	fake.nodesArgsForCall = append(fake.nodesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Nodes", []interface{}{arg1})
	fake.nodesMutex.Unlock()
	if fake.NodesStub != nil {
		return fake.NodesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.nodesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNodeClient) NodesCallCount() int {
	fake.nodesMutex.RLock()
	defer fake.nodesMutex.RUnlock()
	return len(fake.nodesArgsForCall)
}

func (fake *FakeNodeClient) NodesCalls(stub func(context.Context) ([]*ipfs.NodeInfo, error)) {
	fake.nodesMutex.Lock()
	defer fake.nodesMutex.Unlock()
	fake.NodesStub = stub
}

func (fake *FakeNodeClient) NodesArgsForCall(i int) context.Context {
	fake.nodesMutex.RLock()
	defer fake.nodesMutex.RUnlock()
	argsForCall := fake.nodesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNodeClient) NodesReturns(result1 []*ipfs.NodeInfo, result2 error) {
	fake.nodesMutex.Lock()
	defer fake.nodesMutex.Unlock()
	fake.NodesStub = nil
	fake.nodesReturns = struct {
		result1 []*ipfs.NodeInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) NodesReturnsOnCall(i int, result1 []*ipfs.NodeInfo, result2 error) {
	fake.nodesMutex.Lock()
	defer fake.nodesMutex.Unlock()
	fake.NodesStub = nil
	if fake.nodesReturnsOnCall == nil {
		fake.nodesReturnsOnCall = make(map[int]struct {
			result1 []*ipfs.NodeInfo
			result2 error
		})
	}
	fake.nodesReturnsOnCall[i] = struct {
		result1 []*ipfs.NodeInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) RemoveNode(arg1 context.Context, arg2 string) error {
	fake.removeNodeMutex.Lock()
	ret, specificReturn := fake.removeNodeReturnsOnCall[len(fake.removeNodeArgsForCall)]
	fake.removeNodeArgsForCall = append(fake.removeNodeArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("RemoveNode", []interface{}{arg1, arg2})
	fake.removeNodeMutex.Unlock()
	if fake.RemoveNodeStub != nil {
		return fake.RemoveNodeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeNodeReturns
	return fakeReturns.result1
}

func (fake *FakeNodeClient) RemoveNodeCallCount() int {
	fake.removeNodeMutex.RLock()
	defer fake.removeNodeMutex.RUnlock()
	return len(fake.removeNodeArgsForCall)
}

func (fake *FakeNodeClient) RemoveNodeCalls(stub func(context.Context, string) error) {
	fake.removeNodeMutex.Lock()
	defer fake.removeNodeMutex.Unlock()
	fake.RemoveNodeStub = stub
}

func (fake *FakeNodeClient) RemoveNodeArgsForCall(i int) (context.Context, string) {
	fake.removeNodeMutex.RLock()
	defer fake.removeNodeMutex.RUnlock()
	argsForCall := fake.removeNodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNodeClient) RemoveNodeReturns(result1 error) {
	fake.removeNodeMutex.Lock()
	defer fake.removeNodeMutex.Unlock()
	fake.RemoveNodeStub = nil
	fake.removeNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeClient) RemoveNodeReturnsOnCall(i int, result1 error) {
	fake.removeNodeMutex.Lock()
	defer fake.removeNodeMutex.Unlock()
	fake.RemoveNodeStub = nil
	if fake.removeNodeReturnsOnCall == nil {
		fake.removeNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeClient) StopNode(arg1 context.Context, arg2 *ipfs.NodeInfo) error {
	fake.stopNodeMutex.Lock()
	ret, specificReturn := fake.stopNodeReturnsOnCall[len(fake.stopNodeArgsForCall)]
	fake.stopNodeArgsForCall = append(fake.stopNodeArgsForCall, struct {
		arg1 context.Context
		arg2 *ipfs.NodeInfo
	}{arg1, arg2})
	fake.recordInvocation("StopNode", []interface{}{arg1, arg2})
	fake.stopNodeMutex.Unlock()
	if fake.StopNodeStub != nil {
		return fake.StopNodeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stopNodeReturns
	return fakeReturns.result1
}

func (fake *FakeNodeClient) StopNodeCallCount() int {
	fake.stopNodeMutex.RLock()
	defer fake.stopNodeMutex.RUnlock()
	return len(fake.stopNodeArgsForCall)
}

func (fake *FakeNodeClient) StopNodeCalls(stub func(context.Context, *ipfs.NodeInfo) error) {
	fake.stopNodeMutex.Lock()
	defer fake.stopNodeMutex.Unlock()
	fake.StopNodeStub = stub
}

func (fake *FakeNodeClient) StopNodeArgsForCall(i int) (context.Context, *ipfs.NodeInfo) {
	fake.stopNodeMutex.RLock()
	defer fake.stopNodeMutex.RUnlock()
	argsForCall := fake.stopNodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNodeClient) StopNodeReturns(result1 error) {
	fake.stopNodeMutex.Lock()
	defer fake.stopNodeMutex.Unlock()
	fake.StopNodeStub = nil
	fake.stopNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeClient) StopNodeReturnsOnCall(i int, result1 error) {
	fake.stopNodeMutex.Lock()
	defer fake.stopNodeMutex.Unlock()
	fake.StopNodeStub = nil
	if fake.stopNodeReturnsOnCall == nil {
		fake.stopNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeClient) Watch(arg1 context.Context) (<-chan ipfs.Event, <-chan error) {
	fake.watchMutex.Lock()
	ret, specificReturn := fake.watchReturnsOnCall[len(fake.watchArgsForCall)]
	fake.watchArgsForCall = append(fake.watchArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Watch", []interface{}{arg1})
	fake.watchMutex.Unlock()
	if fake.WatchStub != nil {
		return fake.WatchStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.watchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNodeClient) WatchCallCount() int {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	return len(fake.watchArgsForCall)
}

func (fake *FakeNodeClient) WatchCalls(stub func(context.Context) (<-chan ipfs.Event, <-chan error)) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = stub
}

func (fake *FakeNodeClient) WatchArgsForCall(i int) context.Context {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	argsForCall := fake.watchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNodeClient) WatchReturns(result1 <-chan ipfs.Event, result2 <-chan error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	fake.watchReturns = struct {
		result1 <-chan ipfs.Event
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeNodeClient) WatchReturnsOnCall(i int, result1 <-chan ipfs.Event, result2 <-chan error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	if fake.watchReturnsOnCall == nil {
		fake.watchReturnsOnCall = make(map[int]struct {
			result1 <-chan ipfs.Event
			result2 <-chan error
		})
	}
	fake.watchReturnsOnCall[i] = struct {
		result1 <-chan ipfs.Event
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeNodeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createNodeMutex.RLock()
	defer fake.createNodeMutex.RUnlock()
	fake.nodeStatsMutex.RLock()
	defer fake.nodeStatsMutex.RUnlock()
	fake.nodesMutex.RLock()
	defer fake.nodesMutex.RUnlock()
	fake.removeNodeMutex.RLock()
	defer fake.removeNodeMutex.RUnlock()
	fake.stopNodeMutex.RLock()
	defer fake.stopNodeMutex.RUnlock()
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNodeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ ipfs.NodeClient = new(FakeNodeClient)
