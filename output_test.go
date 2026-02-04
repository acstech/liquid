package liquid

import (
	"strconv"
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestOutputHandlesEmptyOutput(t *testing.T) {
	output, err := newOutput(core.NewParser([]byte("{{}}")))
	assert.Nil(t, output)
	assert.Nil(t, err)
}

func TestOutputHandlesSpaceOnlyOutput(t *testing.T) {
	output, err := newOutput(core.NewParser([]byte("{{   }}")))
	assert.Nil(t, output)
	assert.Nil(t, err)
}

func TestOutputExtractsASimpleStatic(t *testing.T) {
	output, _ := newOutput(core.NewParser([]byte("{{  'over 9000'}}")))
	assertRender(t, output, nil, "over 9000")
}

func TestOutputExtractsAComplexStatic(t *testing.T) {
	output, _ := newOutput(core.NewParser([]byte("{{'it\\'s over \\9000'}}")))
	assertRender(t, output, nil, "it's over \\9000")
}

func TestOutputExtractsAStaticWithAnEndingQuote(t *testing.T) {
	output, _ := newOutput(core.NewParser([]byte("{{'it\\''}}")))
	assertRender(t, output, nil, "it'")
}

func TestOutputExtractionGivesErrorForUnclosedStatic(t *testing.T) {
	output, err := newOutput(core.NewParser([]byte("{{ 'failure }}")))
	assert.Nil(t, output)
	assert.Equal(t, err.Error(), `Invalid value, a single quote might be missing ("{{ 'failure }}" - line 1)`, "error message should match")
}

func TestOutputNoFiltersForStatic(t *testing.T) {
	output, _ := newOutput(core.NewParser([]byte("{{'fun'}}")))
	assert.Equal(t, len(output.(*Output).Filters), 0, "no filters should be found")
}

func TestOutputGeneratesErrorOnUnknownFilter(t *testing.T) {
	_, err := newOutput(core.NewParser([]byte("{{'fun' | unknown }}")))
	assert.Equal(t, err.Error(), `Unknown filter "unknown" ("{{'fun' | unknown }}" - line 1)`, "error message should match")
}

func TestOutputGeneratesErrorOnInvalidParameter(t *testing.T) {
	_, err := newOutput(core.NewParser([]byte("{{'fun' | debug: 'missing }}")))
	assert.Equal(t, err.Error(), `Invalid value, a single quote might be missing ("{{'fun' | debug: 'missing }}" - line 1)`, "error message should match")
}

func TestOutputWithASingleFilter(t *testing.T) {
	output, _ := newOutput(core.NewParser([]byte("{{'fun' | debug }}")))
	assertFilters(t, output, "debug(0)")
}

func TestOutputWithMultipleFilters(t *testing.T) {
	output, _ := newOutput(core.NewParser([]byte("{{'fun' | debug | debug}}")))
	assertFilters(t, output, "debug(0)", "debug(1)")
}

func TestOutputWithMultipleFiltersHavingParameters(t *testing.T) {
	output, err := newOutput(core.NewParser([]byte("{{'fun' | debug:1,2 | debug:'test' | debug : 'test' , 5}}")))
	assert.Nil(t, err)
	assertFilters(t, output, "debug(0, 1, 2)", "debug(1, test)", "debug(2, test, 5)")
}

func TestOutputWithAnEscapeParameter(t *testing.T) {
	output, err := newOutput(core.NewParser([]byte("{{'fun' | debug: 'te\\'st'}}")))
	assert.Nil(t, err)
	assertFilters(t, output, "debug(0, te'st)")
}

func assertFilters(t *testing.T, output core.Code, expected ...string) {
	filters := output.(*Output).Filters
	assert.Equal(t, len(filters), len(expected))
	for index, filter := range filters {
		actual := string(filter(strconv.Itoa(index), nil).([]byte))
		assert.Equal(t, actual, expected[index])
	}
}
