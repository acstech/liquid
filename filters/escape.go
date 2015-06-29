package filters

import (
	"html"

	"github.com/acstech/liquid/core"
)

// Creates an escape filter
func EscapeFactory(parameters []core.Value) core.Filter {
	return Escape
}

// Escapes html
func Escape(input interface{}, data map[string]interface{}) interface{} {
	switch typed := input.(type) {
	case string:
		return html.EscapeString(typed)
	case []byte:
		return html.EscapeString(string(typed))
	default:
		return input
	}
}
