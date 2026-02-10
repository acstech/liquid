package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestJoinsStringsWithTheSpecifiedGlue(t *testing.T) {
	filter := JoinFactory([]core.Value{stringValue("..")})
	assert.Equal(t, string(filter([]string{"leto", "atreides"}, nil).([]byte)), "leto..atreides")
}

func TestJoinsVariousTypesWithTheDefaultGlue(t *testing.T) {
	filter := JoinFactory(nil)
	assert.Equal(t, string(filter([]interface{}{"leto", 123, true}, nil).([]byte)), "leto 123 true")
}

func TestJoinPassthroughOnEmptyArray(t *testing.T) {
	filter := JoinFactory(nil)
	arr := [0]int{}
	assert.Equal(t, filter(arr, nil).([0]int), arr)
}

func TestJoinPassthroughOnEmptySlice(t *testing.T) {
	filter := JoinFactory(nil)
	arr := []int{}
	assert.Equal(t, len(filter(arr, nil).([]int)), 0)
}

func TestJoinPassthroughOnInvalidType(t *testing.T) {
	filter := JoinFactory(nil)
	assert.Equal(t, filter("hahah", nil).(string), "hahah")
}
