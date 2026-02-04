package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripsNewLinesFromStirng(t *testing.T) {
	filter := StripNewLinesFactory(nil)
	assert.Equal(t, filter("f\no\ro\n\r", nil).(string), "foo")
}
