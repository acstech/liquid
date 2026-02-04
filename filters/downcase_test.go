package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDowncasesAString(t *testing.T) {
	filter := DowncaseFactory(nil)
	assert.Equal(t, filter("DBZ", nil).(string), "dbz")
}

func TestDowncasesBytes(t *testing.T) {
	filter := DowncaseFactory(nil)
	assert.Equal(t, string(filter([]byte("DBZ"), nil).([]byte)), "dbz")
}

func TestDowncasesPassThroughOnInvalidType(t *testing.T) {
	filter := DowncaseFactory(nil)
	assert.Equal(t, filter(123, nil).(int), 123)
}
