package tags

import (
	"io"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/acstech/liquid/core"
)

const includeWithScopeType = "with"
const includeForScopeType = "for"

// Creates an include tag
func IncludeFactory(p *core.Parser, config *core.Configuration) (core.Tag, error) {
	includeName, err := p.ReadValue()
	if err != nil {
		return nil, err
	}
	if includeName == nil {
		return nil, p.Error("Invalid include value")
	}

	var scopeType string
	var scopeKey string
	var scope core.Value
	params := make(map[string]core.Value)

	var next = p.SkipSpaces()
	if next == 'w' || next == 'f' {
		scopeType = p.ReadName()

		if scopeType == includeWithScopeType || scopeType == includeForScopeType {
			includeNameString := core.ToString(includeName.Resolve(nil))
			scopeKey = strings.TrimSuffix(includeNameString, filepath.Ext(includeNameString))

			scope, err = p.ReadValue()
			if err != nil {
				return nil, err
			}
		}

		next = p.SkipSpaces()
	}

	if next == ',' {
		p.Forward()

		for name := p.ReadName(); name != ""; name = p.ReadName() {
			if p.SkipSpaces() == ':' {
				p.Forward()

				params[name], err = p.ReadValue()
				if err != nil {
					return nil, err
				}
			}

			if p.SkipSpaces() == ',' {
				p.Forward()
			}
		}
	}

	p.SkipPastTag()

	return &Include{
		includeName: includeName,
		handler:     config.GetIncludeHandler(),
		scopeType:   scopeType,
		scopeKey:    scopeKey,
		scope:       scope,
		parameters:  params,
	}, nil
}

type Include struct {
	includeName core.Value
	handler     core.IncludeHandler
	scopeType   string
	scopeKey    string
	scope       core.Value
	parameters  map[string]core.Value
}

func (i *Include) AddCode(code core.Code) {
	panic("Addcode should not have been called on a Include")
}

func (i *Include) AddSibling(tag core.Tag) error {
	panic("AddSibling should not have been called on a Include")
}

func (i *Include) LastSibling() core.Tag {
	return nil
}

func (i *Include) Execute(writer io.Writer, data map[string]interface{}) core.ExecuteState {
	if i.handler == nil {
		return core.Normal
	}

	template := core.ToString(i.includeName.Resolve(data))

	// Resolve values for all our parameters
	for key, value := range i.parameters {
		data[key] = value.Resolve(data)
	}

	if i.scope != nil {
		scope := i.scope.Resolve(data)

		if i.scopeType == includeForScopeType {
			// Resolve returns a byte array when resolved data is nil that we can't do
			// anything with. Bail so we dont just iterate through an array of bytes.
			if _, byteOk := scope.([]byte); !byteOk {

				switch reflect.TypeOf(scope).Kind() {
				case reflect.Slice:
					s := reflect.ValueOf(scope)

					for idx := 0; idx < s.Len(); idx++ {
						data[i.scopeKey] = s.Index(idx).Interface()
						i.handler(template, writer, data)
					}
				}

				return core.Normal
			}
		}

		data[i.scopeKey] = scope
	}

	i.handler(template, writer, data)

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
