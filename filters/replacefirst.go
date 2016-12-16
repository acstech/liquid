package filters

import (
	"context"

	"github.com/acstech/liquid/core"
)

func ReplaceFirstFactory(ctx context.Context, parameters []core.Value) core.Filter {
	if len(parameters) != 2 {
		return Noop
	}
	return (&ReplaceFilter{parameters[0], parameters[1], 1}).Replace
}
