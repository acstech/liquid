package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
)

func TestStripsNewLinesFromStirng(t *testing.T) {
	spec := gspec.New(t)
	filter := StripNewLinesFactory(nil, nil)
	spec.Expect(filter("f\no\ro\n\r", nil).(string)).ToEqual("foo")
}
