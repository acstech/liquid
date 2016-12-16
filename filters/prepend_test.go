package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/karlseguin/gspec"
)

func TestPrependToAString(t *testing.T) {
	spec := gspec.New(t)
	filter := PrependFactory(nil, []core.Value{stringValue("?!")})
	spec.Expect(filter("dbz", nil).(string)).ToEqual("?!dbz")
}

func TestPrependToBytes(t *testing.T) {
	spec := gspec.New(t)
	filter := PrependFactory(nil, []core.Value{stringValue("boring")})
	spec.Expect(filter([]byte("so"), nil).(string)).ToEqual("boringso")
}

func TestPrependADynamicValue(t *testing.T) {
	spec := gspec.New(t)
	filter := PrependFactory(nil, []core.Value{dynamicValue("local.currency")})
	data := map[string]interface{}{
		"local": map[string]string{
			"currency": "$",
		},
	}
	spec.Expect(filter("100", data).(string)).ToEqual("$100")
}
