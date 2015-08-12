package tags

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/acstech/liquid/core"
)

// Creates an include tag
func IncludeFactory(p *core.Parser, config *core.Configuration) (core.Tag, error) {
	value, err := p.ReadValue()
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, p.Error("Invalid include value")
	}

	scopeType := p.ReadName()
	var scope core.Value
	if scopeType == "with" || scopeType == "for" {
		scope, err = p.ReadValue()
		if err != nil {
			return nil, err
		}
	} else {
		scopeType = ""
	}

	p.SkipPastTag()
	return &Include{value, config.GetIncludeHandler(), scopeType, scope}, nil
}

type Include struct {
	value     core.Value
	handler   core.IncludeHandler
	scopeType string
	scope     core.Value
}

func (i *Include) AddCode(code core.Code) {
	panic("Addcode should not have been called on a Include")
}

func (i *Include) AddSibling(tag core.Tag) error {
	panic("AddSibling should not have been called on a Include")
}

func (i *Include) Execute(writer io.Writer, data map[string]interface{}) core.ExecuteState {
	template := core.ToString(i.value.Resolve(data))
	_, filename := filepath.Split(template)
	extension := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(extension)]
	contextVariableName := strings.ToLower(name)

	var templateData = make([]map[string]interface{}, 1)

	switch i.scopeType {
	case "with":
		scopedData := i.scope.Resolve(data)
		templateData[0] = toMap(scopedData, contextVariableName)
	case "for":
		scopedData := i.scope.Resolve(data)

		switch typed := scopedData.(type) {
		case []interface{}:
			templateData = make([]map[string]interface{}, len(typed))
			for i, item := range typed {
				templateData[i] = toMap(item, contextVariableName)
			}
		case []map[string]interface{}:
			templateData = make([]map[string]interface{}, len(typed))
			for i, item := range typed {
				templateData[i] = toMap(item, contextVariableName)
			}
		default:
			templateData[0] = toMap(scopedData, contextVariableName)
		}
	default:
		templateData[0] = data
	}

	if i.handler != nil {
		for _, item := range templateData {
			i.handler(template, writer, item)
		}
	}
	return core.Normal
}

func toMap(data interface{}, contextVariableName string) map[string]interface{} {
	if data == nil {
		return nil
	}

	var context map[string]interface{}
	if typed, ok := data.(map[string]interface{}); ok {
		context = typed
	} else {
		context := make(map[string]interface{})
		context[contextVariableName] = typed
	}
	return context
}

func (i *Include) Name() string {
	return "include"
}

func (i *Include) Type() core.TagType {
	return core.StandaloneTag
}
