package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestModuloAnIntToAnInt(t *testing.T) {
	filter := ModuloFactory([]core.Value{intValue(5)})
	assert.Equal(t, filter(43, nil).(int), 3)
}

func TestModuloAnFloatToAnInt(t *testing.T) {
	filter := ModuloFactory([]core.Value{floatValue(5.2)})
	assert.Equal(t, filter(43, nil).(int), 3)
}

func TestModuloAnIntToAStringAsAnInt(t *testing.T) {
	filter := ModuloFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter("33", nil).(int), 5)
}

func TestModuloAnIntToBytesAsAnInt(t *testing.T) {
	filter := ModuloFactory([]core.Value{intValue(7)})
	assert.Equal(t, filter([]byte("34"), nil).(int), 6)
}

func TestModuloAnDynamicIntToBytesAsAnInt(t *testing.T) {
	filter := ModuloFactory([]core.Value{dynamicValue("fee")})
	assert.Equal(t, filter([]byte("34"), params("fee", 5)).(int), 4)
}
