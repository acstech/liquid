package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestTruncateAString(t *testing.T) {
	filter := TruncateFactory([]core.Value{intValue(5)})
	assert.Equal(t, filter("1234567", nil).(string), "12...")
}

func TestTruncateAShortString(t *testing.T) {
	filter := TruncateFactory([]core.Value{intValue(100), stringValue("")})
	assert.Equal(t, filter("1234567", nil).(string), "1234567")
}

func TestTruncateAPerfectString(t *testing.T) {
	filter := TruncateFactory([]core.Value{intValue(7), stringValue("")})
	assert.Equal(t, filter("1234567", nil).(string), "1234567")
}

func TestTruncateAnAlmostPerfectString(t *testing.T) {
	filter := TruncateFactory([]core.Value{intValue(6), stringValue("")})
	assert.Equal(t, filter("1234567", nil).(string), "123456")
}

func TestTruncateAStringFromAFloat(t *testing.T) {
	filter := TruncateFactory([]core.Value{floatValue(3.3), stringValue(".")})
	assert.Equal(t, filter("1234567", nil).(string), "12.")
}

func TestTruncateAStringFromAString(t *testing.T) {
	filter := TruncateFactory([]core.Value{stringValue("4"), stringValue("")})
	assert.Equal(t, filter("1234567", nil).(string), "1234")
}

func TestTruncateAStringFromAnInvalidString(t *testing.T) {
	filter := TruncateFactory([]core.Value{stringValue("abc"), stringValue("")})
	assert.Equal(t, filter("1234567", nil).(string), "1234567")
}

func TestTruncateAnInvalidValue(t *testing.T) {
	filter := TruncateFactory([]core.Value{intValue(4)})
	assert.Equal(t, filter(555, nil).(string), "5...")
}
