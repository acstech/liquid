package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/karlseguin/gspec"
)

func TestReplaceFirstValueInAString(t *testing.T) {
	spec := gspec.New(t)
	filter := ReplaceFirstFactory([]core.Value{stringValue("foo"), stringValue("bar")})
	spec.Expect(filter("foobarforfoo", nil).(string)).ToEqual("barbarforfoo")
}
