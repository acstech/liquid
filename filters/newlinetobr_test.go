package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
)

func TestReplacesNewlinesWithBr(t *testing.T) {
	spec := gspec.New(t)
	filter := NewLineToBrFactory(nil)
	spec.Expect(filter("f\no\ro\n\r", nil).(string)).ToEqual("f<br />\no<br />\no<br />\n")
}
