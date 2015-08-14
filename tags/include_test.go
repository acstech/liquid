package tags

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/acstech/liquid/core"
	"github.com/karlseguin/gspec"
)

func TestIncludeFactory(t *testing.T) {
	spec := gspec.New(t)
	parser := newParser("'test' %}Z")
	tag, err := IncludeFactory(parser, new(core.Configuration))
	spec.Expect(err).ToBeNil()
	spec.Expect(tag.Name()).ToEqual("include")
	spec.Expect(parser.Current()).ToEqual(byte('Z'))
}

func TestIncludeTagExecutes(t *testing.T) {
	spec := gspec.New(t)
	parser := newParser("'test' %}")
	testData := make(map[string]interface{})
	testData["key"] = "test"

	config := new(core.Configuration).IncludeHandler(func(name string, writer io.Writer, data map[string]interface{}) {
		writer.Write([]byte(fmt.Sprintf("%v", data["key"])))
	})
	tag, err := IncludeFactory(parser, config)

	spec.Expect(err).ToBeNil()

	writer := new(bytes.Buffer)
	tag.Execute(writer, testData)
	spec.Expect(writer.String()).ToEqual(testData["key"])
}

func TestIncludeTagWithExecutes(t *testing.T) {
	spec := gspec.New(t)
	parser := newParser("'test' with context %}")
	testData := make(map[string]interface{})
	testContext := make(map[string]interface{})
	testContext["key"] = "test"
	testData["context"] = testContext

	config := new(core.Configuration).IncludeHandler(func(name string, writer io.Writer, data map[string]interface{}) {
		writer.Write([]byte(fmt.Sprintf("%v", data["key"])))
	})
	tag, err := IncludeFactory(parser, config)

	spec.Expect(err).ToBeNil()

	writer := new(bytes.Buffer)
	tag.Execute(writer, testData)
	spec.Expect(writer.String()).ToEqual(testContext["key"])
}

func TestIncludeTagForExecutes(t *testing.T) {
	spec := gspec.New(t)
	parser := newParser("'test' for context %}")

	testData := make(map[string]interface{})
	testArray := make([]string, 3)
	testArray[0] = "1"
	testArray[1] = "2"
	testArray[2] = "3"
	testData["context"] = testArray

	config := new(core.Configuration).IncludeHandler(func(name string, writer io.Writer, data map[string]interface{}) {
		writer.Write([]byte(fmt.Sprintf("%v", data["test"])))
	})
	tag, err := IncludeFactory(parser, config)

	spec.Expect(err).ToBeNil()

	writer := new(bytes.Buffer)
	tag.Execute(writer, testData)
	spec.Expect(writer.String()).ToEqual("123")
}
