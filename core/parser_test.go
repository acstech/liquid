package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserToMarkupWhenTheresNoMarkup(t *testing.T) {
	parser := newParser("hello world")
	pre, mt := parser.ToMarkup(false)
	assert.Equal(t, mt, NoMarkup)
	assert.Equal(t, string(pre), "hello world")
	assert.Equal(t, parser.HasMore(), false)
}

func TestParserToMarkupWhenThereIsAnOutputMarkup(t *testing.T) {
	parser := newParser("hello world {{ hello }}")
	pre, mt := parser.ToMarkup(false)
	assert.Equal(t, mt, OutputMarkup)
	assert.Equal(t, string(pre), "hello world ")
	assert.Equal(t, parser.HasMore(), true)
	assert.Equal(t, parser.Position, 12)
}

func TestParserToMarkupWhenThereIsATagMarkup(t *testing.T) {
	parser := newParser("hello world {% hello %}")
	pre, mt := parser.ToMarkup(false)
	assert.Equal(t, mt, TagMarkup)
	assert.Equal(t, string(pre), "hello world ")
	assert.Equal(t, parser.HasMore(), true)
	assert.Equal(t, parser.Position, 12)
}

func TestParserSkipsSpacesWhenThereAreNoSpaces(t *testing.T) {
	parser := newParser("hello")
	parser.SkipSpaces()
	assert.Equal(t, parser.Position, 0)
}

func TestParserSkipsSpacesWhenThereAreSpaces(t *testing.T) {
	parser := newParser("  hello")
	parser.SkipSpaces()
	assert.Equal(t, parser.Position, 2)
}

func TestParserParsesAnEmptyValue(t *testing.T) {
	parser := newParser("  ")
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Nil(t, value)
	assert.Equal(t, parser.Position, 2)
}

func TestParserParsesAnEmptyValue2(t *testing.T) {
	parser := newParser("  }}")
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Nil(t, value)
	assert.Equal(t, parser.Position, 2)
}

func TestParserParsesAStaticValue(t *testing.T) {
	parser := newParser(` 'hel"lo' `)
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Equal(t, value.Resolve(nil).(string), `hel"lo`)
	assert.Equal(t, parser.Position, 9)
}

func TestParserParsesAStaticValueWithDoubleQuotes(t *testing.T) {
	parser := newParser(` "hello'" `)
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Equal(t, value.Resolve(nil).(string), "hello'")
	assert.Equal(t, parser.Position, 9)
}

func TestParserParsesTrueBoolean(t *testing.T) {
	parser := newParser(" true ")
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Equal(t, value.Resolve(nil).(bool), true)
	assert.Equal(t, parser.Position, 5)
}

func TestParserParsesFalseBoolean(t *testing.T) {
	parser := newParser(" false ")
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Equal(t, value.Resolve(nil).(bool), false)
	assert.Equal(t, parser.Position, 6)
}

func TestParserParsesEmpty(t *testing.T) {
	parser := newParser(" empty ")
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Equal(t, value.Resolve(nil).(string), "liquid:empty")
	assert.Equal(t, parser.Position, 6)
}

func TestParserParsesAnInteger(t *testing.T) {
	parser := newParser(" 938 ")
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Equal(t, value.Resolve(nil).(int), 938)
	assert.Equal(t, parser.Position, 4)
}

func TestParserParsesANegativeInteger(t *testing.T) {
	parser := newParser(" -331 ")
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Equal(t, value.Resolve(nil).(int), -331)
	assert.Equal(t, parser.Position, 5)
}

func TestParserParsesAFloat(t *testing.T) {
	parser := newParser(" 9000.1 ")
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Equal(t, value.Resolve(nil).(float64), 9000.1)
	assert.Equal(t, parser.Position, 7)
}

func TestParserParsesANegativeFloat(t *testing.T) {
	parser := newParser(" -331.89 ")
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Equal(t, value.Resolve(nil).(float64), -331.89)
	assert.Equal(t, parser.Position, 8)
}

func TestParserParsesAStaticValueWithEscapedQuote(t *testing.T) {
	parser := newParser(" 'hello \\'You\\' ' ")
	value, err := parser.ReadValue()
	assert.Nil(t, err)
	assert.Equal(t, value.Resolve(nil).(string), "hello 'You' ")
	assert.Equal(t, parser.Position, 17)
}

func TestParserParsesAStaticWithMissingClosingQuote(t *testing.T) {
	parser := newParser(" 'hello ")
	_, err := parser.ReadValue()
	assert.Equal(t, err.Error(), `Invalid value, a single quote might be missing (" 'hello " - line 1)`)
}

func TestParserParsesASingleLevelDynamicValue(t *testing.T) {
	parser := newParser(" user ")
	v, err := parser.ReadValue()
	values := v.(*DynamicValue)
	assert.Nil(t, err)
	assert.Equal(t, len(values.Fields), 1)
	assert.Equal(t, values.Fields[0], "user")
	assert.Equal(t, parser.Position, 5)
}

func TestParserParsesAMultiLevelDynamicValue(t *testing.T) {
	parser := newParser(" user.NaMe.first}}")
	v, err := parser.ReadValue()
	values := v.(*DynamicValue)
	assert.Nil(t, err)
	assert.Equal(t, len(values.Fields), 3)
	assert.Equal(t, values.Fields[0], "user")
	assert.Equal(t, values.Fields[1], "name")
	assert.Equal(t, values.Fields[2], "first")
	assert.Equal(t, parser.Position, 16)
}

