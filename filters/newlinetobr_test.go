package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplacesNewlinesWithBr(t *testing.T) {
	filter := NewLineToBrFactory(nil)
	assert.Equal(t, filter("f\no\ro\n\r", nil).(string), "f<br />\no<br />\no<br />\n")
}
