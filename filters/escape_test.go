package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
)

func TestEscapesAString(t *testing.T) {
	spec := gspec.New(t)
	filter := EscapeFactory(nil, nil)
	spec.Expect(filter("<script>hack</script>", nil).(string)).ToEqual("&lt;script&gt;hack&lt;/script&gt;")
}
