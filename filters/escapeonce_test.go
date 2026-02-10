package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscapesOnceAString(t *testing.T) {
	filter := EscapeOnceFactory(nil)
	assert.Equal(t, filter("<b>hello</b>", nil).(string), "&lt;b&gt;hello&lt;/b&gt;")
}

func TestEscapesOnceAStringWithEscapedValues(t *testing.T) {
	filter := EscapeOnceFactory(nil)
	assert.Equal(t, filter("<b>hello</b>&lt;b&gt;hello&lt;/b&gt;", nil).(string), "&lt;b&gt;hello&lt;/b&gt;&lt;b&gt;hello&lt;/b&gt;")
}
