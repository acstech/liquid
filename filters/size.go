package filters

import (
	"context"

	"github.com/acstech/liquid/core"
)

// Creates a size filter
func SizeFactory(ctx context.Context, parameters []core.Value) core.Filter {
	return Size
}

// Gets the size of a string or array
func Size(input interface{}, data map[string]interface{}) interface{} {
	if n, ok := core.ToLength(input); ok {
		return n
	}
	return input
}
