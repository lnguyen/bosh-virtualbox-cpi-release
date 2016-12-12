// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-utils/crypto"
)

type FakeDigestProvider struct {
	CreateFromFileStub        func(path string, algorithm crypto.DigestAlgorithm) (crypto.Digest, error)
	createFromFileMutex       sync.RWMutex
	createFromFileArgsForCall []struct {
		path      string
		algorithm crypto.DigestAlgorithm
	}
	createFromFileReturns struct {
		result1 crypto.Digest
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDigestProvider) CreateFromFile(path string, algorithm crypto.DigestAlgorithm) (crypto.Digest, error) {
	fake.createFromFileMutex.Lock()
	fake.createFromFileArgsForCall = append(fake.createFromFileArgsForCall, struct {
		path      string
		algorithm crypto.DigestAlgorithm
	}{path, algorithm})
	fake.recordInvocation("CreateFromFile", []interface{}{path, algorithm})
	fake.createFromFileMutex.Unlock()
	if fake.CreateFromFileStub != nil {
		return fake.CreateFromFileStub(path, algorithm)
	} else {
		return fake.createFromFileReturns.result1, fake.createFromFileReturns.result2
	}
}

func (fake *FakeDigestProvider) CreateFromFileCallCount() int {
	fake.createFromFileMutex.RLock()
	defer fake.createFromFileMutex.RUnlock()
	return len(fake.createFromFileArgsForCall)
}

func (fake *FakeDigestProvider) CreateFromFileArgsForCall(i int) (string, crypto.DigestAlgorithm) {
	fake.createFromFileMutex.RLock()
	defer fake.createFromFileMutex.RUnlock()
	return fake.createFromFileArgsForCall[i].path, fake.createFromFileArgsForCall[i].algorithm
}

func (fake *FakeDigestProvider) CreateFromFileReturns(result1 crypto.Digest, result2 error) {
	fake.CreateFromFileStub = nil
	fake.createFromFileReturns = struct {
		result1 crypto.Digest
		result2 error
	}{result1, result2}
}

func (fake *FakeDigestProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createFromFileMutex.RLock()
	defer fake.createFromFileMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDigestProvider) recordInvocation(key string, args []interface{}) {
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

var _ crypto.DigestProvider = new(FakeDigestProvider)
