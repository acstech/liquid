package liquid

import (
	"github.com/acstech/liquid/core"
	"github.com/karlseguin/bytepool"
)

var (
	defaultConfig = Configure()
	//A Configuration with caching disabled
	NoCache = Configure().Cache(nil)
)

// Entry into the fluent-configuration
func Configure() *core.Configuration {
	c := new(core.Configuration)
	return c.Cache(TemplateCache)
}

// Set's the count and size of the internal bytepool
func SetInternalBuffer(count, size int) {
	core.BytePool = bytepool.New(count, size)
}
