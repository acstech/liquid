package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestRemovesFirstValueFromAString(t *testing.T) {
	filter := RemoveFirstFactory([]core.Value{stringValue("foo")})
	assert.Equal(t, filter("foobarforfoo", nil).(string), "barforfoo")
}
