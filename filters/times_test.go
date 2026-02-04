package filters

import (
	"strings"
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestTimesAnIntToAnInt(t *testing.T) {
	filter := TimesFactory([]core.Value{intValue(5)})
	assert.Equal(t, filter(43, nil).(int), 215)
}

func TestTimesAnIntToAFloat(t *testing.T) {
	filter := TimesFactory([]core.Value{intValue(2)})
	assert.Equal(t, filter(43.3, nil).(float64), 86.6)
}

func TestTimesAnIntToAStringAsAnInt(t *testing.T) {
	filter := TimesFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter("33", nil).(int), 231)
}

func TestTimesAnIntToBytesAsAnInt(t *testing.T) {
	filter := TimesFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter([]byte("34"), nil).(int), 238)
}

func TestTimesAnIntToAStringAsAString(t *testing.T) {
	filter := TimesFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter("abc", nil).(string), "abc")
}

func TestTimesAnIntToBytesAsAString(t *testing.T) {
	filter := TimesFactory([]core.Value{intValue(8)})
	assert.Equal(t, filter([]byte("abb"), nil).(string), "abb")
}

func TestTimesAFloatToAnInt(t *testing.T) {
	filter := TimesFactory([]core.Value{floatValue(1.10)})
	assert.Equal(t, filter(43, nil).(float64), 47.300000000000004)
}

func TestTimesAFloatToAFloat(t *testing.T) {
	filter := TimesFactory([]core.Value{floatValue(5.3)})
	assert.Equal(t, filter(43.3, nil).(float64), 229.48999999999998)
}

func TestTimesAFloatToAStringAsAnInt(t *testing.T) {
	filter := TimesFactory([]core.Value{floatValue(7.11)})
	assert.Equal(t, filter("33", nil).(float64), 234.63000000000002)
}

func TestTimesADynamicIntValue(t *testing.T) {
	filter := TimesFactory([]core.Value{dynamicValue("count")})
	assert.Equal(t, filter("33", params("count", 112)).(int), 3696)
}

func TestTimesADynamicFloatValue(t *testing.T) {
	filter := TimesFactory([]core.Value{dynamicValue("count")})
	assert.Equal(t, filter("12", params("count", 44.2)).(float64), 530.4000000000001)
}

func TestTimesDynamicNoop(t *testing.T) {
	filter := TimesFactory([]core.Value{dynamicValue("count")})
	assert.Equal(t, filter("12", params("count", "22")).(string), "12")
}

func stringValue(s string) core.Value {
	return &core.StaticStringValue{s}
}

func intValue(n int) core.Value {
	return &core.StaticIntValue{n}
}

func floatValue(f float64) core.Value {
	return &core.StaticFloatValue{f}
}

func dynamicValue(s string) core.Value {
	return core.NewDynamicValue(strings.Split(s, "."))
}

func params(values ...interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(values); i += 2 {
		m[values[i].(string)] = values[i+1]
	}
	return m
}
