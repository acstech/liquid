package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
)

func TestStringHtml(t *testing.T) {
	spec := gspec.New(t)
	filter := StripHtmlFactory(nil)
	spec.Expect(filter("<style>*{margin:0}</style>hello <b>world</b>", nil).(string)).ToEqual("hello world")
}
