package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringHtml(t *testing.T) {
	filter := StripHtmlFactory(nil)
	assert.Equal(t, filter("<style>*{margin:0}</style>hello <b>world</b>", nil).(string), "hello world")
}
