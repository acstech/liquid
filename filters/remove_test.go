package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/karlseguin/gspec"
)

func TestRemovesValuesFromAString(t *testing.T) {
	spec := gspec.New(t)
	filter := RemoveFactory([]core.Value{stringValue("foo")})
	spec.Expect(filter("foobarforfoo", nil).(string)).ToEqual("barfor")
}
