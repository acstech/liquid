// liquid template parser
package liquid

import (
	"context"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/acstech/liquid/core"
)

// A compiled liquid template
type Template struct {
	Code []core.Code
}

func (t *Template) AddCode(code core.Code) {
	t.Code = append(t.Code, code)
}

func (t *Template) AddSibling(tag core.Tag) error {
	return nil
}

func (t *Template) Type() core.TagType {
	return core.ContainerTag
}

func (t *Template) Name() string {
	return "root"
}

// Parse the bytes into a Liquid template
func Parse(ctx context.Context, data []byte, config *core.Configuration) (*Template, error) {
	if config == nil {
		config = defaultConfig
	}
	cache := config.GetCache()
	if cache == nil {
		return buildTemplate(ctx, data, config)
	}
	hasher := sha1.New()
	hasher.Write(data)
	key := fmt.Sprintf("%x", hasher.Sum(nil))

	template := cache.Get(key)
	if template == nil {
		var err error
		template, err = buildTemplate(ctx, data, config)
		if err != nil {
			return nil, err
		}
		cache.Set(key, template)
	}
	return template.(*Template), nil
}

// Parse the string into a liquid template
func ParseString(ctx context.Context, data string, config *core.Configuration) (*Template, error) {
	return Parse(ctx, []byte(data), config)
}

// Turn the contents of the specified file into a liquid template
func ParseFile(ctx context.Context, path string, config *core.Configuration) (*Template, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(ctx, data, config)
}

func (t *Template) Render(writer io.Writer, data map[string]interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}
	t.Execute(writer, data)
}

func (t *Template) Execute(writer io.Writer, data map[string]interface{}) core.ExecuteState {
	for _, code := range t.Code {
		if state := code.Execute(writer, data); state == core.Break {
			return core.Normal
		}
	}
	return core.Normal
}

func buildTemplate(ctx context.Context, data []byte, config *core.Configuration) (*Template, error) {
	parser := core.NewParser(ctx, data)
	template := new(Template)
	if err := extractTokens(parser, template, config); err != nil {
		return nil, err
	}
	return template, nil
}

func extractTokens(parser *core.Parser, container core.Tag, config *core.Configuration) error {
	stack := []core.Tag{container}
	preserveWhiteSpace := config.GetPreserveWhitespace()
	for parser.HasMore() {
		pre, markupType := parser.ToMarkup(preserveWhiteSpace)
		if len(pre) > 0 {
			container.AddCode(newLiteral(pre))
		}
		if markupType == core.OutputMarkup {
			code, err := newOutput(parser)
			if err != nil {
				return err
			}
			if code != nil {
				container.AddCode(code)
			}
		} else if markupType == core.TagMarkup {
			tag, err := newTag(parser, config)
			if err != nil {
				return err
			}
			switch tag.Type() {
			case core.ContainerTag, core.LoopTag:
				container.AddCode(tag)
				container = tag
				stack = append(stack, container)
			case core.EndTag:
				l := len(stack) - 1
				container = stack[l]
				if tag.Name() != container.Name() {
					return parser.Error(fmt.Sprintf("end tag \"end%s\" cannot terminate %q", tag.Name(), container.Name()))
				}
				stack = stack[0:l]
				container = stack[l-1]
				parser.SkipPastTag()
			case core.SiblingTag:
				if err := stack[len(stack)-1].AddSibling(tag); err != nil {
					return err
				}
				container = tag
			case core.StandaloneTag:
				container.AddCode(tag)
			}
		} else {
			break
		}
	}
	return nil
}
