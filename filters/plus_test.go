package filters

import (
	"testing"
	"time"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestPlusAnIntToAnInt(t *testing.T) {
	filter := PlusFactory([]core.Value{intValue(5)})
	assert.Equal(t, filter(43, nil).(int), 48)
}

func TestPlusAFloattToAnInt(t *testing.T) {
	filter := PlusFactory([]core.Value{floatValue(5.11)})
	assert.Equal(t, filter(43, nil).(float64), 48.11)
}

func TestPlusAnIntToAFloat(t *testing.T) {
	filter := PlusFactory([]core.Value{intValue(5)})
	assert.Equal(t, filter(43.2, nil).(float64), 48.2)
}

func TestPlusAnIntToNow(t *testing.T) {
	filter := PlusFactory([]core.Value{intValue(61)})
	assert.Equal(t, filter("now", nil).(time.Time), core.Now().Add(time.Minute*61))
}

func TestPlusAnIntToATime(t *testing.T) {
	now := time.Now()
	filter := PlusFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter(now, nil).(time.Time), now.Add(time.Minute*7))
}

func TestPlusAnIntToAStringAsAnInt(t *testing.T) {
	filter := PlusFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter("33", nil).(int), 40)
}

func TestPlusAnIntToAStringAsAFloat(t *testing.T) {
	filter := PlusFactory([]core.Value{floatValue(2.2)})
	assert.Equal(t, filter("33.11", nil).(float64), 35.31)
}

func TestPlusAnIntToBytesAsAnInt(t *testing.T) {
	filter := PlusFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter([]byte("34"), nil).(int), 41)
}

func TestPlusAnDynamicIntToBytesAsAnInt(t *testing.T) {
	filter := PlusFactory([]core.Value{dynamicValue("fee")})
	assert.Equal(t, filter([]byte("34"), params("fee", 5)).(int), 39)
}

func TestPlusAnDynamicFloatToBytesAsAnInt(t *testing.T) {
	filter := PlusFactory([]core.Value{dynamicValue("fee")})
	assert.Equal(t, filter([]byte("34"), params("fee", 5.1)).(float64), 39.1)
}
