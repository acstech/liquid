package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolvesAnInvalidValueFromNil(t *testing.T) {
	assert.Nil(t, Resolve(nil, "something"))
}

func TestResolvesAnInvalidValueFromAMap(t *testing.T) {
	m := map[string]string{"name": "leto"}
	assert.Equal(t, Resolve(m, "age"), "")
}

func TestResolvesAnInvalidValueFromAStruct(t *testing.T) {
	m := &Person{"Leto", 3231}
	assert.Nil(t, Resolve(m, "IsGholas"))
}

func TestResolvesAValueFromAMap(t *testing.T) {
	m := map[string]string{"name": "leto"}
	assert.Equal(t, Resolve(m, "name"), "leto")
}

func TestResolvesAValueFromAStruct(t *testing.T) {
	m := &Person{"Leto", 3231}
	assert.Equal(t, Resolve(m, "name"), "Leto")
	assert.Equal(t, Resolve(m, "age"), 3231)
}

type Person struct {
	Name string
	Age  int
}
