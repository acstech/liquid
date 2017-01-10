package filters

import (
	"context"
	"regexp"

	"github.com/acstech/liquid/core"
)

var stripNewLines = &ReplacePattern{regexp.MustCompile("(\n|\r)"), ""}

func StripNewLinesFactory(ctx context.Context, parameters []core.Value) core.Filter {
	return stripNewLines.Replace
}
