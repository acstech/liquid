package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseDoesNothingOnInvalidType(t *testing.T) {
	filter := ReverseFactory(nil)
	assert.Equal(t, filter(123, nil).(int), 123)
}

func TestReverseAnEvenLengthString(t *testing.T) {
	filter := ReverseFactory(nil)
	assert.Equal(t, string(filter("123456", nil).([]byte)), "654321")
}

func TestReverseAnOddLengthString(t *testing.T) {
	filter := ReverseFactory(nil)
	assert.Equal(t, string(filter("12345", nil).([]byte)), "54321")
}

func TestReverseASingleCharacterString(t *testing.T) {
	filter := ReverseFactory(nil)
	assert.Equal(t, string(filter("1", nil).([]byte)), "1")
}

func TestReverseAnEvenLengthArray(t *testing.T) {
	filter := ReverseFactory(nil)
	values := filter([]int{1, 2, 3, 4}, nil).([]int)
	assert.Equal(t, len(values), 4)
	assert.Equal(t, values[0], 4)
	assert.Equal(t, values[1], 3)
	assert.Equal(t, values[2], 2)
	assert.Equal(t, values[3], 1)
}

func TestReverseAnOddLengthArray(t *testing.T) {
	filter := ReverseFactory(nil)
	values := filter([]float64{1.1, 2.2, 3.3}, nil).([]float64)
	assert.Equal(t, len(values), 3)
	assert.Equal(t, values[0], 3.3)
	assert.Equal(t, values[1], 2.2)
	assert.Equal(t, values[2], 1.1)
}

func TestReverseASingleElementArray(t *testing.T) {
	filter := ReverseFactory(nil)
	values := filter([]bool{true}, nil).([]bool)
	assert.Equal(t, len(values), 1)
	assert.Equal(t, values[0], true)
}
