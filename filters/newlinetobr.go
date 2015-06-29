package filters

import (
	"regexp"

	"github.com/acstech/liquid/core"
)

var newLinesToBr = &ReplacePattern{regexp.MustCompile("(\n\r|\n|\r)"), "<br />\n"}

func NewLineToBrFactory(parameters []core.Value) core.Filter {
	return newLinesToBr.Replace
}
