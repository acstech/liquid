package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapitalizesAString(t *testing.T) {
	filter := CapitalizeFactory(nil)
	assert.Equal(t, string(filter("tiger got to hunt, bird got to fly", nil).([]byte)), "Tiger Got To Hunt, Bird Got To Fly")
}

func TestCapitalizesBytes(t *testing.T) {
	filter := CapitalizeFactory(nil)
	assert.Equal(t, string(filter([]byte("Science is magic that works "), nil).([]byte)), "Science Is Magic That Works ")
}

func TestCapitalizePassThroughOnInvalidType(t *testing.T) {
	filter := CapitalizeFactory(nil)
	assert.Equal(t, filter(123, nil).(int), 123)
}
