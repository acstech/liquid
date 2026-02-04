package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscapesAString(t *testing.T) {
	filter := EscapeFactory(nil)
	assert.Equal(t, filter("<script>hack</script>", nil).(string), "&lt;script&gt;hack&lt;/script&gt;")
}
