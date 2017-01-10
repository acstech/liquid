package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
)

func TestUpcasesAString(t *testing.T) {
	spec := gspec.New(t)
	filter := UpcaseFactory(nil, nil)
	spec.Expect(filter("dbz", nil).(string)).ToEqual("DBZ")
}

func TestUpcasesBytes(t *testing.T) {
	spec := gspec.New(t)
	filter := UpcaseFactory(nil, nil)
	spec.Expect(string(filter([]byte("dbz"), nil).([]byte))).ToEqual("DBZ")
}

func TestUpcasesPassThroughOnInvalidType(t *testing.T) {
	spec := gspec.New(t)
	filter := UpcaseFactory(nil, nil)
	spec.Expect(filter(123, nil).(int)).ToEqual(123)
}
