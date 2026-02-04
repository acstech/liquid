package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestPrependToAString(t *testing.T) {
	filter := PrependFactory([]core.Value{stringValue("?!")})
	assert.Equal(t, filter("dbz", nil).(string), "?!dbz")
}

func TestPrependToBytes(t *testing.T) {
	filter := PrependFactory([]core.Value{stringValue("boring")})
	assert.Equal(t, filter([]byte("so"), nil).(string), "boringso")
}

func TestPrependADynamicValue(t *testing.T) {
	filter := PrependFactory([]core.Value{dynamicValue("local.currency")})
	data := map[string]interface{}{
		"local": map[string]string{
			"currency": "$",
		},
	}
	assert.Equal(t, filter("100", data).(string), "$100")
}
