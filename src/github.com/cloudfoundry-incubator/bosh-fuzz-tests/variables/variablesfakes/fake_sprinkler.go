// Code generated by counterfeiter. DO NOT EDIT.
package variablesfakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bosh-fuzz-tests/variables"
)

type FakeSprinkler struct {
	SprinklePlaceholdersStub        func(manifestPath string) (map[string]interface{}, error)
	sprinklePlaceholdersMutex       sync.RWMutex
	sprinklePlaceholdersArgsForCall []struct {
		manifestPath string
	}
	sprinklePlaceholdersReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	sprinklePlaceholdersReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSprinkler) SprinklePlaceholders(manifestPath string) (map[string]interface{}, error) {
	fake.sprinklePlaceholdersMutex.Lock()
	ret, specificReturn := fake.sprinklePlaceholdersReturnsOnCall[len(fake.sprinklePlaceholdersArgsForCall)]
	fake.sprinklePlaceholdersArgsForCall = append(fake.sprinklePlaceholdersArgsForCall, struct {
		manifestPath string
	}{manifestPath})
	fake.recordInvocation("SprinklePlaceholders", []interface{}{manifestPath})
	fake.sprinklePlaceholdersMutex.Unlock()
	if fake.SprinklePlaceholdersStub != nil {
		return fake.SprinklePlaceholdersStub(manifestPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.sprinklePlaceholdersReturns.result1, fake.sprinklePlaceholdersReturns.result2
}

func (fake *FakeSprinkler) SprinklePlaceholdersCallCount() int {
	fake.sprinklePlaceholdersMutex.RLock()
	defer fake.sprinklePlaceholdersMutex.RUnlock()
	return len(fake.sprinklePlaceholdersArgsForCall)
}

func (fake *FakeSprinkler) SprinklePlaceholdersArgsForCall(i int) string {
	fake.sprinklePlaceholdersMutex.RLock()
	defer fake.sprinklePlaceholdersMutex.RUnlock()
	return fake.sprinklePlaceholdersArgsForCall[i].manifestPath
}

func (fake *FakeSprinkler) SprinklePlaceholdersReturns(result1 map[string]interface{}, result2 error) {
	fake.SprinklePlaceholdersStub = nil
	fake.sprinklePlaceholdersReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeSprinkler) SprinklePlaceholdersReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.SprinklePlaceholdersStub = nil
	if fake.sprinklePlaceholdersReturnsOnCall == nil {
		fake.sprinklePlaceholdersReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.sprinklePlaceholdersReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeSprinkler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sprinklePlaceholdersMutex.RLock()
	defer fake.sprinklePlaceholdersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSprinkler) recordInvocation(key string, args []interface{}) {
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

var _ variables.Sprinkler = new(FakeSprinkler)
