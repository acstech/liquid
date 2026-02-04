package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestSplitsAStringOnDefaultSpace(t *testing.T) {
	filter := SplitFactory([]core.Value{})
	values := filter("hello world", nil).([]string)
	assert.Equal(t, len(values), 2)
	assert.Equal(t, values[0], "hello")
	assert.Equal(t, values[1], "world")
}

func TestSplitsAStringOnSpecifiedValue(t *testing.T) {
	filter := SplitFactory([]core.Value{stringValue("..")})
	values := filter([]byte("hel..lowo..rl..d"), nil).([]string)
	assert.Equal(t, len(values), 4)
	assert.Equal(t, values[0], "hel")
	assert.Equal(t, values[1], "lowo")
	assert.Equal(t, values[2], "rl")
	assert.Equal(t, values[3], "d")
}

func TestSplitsAStringOnADynamicValue(t *testing.T) {
	filter := SplitFactory([]core.Value{dynamicValue("sep")})
	values := filter("over;9000;!", params("sep", ";")).([]string)
	assert.Equal(t, len(values), 3)
	assert.Equal(t, values[0], "over")
	assert.Equal(t, values[1], "9000")
	assert.Equal(t, values[2], "!")
}
