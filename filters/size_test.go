package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSizeOfString(t *testing.T) {
	filter := SizeFactory(nil)
	assert.Equal(t, filter("dbz", nil).(int), 3)
}

func TestSizeOfByteArray(t *testing.T) {
	filter := SizeFactory(nil)
	assert.Equal(t, filter([]byte("7 123"), nil).(int), 5)
}

func TestSizeOfIntArray(t *testing.T) {
	filter := SizeFactory(nil)
	assert.Equal(t, filter([]int{2, 4, 5, 6}, nil).(int), 4)
}

func TestSizeOfBoolArray(t *testing.T) {
	filter := SizeFactory(nil)
	assert.Equal(t, filter([]bool{true, false, true, true, false}, nil).(int), 5)
}

func TestSizeOfMap(t *testing.T) {
	filter := SizeFactory(nil)
	assert.Equal(t, filter(map[string]int{"over": 9000}, nil).(int), 1)
}

func TestSizeOfSometingInvalid(t *testing.T) {
	filter := SizeFactory(nil)
	assert.Equal(t, filter(false, nil).(bool), false)
}
