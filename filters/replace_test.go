package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/karlseguin/gspec"
)

func TestReplaceValuesInAString(t *testing.T) {
	spec := gspec.New(t)
	filter := ReplaceFactory([]core.Value{stringValue("foo"), stringValue("bar")})
	spec.Expect(filter("foobarforfoo", nil).(string)).ToEqual("barbarforbar")
}

func TestReplaceWithDynamicValues(t *testing.T) {
	spec := gspec.New(t)
	filter := ReplaceFactory([]core.Value{dynamicValue("f"), dynamicValue("b")})
	spec.Expect(filter("foobarforfoo", params("f", "oo", "b", "br")).(string)).ToEqual("fbrbarforfbr")
}
