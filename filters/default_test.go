package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestDefaultWithBuiltinValue(t *testing.T) {
	filter := DefaultFactory(nil)
	assert.Equal(t, filter(nil, nil).(string), "")
}

func TestDefaultWithValueOnString(t *testing.T) {
	filter := DefaultFactory([]core.Value{stringValue("d")})
	assert.Equal(t, filter("", nil), "d")
}

func TestDefaultWithValueOnArray(t *testing.T) {
	filter := DefaultFactory([]core.Value{stringValue("n/a")})
	assert.Equal(t, filter([]int{}, nil), "n/a")
}
