package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpcasesAString(t *testing.T) {
	filter := UpcaseFactory(nil)
	assert.Equal(t, filter("dbz", nil).(string), "DBZ")
}

func TestUpcasesBytes(t *testing.T) {
	filter := UpcaseFactory(nil)
	assert.Equal(t, string(filter([]byte("dbz"), nil).([]byte)), "DBZ")
}

func TestUpcasesPassThroughOnInvalidType(t *testing.T) {
	filter := UpcaseFactory(nil)
	assert.Equal(t, filter(123, nil).(int), 123)
}
