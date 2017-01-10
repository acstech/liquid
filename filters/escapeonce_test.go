package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
)

func TestEscapesOnceAString(t *testing.T) {
	spec := gspec.New(t)
	filter := EscapeOnceFactory(nil, nil)
	spec.Expect(filter("<b>hello</b>", nil).(string)).ToEqual("&lt;b&gt;hello&lt;/b&gt;")
}

func TestEscapesOnceAStringWithEscapedValues(t *testing.T) {
	spec := gspec.New(t)
	filter := EscapeOnceFactory(nil, nil)
	spec.Expect(filter("<b>hello</b>&lt;b&gt;hello&lt;/b&gt;", nil).(string)).ToEqual("&lt;b&gt;hello&lt;/b&gt;&lt;b&gt;hello&lt;/b&gt;")
}
