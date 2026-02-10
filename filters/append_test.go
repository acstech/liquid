package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestAppendToAString(t *testing.T) {
	filter := AppendFactory([]core.Value{stringValue("?!")})
	assert.Equal(t, filter("dbz", nil).(string), "dbz?!")
}

func TestAppendToBytes(t *testing.T) {
	filter := AppendFactory([]core.Value{stringValue("boring")})
	assert.Equal(t, filter([]byte("so"), nil).(string), "soboring")
}

func TestAppendADynamicValue(t *testing.T) {
	filter := AppendFactory([]core.Value{dynamicValue("local.currency")})
	data := map[string]interface{}{
		"local": map[string]string{
			"currency": "$",
		},
	}
	assert.Equal(t, filter("100", data).(string), "100$")
}
