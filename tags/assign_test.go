package tags

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/acstech/liquid/filters"
	"github.com/stretchr/testify/assert"
)

func init() {
	core.RegisterFilter("minus", filters.MinusFactory)
}

func TestAssignForAStaticAndNoFilters(t *testing.T) {
	parser := newParser(" var = 'abc123'%}B")
	assertStringAssign(t, parser, "var", "abc123")
	assert.Equal(t, parser.Current(), byte('B'))
}

func TestAssignForAStaticWithFilters(t *testing.T) {
	parser := newParser("sale  =  213  |minus: 4  %}o")
	assertIntAssign(t, parser, "sale", 209)
	assert.Equal(t, parser.Current(), byte('o'))
}

func TestAssignForAVariableWithFilters(t *testing.T) {
	parser := newParser("sale = price  |minus: 9  %}o")
	assertIntAssign(t, parser, "sale", 91)
}

func assertStringAssign(t *testing.T, parser *core.Parser, variableName, value string) {
	tag, err := AssignFactory(parser, nil)
	assert.Nil(t, err)
	assert.Equal(t, tag.Name(), "assign")
	m := make(map[string]interface{})
	tag.Execute(nil, m)
	assert.Equal(t, m[variableName].(string), value)
}

func assertIntAssign(t *testing.T, parser *core.Parser, variableName string, value int) {
	tag, err := AssignFactory(parser, nil)
	assert.Nil(t, err)
	assert.Equal(t, tag.Name(), "assign")
	m := map[string]interface{}{
		"price": 100,
	}
	tag.Execute(nil, m)
	assert.Equal(t, m[variableName].(int), value)
}
