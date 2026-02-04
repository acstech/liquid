package tags

import (
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func TestCommentFactoryForNormalComment(t *testing.T) {
	parser := newParser(" %} hack {%endcomment%}Z")
	tag, err := CommentFactory(parser, nil)
	assert.Nil(t, err)
	assert.Equal(t, tag.Name(), "comment")
	assert.Equal(t, parser.Current(), byte('Z'))
}

func TestCommentFactoryForNestedComment(t *testing.T) {
	parser := newParser(" %} ha {%comment%} {%if%} ck {%endcomment%} {%  endcomment  %}XZ ")
	tag, err := CommentFactory(parser, nil)
	assert.Nil(t, err)
	assert.Equal(t, tag.Name(), "comment")
	assert.Equal(t, parser.Current(), byte('X'))
}

func TestCommentFactoryHandlesUnclosedComment(t *testing.T) {
	parser := newParser(" %} ouch ")
	tag, err := CommentFactory(parser, nil)
	assert.Nil(t, err)
	assert.Equal(t, tag.Name(), "comment")
	assert.Equal(t, parser.HasMore(), false)
}

func newParser(s string) *core.Parser {
	return core.NewParser([]byte(s))
}
