package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestDivideByAnIntToAnInt(t *testing.T) {
	filter := DivideByFactory([]core.Value{intValue(5)})
	assert.Equal(t, filter(43, nil).(float64), 8.6)
}

func TestDivideByAnIntToAFloat(t *testing.T) {
	filter := DivideByFactory([]core.Value{intValue(2)})
	assert.Equal(t, filter(43.3, nil).(float64), 21.65)
}

func TestDivideByAnIntToAStringAsAnInt(t *testing.T) {
	filter := DivideByFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter("33", nil).(float64), 4.714285714285714)
}

func TestDivideByAnIntToBytesAsAnInt(t *testing.T) {
	filter := DivideByFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter([]byte("34"), nil).(float64), 4.857142857142857)
}

func TestDivideByAnIntToAStringAsAString(t *testing.T) {
	filter := DivideByFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter("abc", nil).(string), "abc")
}

func TestDivideByAnIntToBytesAsAString(t *testing.T) {
	filter := DivideByFactory([]core.Value{intValue(8)})
	assert.Equal(t, filter([]byte("abb"), nil).(string), "abb")
}

func TestDivideByAFloatToAnInt(t *testing.T) {
	filter := DivideByFactory([]core.Value{floatValue(1.10)})
	assert.Equal(t, filter(43, nil).(float64), 39.090909090909086)
}

func TestDivideByAFloatToAFloat(t *testing.T) {
	filter := DivideByFactory([]core.Value{floatValue(5.3)})
	assert.Equal(t, filter(43.3, nil).(float64), 8.169811320754716)
}

func TestDivideByAFloatToAStringAsAnInt(t *testing.T) {
	filter := DivideByFactory([]core.Value{floatValue(7.11)})
	assert.Equal(t, filter("33", nil).(float64), 4.641350210970464)
}

func TestDivideByADynamicIntValue(t *testing.T) {
	filter := DivideByFactory([]core.Value{dynamicValue("count")})
	assert.Equal(t, filter("33", params("count", 112)).(float64), 0.29464285714285715)
}

func TestDivideByADynamicFloatValue(t *testing.T) {
	filter := DivideByFactory([]core.Value{dynamicValue("count")})
	assert.Equal(t, filter("12", params("count", 44.2)).(float64), 0.27149321266968324)
}

func TestDivideByDynamicNoop(t *testing.T) {
	filter := DivideByFactory([]core.Value{dynamicValue("count")})
	assert.Equal(t, filter("12", params("count", "22")).(string), "12")
}
