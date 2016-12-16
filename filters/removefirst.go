package filters

import (
	"context"

	"github.com/acstech/liquid/core"
)

var (
	EmptyValue = &core.StaticStringValue{""}
)

func RemoveFirstFactory(ctx context.Context, parameters []core.Value) core.Filter {
	if len(parameters) != 1 {
		return Noop
	}
	return (&ReplaceFilter{parameters[0], EmptyValue, 1}).Replace
}
