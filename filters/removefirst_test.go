package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/karlseguin/gspec"
)

func TestRemovesFirstValueFromAString(t *testing.T) {
	spec := gspec.New(t)
	filter := RemoveFirstFactory(nil, []core.Value{stringValue("foo")})
	spec.Expect(filter("foobarforfoo", nil).(string)).ToEqual("barforfoo")
}
