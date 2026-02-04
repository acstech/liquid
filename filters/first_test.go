package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsTheFirstItem(t *testing.T) {
	filter := FirstFactory(nil)
	assert.Equal(t, filter([]string{"leto", "atreides"}, nil).(string), "leto")
}

func TestReturnsTheFirstItemIfOnlyOneItem(t *testing.T) {
	filter := FirstFactory(nil)
	assert.Equal(t, filter([]string{"leto"}, nil).(string), "leto")
}

func TestReturnsTheFirstItemOfAnArray(t *testing.T) {
	filter := FirstFactory(nil)
	arr := [4]int{12, 2, 3, 48}
	assert.Equal(t, filter(arr, nil).(int), 12)
}

func TestFirstPassthroughOnEmptyArray(t *testing.T) {
	filter := FirstFactory(nil)
	arr := [0]int{}
	assert.Equal(t, filter(arr, nil).([0]int), arr)
}

func TestFirstPassthroughOnEmptySlice(t *testing.T) {
	filter := FirstFactory(nil)
	arr := []int{}
	assert.Equal(t, len(filter(arr, nil).([]int)), 0)
}

func TestFirstPassthroughOnInvalidType(t *testing.T) {
	filter := FirstFactory(nil)
	assert.Equal(t, filter("hahah", nil).(string), "hahah")
}
