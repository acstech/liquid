package filters

import (
	"testing"
	"time"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestMinusAnIntToAnInt(t *testing.T) {
	filter := MinusFactory([]core.Value{intValue(5)})
	assert.Equal(t, filter(43, nil).(int), 38)
}

func TestMinusAFloattToAnInt(t *testing.T) {
	filter := MinusFactory([]core.Value{floatValue(5.11)})
	assert.Equal(t, filter(43, nil).(float64), 37.89)
}

func TestMinusAnIntToAFloat(t *testing.T) {
	filter := MinusFactory([]core.Value{intValue(5)})
	assert.Equal(t, filter(43.2, nil).(float64), 38.2)
}

func TestMinusAnIntToATime(t *testing.T) {
	now := time.Now()
	filter := MinusFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter(now, nil).(time.Time), now.Add(time.Minute*-7))
}

func TestMinusAnIntToAStringAsAnInt(t *testing.T) {
	filter := MinusFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter("33", nil).(int), 26)
}

func TestMinusAnIntToAStringAsAFloat(t *testing.T) {
	filter := MinusFactory([]core.Value{floatValue(2.2)})
	assert.Equal(t, filter("33.11", nil).(float64), 30.91)
}

func TestMinusAnIntToBytesAsAnInt(t *testing.T) {
	filter := MinusFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter([]byte("34"), nil).(int), 27)
}

func TestMinusAnDynamicIntToBytesAsAnInt(t *testing.T) {
	filter := MinusFactory([]core.Value{dynamicValue("fee")})
	assert.Equal(t, filter([]byte("34"), params("fee", 5)).(int), 29)
}

func TestMinusAnDynamicFloatToBytesAsAnInt(t *testing.T) {
	filter := MinusFactory([]core.Value{dynamicValue("fee")})
	assert.Equal(t, filter([]byte("34"), params("fee", 5.1)).(float64), 28.9)
}
