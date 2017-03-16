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

	// Resolve values for all our parameters
	for key, value := range i.parameters {
		data[key] = value.Resolve(data)
	}

	if i.scopeType == includeForScopeType {
		if c, ok := i.executeFor(writer, data); ok {
			return c
		}
	}

	template := core.ToString(i.includeName.Resolve(data))

	if i.scope != nil {
		data[i.scopeKey] = i.scope.Resolve(data)
	}

	i.handler(template, writer, data)

	return core.Normal
}

func (i *Include) Name() string {
	return "include"
}

func (i *Include) Type() core.TagType {
	return core.StandaloneTag
}

// executeFor returns false in any case that does not execute the template
func (i *Include) executeFor(writer io.Writer, data map[string]interface{}) (core.ExecuteState, bool) {

	if i.scope == nil {
		return core.Normal, false
	}

	scope := i.scope.Resolve(data)

	// Resolve returns a byte array when resolved data is nil that we can't do
	// anything with. Bail so we dont just iterate through an array of bytes.
	if _, byteOk := scope.([]byte); byteOk {
		return core.Normal, false
	}

	template := core.ToString(i.includeName.Resolve(data))

	switch reflect.TypeOf(scope).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(scope)
		for idx := 0; idx < s.Len(); idx++ {
			data[i.scopeKey] = s.Index(idx).Interface()
			i.handler(template, writer, data)
		}
	default:
		return core.Normal, false
	}

	return core.Normal, true
}
