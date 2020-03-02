// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/api/protocol"
	"github.com/trustbloc/sidetree-core-go/pkg/batch"
	"github.com/trustbloc/sidetree-core-go/pkg/batch/cutter"
)

type BatchContext struct {
	ProtocolStub        func() protocol.Client
	protocolMutex       sync.RWMutex
	protocolArgsForCall []struct{}
	protocolReturns     struct {
		result1 protocol.Client
	}
	protocolReturnsOnCall map[int]struct {
		result1 protocol.Client
	}
	CASStub        func() batch.CASClient
	cASMutex       sync.RWMutex
	cASArgsForCall []struct{}
	cASReturns     struct {
		result1 batch.CASClient
	}
	cASReturnsOnCall map[int]struct {
		result1 batch.CASClient
	}
	BlockchainStub        func() batch.BlockchainClient
	blockchainMutex       sync.RWMutex
	blockchainArgsForCall []struct{}
	blockchainReturns     struct {
		result1 batch.BlockchainClient
	}
	blockchainReturnsOnCall map[int]struct {
		result1 batch.BlockchainClient
	}
	OperationQueueStub        func() cutter.OperationQueue
	operationQueueMutex       sync.RWMutex
	operationQueueArgsForCall []struct{}
	operationQueueReturns     struct {
		result1 cutter.OperationQueue
	}
	operationQueueReturnsOnCall map[int]struct {
		result1 cutter.OperationQueue
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BatchContext) Protocol() protocol.Client {
	fake.protocolMutex.Lock()
	ret, specificReturn := fake.protocolReturnsOnCall[len(fake.protocolArgsForCall)]
	fake.protocolArgsForCall = append(fake.protocolArgsForCall, struct{}{})
	fake.recordInvocation("Protocol", []interface{}{})
	fake.protocolMutex.Unlock()
	if fake.ProtocolStub != nil {
		return fake.ProtocolStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.protocolReturns.result1
}

func (fake *BatchContext) ProtocolCallCount() int {
	fake.protocolMutex.RLock()
	defer fake.protocolMutex.RUnlock()
	return len(fake.protocolArgsForCall)
}

func (fake *BatchContext) ProtocolReturns(result1 protocol.Client) {
	fake.ProtocolStub = nil
	fake.protocolReturns = struct {
		result1 protocol.Client
	}{result1}
}

func (fake *BatchContext) ProtocolReturnsOnCall(i int, result1 protocol.Client) {
	fake.ProtocolStub = nil
	if fake.protocolReturnsOnCall == nil {
		fake.protocolReturnsOnCall = make(map[int]struct {
			result1 protocol.Client
		})
	}
	fake.protocolReturnsOnCall[i] = struct {
		result1 protocol.Client
	}{result1}
}

func (fake *BatchContext) CAS() batch.CASClient {
	fake.cASMutex.Lock()
	ret, specificReturn := fake.cASReturnsOnCall[len(fake.cASArgsForCall)]
	fake.cASArgsForCall = append(fake.cASArgsForCall, struct{}{})
	fake.recordInvocation("CAS", []interface{}{})
	fake.cASMutex.Unlock()
	if fake.CASStub != nil {
		return fake.CASStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cASReturns.result1
}

func (fake *BatchContext) CASCallCount() int {
	fake.cASMutex.RLock()
	defer fake.cASMutex.RUnlock()
	return len(fake.cASArgsForCall)
}

func (fake *BatchContext) CASReturns(result1 batch.CASClient) {
	fake.CASStub = nil
	fake.cASReturns = struct {
		result1 batch.CASClient
	}{result1}
}

func (fake *BatchContext) CASReturnsOnCall(i int, result1 batch.CASClient) {
	fake.CASStub = nil
	if fake.cASReturnsOnCall == nil {
		fake.cASReturnsOnCall = make(map[int]struct {
			result1 batch.CASClient
		})
	}
	fake.cASReturnsOnCall[i] = struct {
		result1 batch.CASClient
	}{result1}
}

func (fake *BatchContext) Blockchain() batch.BlockchainClient {
	fake.blockchainMutex.Lock()
	ret, specificReturn := fake.blockchainReturnsOnCall[len(fake.blockchainArgsForCall)]
	fake.blockchainArgsForCall = append(fake.blockchainArgsForCall, struct{}{})
	fake.recordInvocation("Blockchain", []interface{}{})
	fake.blockchainMutex.Unlock()
	if fake.BlockchainStub != nil {
		return fake.BlockchainStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.blockchainReturns.result1
}

func (fake *BatchContext) BlockchainCallCount() int {
	fake.blockchainMutex.RLock()
	defer fake.blockchainMutex.RUnlock()
	return len(fake.blockchainArgsForCall)
}

func (fake *BatchContext) BlockchainReturns(result1 batch.BlockchainClient) {
	fake.BlockchainStub = nil
	fake.blockchainReturns = struct {
		result1 batch.BlockchainClient
	}{result1}
}

func (fake *BatchContext) BlockchainReturnsOnCall(i int, result1 batch.BlockchainClient) {
	fake.BlockchainStub = nil
	if fake.blockchainReturnsOnCall == nil {
		fake.blockchainReturnsOnCall = make(map[int]struct {
			result1 batch.BlockchainClient
		})
	}
	fake.blockchainReturnsOnCall[i] = struct {
		result1 batch.BlockchainClient
	}{result1}
}

func (fake *BatchContext) OperationQueue() cutter.OperationQueue {
	fake.operationQueueMutex.Lock()
	ret, specificReturn := fake.operationQueueReturnsOnCall[len(fake.operationQueueArgsForCall)]
	fake.operationQueueArgsForCall = append(fake.operationQueueArgsForCall, struct{}{})
	fake.recordInvocation("OperationQueue", []interface{}{})
	fake.operationQueueMutex.Unlock()
	if fake.OperationQueueStub != nil {
		return fake.OperationQueueStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.operationQueueReturns.result1
}

func (fake *BatchContext) OperationQueueCallCount() int {
	fake.operationQueueMutex.RLock()
	defer fake.operationQueueMutex.RUnlock()
	return len(fake.operationQueueArgsForCall)
}

func (fake *BatchContext) OperationQueueReturns(result1 cutter.OperationQueue) {
	fake.OperationQueueStub = nil
	fake.operationQueueReturns = struct {
		result1 cutter.OperationQueue
	}{result1}
}

func (fake *BatchContext) OperationQueueReturnsOnCall(i int, result1 cutter.OperationQueue) {
	fake.OperationQueueStub = nil
	if fake.operationQueueReturnsOnCall == nil {
		fake.operationQueueReturnsOnCall = make(map[int]struct {
			result1 cutter.OperationQueue
		})
	}
	fake.operationQueueReturnsOnCall[i] = struct {
		result1 cutter.OperationQueue
	}{result1}
}

func (fake *BatchContext) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.protocolMutex.RLock()
	defer fake.protocolMutex.RUnlock()
	fake.cASMutex.RLock()
	defer fake.cASMutex.RUnlock()
	fake.blockchainMutex.RLock()
	defer fake.blockchainMutex.RUnlock()
	fake.operationQueueMutex.RLock()
	defer fake.operationQueueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BatchContext) recordInvocation(key string, args []interface{}) {
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

var _ batch.Context = new(BatchContext)
