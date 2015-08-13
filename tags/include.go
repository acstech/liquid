package tags

import (
	"fmt"
	"io"
	"path/filepath"
	"reflect"
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

		// Resolve returns a byte array when resolved data is nil that we can't do
		// anything with. Bail so we dont just iterate through an array of bytes.
		if _, ok := scopedData.([]byte); ok {
			return core.Normal
		}

		switch reflect.TypeOf(scopedData).Kind() {
		case reflect.Slice:
			// Use reflection to iterate over ANY kind of slice (except []byte - see above)
			slice := reflect.ValueOf(scopedData)
			templateData = make([]map[string]interface{}, slice.Len())
			for i := 0; i < slice.Len(); i++ {
				templateData[i] = toMap(slice.Index(i).Interface(), contextVariableName)
			}
		default:
			fmt.Printf("Scoped Data: %+v\n", scopedData)
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

	if typed, ok := data.(map[string]interface{}); ok {
		return typed
	}

	context := make(map[string]interface{})
	context[contextVariableName] = data
	return context
}

func (i *Include) Name() string {
	return "include"
}

func (i *Include) Type() core.TagType {
	return core.StandaloneTag
}
