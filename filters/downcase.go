package filters

import (
	"bytes"
	"context"
	"strings"

	"github.com/acstech/liquid/core"
)

// Creates a downcase filter
func DowncaseFactory(ctx context.Context, parameters []core.Value) core.Filter {
	return Downcase
}

// convert an input string to lowercase
func Downcase(input interface{}, data map[string]interface{}) interface{} {
	switch typed := input.(type) {
	case []byte:
		return bytes.ToLower(typed)
	case string:
		return strings.ToLower(typed)
	default:
		return input
	}
}
