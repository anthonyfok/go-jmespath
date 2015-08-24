package jmespath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type scalars struct {
	Foo string
	Bar string
}

type sliceType struct {
	A string
	B []scalars
	C []*scalars
}

func TestCanSupportEmptyInterface(t *testing.T) {
	assert := assert.New(t)
	data := make(map[string]interface{})
	data["foo"] = "bar"
	result, err := Search("foo", data)
	assert.Nil(err)
	assert.Equal("bar", result)
}

func TestCanSupportUserDefinedStructsValue(t *testing.T) {
	assert := assert.New(t)
	s := scalars{Foo: "one", Bar: "bar"}
	result, err := Search("Foo", s)
	assert.Nil(err)
	assert.Equal("one", result)
}

func TestCanSupportUserDefinedStructsRef(t *testing.T) {
	assert := assert.New(t)
	s := scalars{Foo: "one", Bar: "bar"}
	result, err := Search("Foo", &s)
	assert.Nil(err)
	assert.Equal("one", result)
}

func TestCanSupportStructWithSlice(t *testing.T) {
	assert := assert.New(t)
	data := sliceType{A: "foo", B: []scalars{scalars{"f1", "b1"}, scalars{"correct", "b2"}}}
	result, err := Search("B[-1].Foo", data)
	assert.Nil(err)
	assert.Equal("correct", result)
}

func TestCanSupportStructWithSlicePointer(t *testing.T) {
	assert := assert.New(t)
	data := sliceType{A: "foo", C: []*scalars{&scalars{"f1", "b1"}, &scalars{"correct", "b2"}}}
	result, err := Search("C[-1].Foo", data)
	assert.Nil(err)
	assert.Equal("correct", result)
}

func TestWillAutomaticallyCapitalizeFieldNames(t *testing.T) {
	assert := assert.New(t)
	s := scalars{Foo: "one", Bar: "bar"}
	// Note that there's a lower cased "foo" instead of "Foo",
	// but it should still correspond to the Foo field in the
	// scalars struct
	result, err := Search("foo", &s)
	assert.Nil(err)
	assert.Equal("one", result)
}

func TestCanSupportStructWithSliceLowerCased(t *testing.T) {
	assert := assert.New(t)
	data := sliceType{A: "foo", B: []scalars{scalars{"f1", "b1"}, scalars{"correct", "b2"}}}
	result, err := Search("b[-1].foo", data)
	assert.Nil(err)
	assert.Equal("correct", result)
}