func TestParserReadsAnEmptyName1(t *testing.T) {
	parser := newParser("  ")
	assert.Equal(t, parser.ReadName(), "")
	assert.Equal(t, parser.Position, 2)
}

func TestParserReadsAnEmptyName2(t *testing.T) {
	parser := newParser("   }}")
	assert.Equal(t, parser.ReadName(), "")
	assert.Equal(t, parser.Position, 3)
}

func TestParserReadsAnEmptyName3(t *testing.T) {
	parser := newParser("%}")
	assert.Equal(t, parser.ReadName(), "")
	assert.Equal(t, parser.Position, 0)
}

func TestParserReadsAnEmptyName4(t *testing.T) {
	parser := newParser(" |")
	assert.Equal(t, parser.ReadName(), "")
	assert.Equal(t, parser.Position, 1)
}

func TestParserReadsAName(t *testing.T) {
	parser := newParser(" spice }}")
	assert.Equal(t, parser.ReadName(), "spice")
	assert.Equal(t, parser.Position, 6)
}

func TestParserReadsEmptyParameters(t *testing.T) {
	parser := newParser(" }}")
	values, err := parser.ReadParameters()
	assert.Nil(t, err)
	assert.Equal(t, len(values), 0)
	assert.Equal(t, parser.Position, 1)
}

func TestParserReadsASingleParameter(t *testing.T) {
	parser := newParser(" 'hello'")
	values, err := parser.ReadParameters()
	assert.Nil(t, err)
	assert.Equal(t, len(values), 1)
	assert.Equal(t, values[0].Resolve(nil).(string), "hello")
	assert.Equal(t, parser.Position, 8)
}

func TestParserReadsMultipleParameters(t *testing.T) {
	parser := newParser(" 'hello' , 123 ")
	values, err := parser.ReadParameters()
	assert.Nil(t, err)
	assert.Equal(t, len(values), 2)
	assert.Equal(t, values[0].Resolve(nil).(string), "hello")
	assert.Equal(t, values[1].Resolve(nil).(int), 123)
	assert.Equal(t, parser.Position, 15)
}

func TestParserReadsAUnaryCondition(t *testing.T) {
	parser := newParser(" true %}")
	group, err := parser.ReadConditionGroup()
	assert.Nil(t, err)
	assertParsedConditionGroup(t, group, true, Unary, nil)
}

func TestParserReadsMultipleUnaryConditions(t *testing.T) {
	parser := newParser(" true and false%}")
	group, err := parser.ReadConditionGroup()
	assert.Nil(t, err)
	assertParsedConditionGroup(t, group, true, Unary, nil, AND, false, Unary, nil)
}

func TestParserReadsSingleCondition(t *testing.T) {
	parser := newParser(" true == 123   %}")
	group, err := parser.ReadConditionGroup()
	assert.Nil(t, err)
	assertParsedConditionGroup(t, group, true, Equals, 123)
}

func TestParserReadsContainsCondition(t *testing.T) {
	parser := newParser(" 'xyz'   contains   true%}")
	group, err := parser.ReadConditionGroup()
	assert.Nil(t, err)
	assertParsedConditionGroup(t, group, "xyz", Contains, true)
}

func TestParserReadsMultipleComplexConditions(t *testing.T) {
	parser := newParser(" 'xyz'   contains   true or true and 123 > 445%}")
	group, err := parser.ReadConditionGroup()
	assert.Nil(t, err)
	assertParsedConditionGroup(t, group, "xyz", Contains, true, OR, true, Unary, nil, AND, 123, GreaterThan, 445)
}

func TestParserReadsASinglePartial(t *testing.T) {
	parser := newParser(" true %}")
	group, err := parser.ReadPartialCondition()
	assert.Nil(t, err)
	assertParsedConditionGroup(t, group, true, UnknownComparator, nil)
}

func TestParserReadsMultiplePartials(t *testing.T) {
	parser := newParser(" 1 or 2%}")
	group, err := parser.ReadPartialCondition()
	assert.Nil(t, err)
	assertParsedConditionGroup(t, group, 1, UnknownComparator, nil, OR, 2, UnknownComparator, nil)
}

func newParser(s string) *Parser {
	return NewParser([]byte(s))
}

func assertParsedConditionGroup(t *testing.T, group Verifiable, data ...interface{}) {
	for i := 0; i < len(data); i += 4 {
		actual := group.(*ConditionGroup).conditions[i%3]
		if s, ok := data[i].(string); ok {
			assert.Equal(t, actual.left.ResolveWithNil(nil).(string), s)
		} else {
			assert.Equal(t, actual.left.ResolveWithNil(nil), data[i])
		}
		assert.Equal(t, actual.operator, data[i+1])
		if data[i+2] == nil {
			assert.Nil(t, actual.right)
		} else {
			assert.Equal(t, actual.right.ResolveWithNil(nil), data[i+2])
		}
		if i != len(data)-3 {
			logical := group.(*ConditionGroup).joins[i%3]
			assert.Equal(t, logical, data[i+3])
		}
	}
}
