package generator_test

import (
	"github.com/maxbrunsfeld/counterfeiter/locator"

	. "github.com/maxbrunsfeld/counterfeiter/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generator", func() {
	Describe("generating a fake for a simple interface", func() {
		var fakeFileContents string
		var err error

		BeforeEach(func() {
			model, _ := locator.GetInterfaceFromFilePath("Something", "../fixtures/something.go")

			subject := CodeGenerator{
				Model:       *model,
				StructName:  "FakeSomething",
				PackageName: "fixturesfakes",
			}

			fakeFileContents, err = subject.GenerateFake()
		})

		It("should not fail", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match the correct file contents", func() {
			Expect(fakeFileContents).To(Equal(expectedSimpleFake))
		})
	})

	Describe("generating a fake for a typed function", func() {
		var fakeFileContents string
		var err error

		BeforeEach(func() {
			model, _ := locator.GetInterfaceFromFilePath("RequestFactory", "../fixtures/request_factory.go")

			subject := CodeGenerator{
				Model:       *model,
				StructName:  "FakeRequestFactory",
				PackageName: "fixturesfakes",
			}
			fakeFileContents, err = subject.GenerateFake()
		})

		It("should not fail", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("should produce the correct file contents", func() {
			Expect(fakeFileContents).To(Equal(expectedFuncFake))
		})
	})

	Describe("generating a fake for a package with a hyphenated import", func() {
		var fakeFileContents string
		var err error

		BeforeEach(func() {
			model, _ := locator.GetInterfaceFromFilePath("SomeInterface", "../fixtures/hyphenated_package_same_name/some_package/interface.go")

			subject := CodeGenerator{
				Model:       *model,
				StructName:  "FakeSomeInterface",
				PackageName: "fixturesfakes",
			}
			fakeFileContents, err = subject.GenerateFake()
		})

		It("should not fail", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("should produce the correct file contents", func() {
			Expect(fakeFileContents).To(Equal(expectedHyphenatedFake))
		})
	})

	Describe("generating a fake for a package with a import name differ from package name", func() {
		var fakeFileContents string
		var err error

		BeforeEach(func() {
			model, _ := locator.GetInterfaceFromFilePath("Potato", "../fixtures/alias_import_name/test/potato.go")

			subject := CodeGenerator{
				Model:       *model,
				StructName:  "FakePotato",
				PackageName: "fixturesfakes",
			}
			fakeFileContents, err = subject.GenerateFake()
		})

		It("should not fail", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("should produce the correct file contents", func() {
			Expect(fakeFileContents).To(Equal(expectedAliasFake))
		})
	})

	Describe("generating a fake for a function return like (a, b int)", func() {
		var fakeFileContents string
		var err error

		BeforeEach(func() {
			model, _ := locator.GetInterfaceFromFilePath("SomethingElse", "../fixtures/compound_return.go")

			subject := CodeGenerator{
				Model:       *model,
				StructName:  "FakeSomethingElse",
				PackageName: "fixturesfakes",
			}
			fakeFileContents, err = subject.GenerateFake()
		})

		It("should not fail", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match the correct file contents", func() {
			Expect(fakeFileContents).To(Equal(expectedCompoundReturnFake))
		})
	})

	Describe("generating a fake which imports packages with go-hyphenated names", func() {
		var generatedFileContents string
		var err error

		BeforeEach(func() {
			model, _ := locator.GetInterfaceFromFilePath("ImportsGoHyphenPackage", "../fixtures/imports_go_hyphen_package.go")

			subject := CodeGenerator{
				Model:       *model,
				StructName:  "FakeImportsGoHyphenPackage",
				PackageName: "fixturesfakes",
			}
			generatedFileContents, err = subject.GenerateFake()
		})

		It("should generate the correct code", func() {
			Expect(generatedFileContents).To(Equal(expectedImportsGoHyphenatedPackageFake))
		})

		It("should not fail", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

const expectedSimpleFake string = `// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/fixtures"
)

type FakeSomething struct {
	DoThingsStub        func(string, uint64) (int, error)
	doThingsMutex       sync.RWMutex
	doThingsArgsForCall []struct {
		arg1 string
		arg2 uint64
	}
	doThingsReturns struct {
		result1 int
		result2 error
	}
	doThingsReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	DoNothingStub        func()
	doNothingMutex       sync.RWMutex
	doNothingArgsForCall []struct{}
	DoASliceStub        func([]byte)
	doASliceMutex       sync.RWMutex
	doASliceArgsForCall []struct {
		arg1 []byte
	}
	DoAnArrayStub        func([4]byte)
	doAnArrayMutex       sync.RWMutex
	doAnArrayArgsForCall []struct {
		arg1 [4]byte
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSomething) DoThings(arg1 string, arg2 uint64) (int, error) {
	fake.doThingsMutex.Lock()
	ret, specificReturn := fake.doThingsReturnsOnCall[len(fake.doThingsArgsForCall)]
	fake.doThingsArgsForCall = append(fake.doThingsArgsForCall, struct {
		arg1 string
		arg2 uint64
	}{arg1, arg2})
	fake.recordInvocation("DoThings", []interface{}{arg1, arg2})
	fake.doThingsMutex.Unlock()
	if fake.DoThingsStub != nil {
		return fake.DoThingsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.doThingsReturns.result1, fake.doThingsReturns.result2
}

func (fake *FakeSomething) DoThingsCallCount() int {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return len(fake.doThingsArgsForCall)
}

func (fake *FakeSomething) DoThingsArgsForCall(i int) (string, uint64) {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return fake.doThingsArgsForCall[i].arg1, fake.doThingsArgsForCall[i].arg2
}

func (fake *FakeSomething) DoThingsReturns(result1 int, result2 error) {
	fake.DoThingsStub = nil
	fake.doThingsReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeSomething) DoThingsReturnsOnCall(i int, result1 int, result2 error) {
	fake.DoThingsStub = nil
	if fake.doThingsReturnsOnCall == nil {
		fake.doThingsReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.doThingsReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeSomething) DoNothing() {
	fake.doNothingMutex.Lock()
	fake.doNothingArgsForCall = append(fake.doNothingArgsForCall, struct{}{})
	fake.recordInvocation("DoNothing", []interface{}{})
	fake.doNothingMutex.Unlock()
	if fake.DoNothingStub != nil {
		fake.DoNothingStub()
	}
}

func (fake *FakeSomething) DoNothingCallCount() int {
	fake.doNothingMutex.RLock()
	defer fake.doNothingMutex.RUnlock()
	return len(fake.doNothingArgsForCall)
}

func (fake *FakeSomething) DoASlice(arg1 []byte) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.doASliceMutex.Lock()
	fake.doASliceArgsForCall = append(fake.doASliceArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("DoASlice", []interface{}{arg1Copy})
	fake.doASliceMutex.Unlock()
	if fake.DoASliceStub != nil {
		fake.DoASliceStub(arg1)
	}
}

func (fake *FakeSomething) DoASliceCallCount() int {
	fake.doASliceMutex.RLock()
	defer fake.doASliceMutex.RUnlock()
	return len(fake.doASliceArgsForCall)
}

func (fake *FakeSomething) DoASliceArgsForCall(i int) []byte {
	fake.doASliceMutex.RLock()
	defer fake.doASliceMutex.RUnlock()
	return fake.doASliceArgsForCall[i].arg1
}

func (fake *FakeSomething) DoAnArray(arg1 [4]byte) {
	fake.doAnArrayMutex.Lock()
	fake.doAnArrayArgsForCall = append(fake.doAnArrayArgsForCall, struct {
		arg1 [4]byte
	}{arg1})
	fake.recordInvocation("DoAnArray", []interface{}{arg1})
	fake.doAnArrayMutex.Unlock()
	if fake.DoAnArrayStub != nil {
		fake.DoAnArrayStub(arg1)
	}
}

func (fake *FakeSomething) DoAnArrayCallCount() int {
	fake.doAnArrayMutex.RLock()
	defer fake.doAnArrayMutex.RUnlock()
	return len(fake.doAnArrayArgsForCall)
}

func (fake *FakeSomething) DoAnArrayArgsForCall(i int) [4]byte {
	fake.doAnArrayMutex.RLock()
	defer fake.doAnArrayMutex.RUnlock()
	return fake.doAnArrayArgsForCall[i].arg1
}

func (fake *FakeSomething) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	fake.doNothingMutex.RLock()
	defer fake.doNothingMutex.RUnlock()
	fake.doASliceMutex.RLock()
	defer fake.doASliceMutex.RUnlock()
	fake.doAnArrayMutex.RLock()
	defer fake.doAnArrayMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSomething) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.Something = new(FakeSomething)
`

const expectedFuncFake string = `// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/fixtures"
)

type FakeRequestFactory struct {
	Stub        func(fixtures.Params, map[string]interface{}) (fixtures.Request, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 fixtures.Params
		arg2 map[string]interface{}
	}
	returns struct {
		result1 fixtures.Request
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 fixtures.Request
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRequestFactory) Spy(arg1 fixtures.Params, arg2 map[string]interface{}) (fixtures.Request, error) {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 fixtures.Params
		arg2 map[string]interface{}
	}{arg1, arg2})
	fake.recordInvocation("RequestFactory", []interface{}{arg1, arg2})
	fake.mutex.Unlock()
	if fake.Stub != nil {
		return fake.Stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.returns.result1, fake.returns.result2
}

func (fake *FakeRequestFactory) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeRequestFactory) ArgsForCall(i int) (fixtures.Params, map[string]interface{}) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1, fake.argsForCall[i].arg2
}

func (fake *FakeRequestFactory) Returns(result1 fixtures.Request, result2 error) {
	fake.Stub = nil
	fake.returns = struct {
		result1 fixtures.Request
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestFactory) ReturnsOnCall(i int, result1 fixtures.Request, result2 error) {
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 fixtures.Request
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 fixtures.Request
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRequestFactory) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.RequestFactory = new(FakeRequestFactory).Spy
`

var expectedHyphenatedFake = `// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	some_packagehyphen_ated "github.com/maxbrunsfeld/counterfeiter/fixtures/hyphenated_package_same_name/hyphen-ated/some_package"
	"github.com/maxbrunsfeld/counterfeiter/fixtures/hyphenated_package_same_name/some_package"
)

type FakeSomeInterface struct {
	CreateThingStub        func() some_packagehyphen_ated.Thing
	createThingMutex       sync.RWMutex
	createThingArgsForCall []struct{}
	createThingReturns struct {
		result1 some_packagehyphen_ated.Thing
	}
	createThingReturnsOnCall map[int]struct {
		result1 some_packagehyphen_ated.Thing
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSomeInterface) CreateThing() some_packagehyphen_ated.Thing {
	fake.createThingMutex.Lock()
	ret, specificReturn := fake.createThingReturnsOnCall[len(fake.createThingArgsForCall)]
	fake.createThingArgsForCall = append(fake.createThingArgsForCall, struct{}{})
	fake.recordInvocation("CreateThing", []interface{}{})
	fake.createThingMutex.Unlock()
	if fake.CreateThingStub != nil {
		return fake.CreateThingStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createThingReturns.result1
}

func (fake *FakeSomeInterface) CreateThingCallCount() int {
	fake.createThingMutex.RLock()
	defer fake.createThingMutex.RUnlock()
	return len(fake.createThingArgsForCall)
}

func (fake *FakeSomeInterface) CreateThingReturns(result1 some_packagehyphen_ated.Thing) {
	fake.CreateThingStub = nil
	fake.createThingReturns = struct {
		result1 some_packagehyphen_ated.Thing
	}{result1}
}

func (fake *FakeSomeInterface) CreateThingReturnsOnCall(i int, result1 some_packagehyphen_ated.Thing) {
	fake.CreateThingStub = nil
	if fake.createThingReturnsOnCall == nil {
		fake.createThingReturnsOnCall = make(map[int]struct {
			result1 some_packagehyphen_ated.Thing
		})
	}
	fake.createThingReturnsOnCall[i] = struct {
		result1 some_packagehyphen_ated.Thing
	}{result1}
}

func (fake *FakeSomeInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createThingMutex.RLock()
	defer fake.createThingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSomeInterface) recordInvocation(key string, args []interface{}) {
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

var _ some_package.SomeInterface = new(FakeSomeInterface)
`

const expectedCompoundReturnFake string = `// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/fixtures"
)

type FakeSomethingElse struct {
	ReturnStuffStub        func() (a, b int)
	returnStuffMutex       sync.RWMutex
	returnStuffArgsForCall []struct{}
	returnStuffReturns struct {
		result1 int
		result2 int
	}
	returnStuffReturnsOnCall map[int]struct {
		result1 int
		result2 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSomethingElse) ReturnStuff() (a, b int) {
	fake.returnStuffMutex.Lock()
	ret, specificReturn := fake.returnStuffReturnsOnCall[len(fake.returnStuffArgsForCall)]
	fake.returnStuffArgsForCall = append(fake.returnStuffArgsForCall, struct{}{})
	fake.recordInvocation("ReturnStuff", []interface{}{})
	fake.returnStuffMutex.Unlock()
	if fake.ReturnStuffStub != nil {
		return fake.ReturnStuffStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.returnStuffReturns.result1, fake.returnStuffReturns.result2
}

func (fake *FakeSomethingElse) ReturnStuffCallCount() int {
	fake.returnStuffMutex.RLock()
	defer fake.returnStuffMutex.RUnlock()
	return len(fake.returnStuffArgsForCall)
}

func (fake *FakeSomethingElse) ReturnStuffReturns(result1 int, result2 int) {
	fake.ReturnStuffStub = nil
	fake.returnStuffReturns = struct {
		result1 int
		result2 int
	}{result1, result2}
}

func (fake *FakeSomethingElse) ReturnStuffReturnsOnCall(i int, result1 int, result2 int) {
	fake.ReturnStuffStub = nil
	if fake.returnStuffReturnsOnCall == nil {
		fake.returnStuffReturnsOnCall = make(map[int]struct {
			result1 int
			result2 int
		})
	}
	fake.returnStuffReturnsOnCall[i] = struct {
		result1 int
		result2 int
	}{result1, result2}
}

func (fake *FakeSomethingElse) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.returnStuffMutex.RLock()
	defer fake.returnStuffMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSomethingElse) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.SomethingElse = new(FakeSomethingElse)
`

const expectedAliasFake string = `// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	ioAlias "io"
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/fixtures/alias_import_name/test"
)

type FakePotato struct {
	TomatoStub        func(ioAlias.Reader)
	tomatoMutex       sync.RWMutex
	tomatoArgsForCall []struct {
		arg1 ioAlias.Reader
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePotato) Tomato(arg1 ioAlias.Reader) {
	fake.tomatoMutex.Lock()
	fake.tomatoArgsForCall = append(fake.tomatoArgsForCall, struct {
		arg1 ioAlias.Reader
	}{arg1})
	fake.recordInvocation("Tomato", []interface{}{arg1})
	fake.tomatoMutex.Unlock()
	if fake.TomatoStub != nil {
		fake.TomatoStub(arg1)
	}
}

func (fake *FakePotato) TomatoCallCount() int {
	fake.tomatoMutex.RLock()
	defer fake.tomatoMutex.RUnlock()
	return len(fake.tomatoArgsForCall)
}

func (fake *FakePotato) TomatoArgsForCall(i int) ioAlias.Reader {
	fake.tomatoMutex.RLock()
	defer fake.tomatoMutex.RUnlock()
	return fake.tomatoArgsForCall[i].arg1
}

func (fake *FakePotato) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tomatoMutex.RLock()
	defer fake.tomatoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePotato) recordInvocation(key string, args []interface{}) {
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

var _ test.Potato = new(FakePotato)
`

const expectedImportsGoHyphenatedPackageFake = `// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/fixtures"
	hyphenpackage "github.com/maxbrunsfeld/counterfeiter/fixtures/go-hyphenpackage"
)

type FakeImportsGoHyphenPackage struct {
	UseHyphenTypeStub        func(hyphenpackage.HyphenType)
	useHyphenTypeMutex       sync.RWMutex
	useHyphenTypeArgsForCall []struct {
		arg1 hyphenpackage.HyphenType
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImportsGoHyphenPackage) UseHyphenType(arg1 hyphenpackage.HyphenType) {
	fake.useHyphenTypeMutex.Lock()
	fake.useHyphenTypeArgsForCall = append(fake.useHyphenTypeArgsForCall, struct {
		arg1 hyphenpackage.HyphenType
	}{arg1})
	fake.recordInvocation("UseHyphenType", []interface{}{arg1})
	fake.useHyphenTypeMutex.Unlock()
	if fake.UseHyphenTypeStub != nil {
		fake.UseHyphenTypeStub(arg1)
	}
}

func (fake *FakeImportsGoHyphenPackage) UseHyphenTypeCallCount() int {
	fake.useHyphenTypeMutex.RLock()
	defer fake.useHyphenTypeMutex.RUnlock()
	return len(fake.useHyphenTypeArgsForCall)
}

func (fake *FakeImportsGoHyphenPackage) UseHyphenTypeArgsForCall(i int) hyphenpackage.HyphenType {
	fake.useHyphenTypeMutex.RLock()
	defer fake.useHyphenTypeMutex.RUnlock()
	return fake.useHyphenTypeArgsForCall[i].arg1
}

func (fake *FakeImportsGoHyphenPackage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.useHyphenTypeMutex.RLock()
	defer fake.useHyphenTypeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImportsGoHyphenPackage) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.ImportsGoHyphenPackage = new(FakeImportsGoHyphenPackage)
`
