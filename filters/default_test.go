package filters

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/karlseguin/gspec"
)

func TestDefaultWithBuiltinValue(t *testing.T) {
	spec := gspec.New(t)
	filter := DefaultFactory(nil, nil)
	spec.Expect(filter(nil, nil).(string)).ToEqual("")
}

func TestDefaultWithValueOnString(t *testing.T) {
	spec := gspec.New(t)
	filter := DefaultFactory(nil, []core.Value{stringValue("d")})
	spec.Expect(filter("", nil).(string)).ToEqual("d")
}

func TestDefaultWithValueOnArray(t *testing.T) {
	spec := gspec.New(t)
	filter := DefaultFactory(nil, []core.Value{stringValue("n/a")})
	spec.Expect(filter([]int{}, nil).(string)).ToEqual("n/a")
}
