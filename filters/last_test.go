package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsTheLastItem(t *testing.T) {
	filter := LastFactory(nil)
	assert.Equal(t, filter([]string{"leto", "atreides"}, nil).(string), "atreides")
}

func TestReturnsTheLastItemIfOnlyOneItem(t *testing.T) {
	filter := LastFactory(nil)
	assert.Equal(t, filter([]string{"leto"}, nil).(string), "leto")
}

func TestReturnsTheLastItemOfAnArray(t *testing.T) {
	filter := LastFactory(nil)
	arr := [4]int{1, 2, 3, 48}
	assert.Equal(t, filter(arr, nil).(int), 48)
}

func TestLastPassthroughOnEmptyArray(t *testing.T) {
	filter := LastFactory(nil)
	arr := [0]int{}
	assert.Equal(t, filter(arr, nil).([0]int), arr)
}

func TestLastPassthroughOnEmptySlice(t *testing.T) {
	filter := LastFactory(nil)
	arr := []int{}
	assert.Equal(t, len(filter(arr, nil).([]int)), 0)
}

func TestLastPassthroughOnInvalidType(t *testing.T) {
	filter := LastFactory(nil)
	assert.Equal(t, filter("hahah", nil).(string), "hahah")
}
