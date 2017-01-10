package filters

import (
	"bytes"
	"context"
	"strings"

	"github.com/acstech/liquid/core"
)

// Creates a upcase filter
func UpcaseFactory(ctx context.Context, parameters []core.Value) core.Filter {
	return Upcase
}

// convert an input string to uppercase
func Upcase(input interface{}, data map[string]interface{}) interface{} {
	switch typed := input.(type) {
	case []byte:
		return bytes.ToUpper(typed)
	case string:
		return strings.ToUpper(typed)
	default:
		return input
	}
}
