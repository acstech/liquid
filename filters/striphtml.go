package filters

import (
	"context"
	"regexp"

	"github.com/acstech/liquid/core"
)

var stripHtml = &ReplacePattern{regexp.MustCompile("(?i)<script.*?</script>|<!--.*?-->|<style.*?</style>|<.*?>"), ""}

func StripHtmlFactory(ctx context.Context, parameters []core.Value) core.Filter {
	return stripHtml.Replace
}
