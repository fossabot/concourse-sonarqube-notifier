// Code generated by counterfeiter. DO NOT EDIT.
package sharedfakes

import (
	"sync"

	"github.com/concourse-sonarqube-notifier/assets/shared"
)

type FakeResultSource struct {
	GetResultStub  func(url string, authToken string, component string, metrics string) ([]byte, error)
	getResultMutex sync.RWMutex
	getResultArgsForCall []struct {
		url       string
		authToken string
		component string
		metrics   string
	}
	getResultReturns struct {
		result1 []byte
		result2 error
	}
	getResultReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetVersionsStub  func(baseUrl string, authToken string, component string) ([]byte, error)
	getVersionsMutex sync.RWMutex
	getVersionsArgsForCall []struct {
		baseUrl   string
		authToken string
		component string
	}
	getVersionsReturns struct {
		result1 []byte
		result2 error
	}
	getVersionsReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResultSource) GetResult(url string, authToken string, component string, metrics string) ([]byte, error) {
	fake.getResultMutex.Lock()
	ret, specificReturn := fake.getResultReturnsOnCall[len(fake.getResultArgsForCall)]
	fake.getResultArgsForCall = append(fake.getResultArgsForCall, struct {
		url       string
		authToken string
		component string
		metrics   string
	}{url, authToken, component, metrics})
	fake.recordInvocation("GetResult", []interface{}{url, authToken, component, metrics})
	fake.getResultMutex.Unlock()
	if fake.GetResultStub != nil {
		return fake.GetResultStub(url, authToken, component, metrics)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getResultReturns.result1, fake.getResultReturns.result2
}

func (fake *FakeResultSource) GetResultCallCount() int {
	fake.getResultMutex.RLock()
	defer fake.getResultMutex.RUnlock()
	return len(fake.getResultArgsForCall)
}

func (fake *FakeResultSource) GetResultArgsForCall(i int) (string, string, string, string) {
	fake.getResultMutex.RLock()
	defer fake.getResultMutex.RUnlock()
	return fake.getResultArgsForCall[i].url, fake.getResultArgsForCall[i].authToken, fake.getResultArgsForCall[i].component, fake.getResultArgsForCall[i].metrics
}

func (fake *FakeResultSource) GetResultReturns(result1 []byte, result2 error) {
	fake.GetResultStub = nil
	fake.getResultReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeResultSource) GetResultReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.GetResultStub = nil
	if fake.getResultReturnsOnCall == nil {
		fake.getResultReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getResultReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeResultSource) GetVersions(baseUrl string, authToken string, component string) ([]byte, error) {
	fake.getVersionsMutex.Lock()
	ret, specificReturn := fake.getVersionsReturnsOnCall[len(fake.getVersionsArgsForCall)]
	fake.getVersionsArgsForCall = append(fake.getVersionsArgsForCall, struct {
		baseUrl   string
		authToken string
		component string
	}{baseUrl, authToken, component})
	fake.recordInvocation("GetVersions", []interface{}{baseUrl, authToken, component})
	fake.getVersionsMutex.Unlock()
	if fake.GetVersionsStub != nil {
		return fake.GetVersionsStub(baseUrl, authToken, component)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getVersionsReturns.result1, fake.getVersionsReturns.result2
}

func (fake *FakeResultSource) GetVersionsCallCount() int {
	fake.getVersionsMutex.RLock()
	defer fake.getVersionsMutex.RUnlock()
	return len(fake.getVersionsArgsForCall)
}

func (fake *FakeResultSource) GetVersionsArgsForCall(i int) (string, string, string) {
	fake.getVersionsMutex.RLock()
	defer fake.getVersionsMutex.RUnlock()
	return fake.getVersionsArgsForCall[i].baseUrl, fake.getVersionsArgsForCall[i].authToken, fake.getVersionsArgsForCall[i].component
}

func (fake *FakeResultSource) GetVersionsReturns(result1 []byte, result2 error) {
	fake.GetVersionsStub = nil
	fake.getVersionsReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeResultSource) GetVersionsReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.GetVersionsStub = nil
	if fake.getVersionsReturnsOnCall == nil {
		fake.getVersionsReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getVersionsReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeResultSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getResultMutex.RLock()
	defer fake.getResultMutex.RUnlock()
	fake.getVersionsMutex.RLock()
	defer fake.getVersionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResultSource) recordInvocation(key string, args []interface{}) {
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

var _ shared.ResultSource = new(FakeResultSource)
