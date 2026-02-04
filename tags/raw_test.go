package tags

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TEstRawFactory(t *testing.T) {
	parser := newParser(" %} this {{}} {%} is raw {%endraw%}Z")
	tag, err := RawFactory(parser, nil)
	assert.Nil(t, err)
	assert.Equal(t, tag.Name(), "raw")
	assert.Equal(t, parser.Current(), byte('Z'))
}

func TestRawFactoryHandlesUnclosedRaw(t *testing.T) {
	parser := newParser(" %} this is raw {%enccsad%}X")
	tag, err := RawFactory(parser, nil)
	assert.Nil(t, err)
	assert.Equal(t, tag.Name(), "raw")
	assert.Equal(t, parser.HasMore(), false)
}

func TestRawTagExecutes(t *testing.T) {
	parser := newParser(" %} this {{}} {%} is raw {%endraw%}Z")
	tag, _ := RawFactory(parser, nil)

	writer := new(bytes.Buffer)
	tag.Execute(writer, nil)
	assert.Equal(t, writer.String(), " this {{}} {%} is raw ")
}
