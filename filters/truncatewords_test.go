package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestTruncateWordsWhenTooLong(t *testing.T) {
	filter := TruncateWordsFactory([]core.Value{intValue(2)})
	assert.Equal(t, filter("hello world how's it going", nil).(string), "hello world...")
}

func TestTruncateWordsWhenTooShort(t *testing.T) {
	filter := TruncateWordsFactory([]core.Value{intValue(6)})
	assert.Equal(t, filter("hello world how's it going", nil).(string), "hello world how's it going")
}

func TestTruncateWordsWhenJustRight(t *testing.T) {
	filter := TruncateWordsFactory([]core.Value{intValue(5)})
	assert.Equal(t, filter("hello world how's it going", nil).(string), "hello world how's it going")
}

func TestTruncateWordsWithCustomAppend(t *testing.T) {
	filter := TruncateWordsFactory([]core.Value{intValue(3), stringValue(" (more...)")})
	assert.Equal(t, filter("hello world how's it going", nil).(string), "hello world how's (more...)")
}

func TestTruncateWordsWithShortWords(t *testing.T) {
	filter := TruncateWordsFactory([]core.Value{dynamicValue("max")})
	assert.Equal(t, filter("I  think  a  feature good", params("max", 2)).(string), "I  think  a  feature...")
}
