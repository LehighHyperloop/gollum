package format

import (
	"testing"

	"github.com/trivago/gollum/core"
	"github.com/trivago/tgo/ttesting"
)

func TestFormatterSerialize(t *testing.T) {
	expect := ttesting.NewExpect(t)

	config := core.NewPluginConfig("", "format.Serialize")

	plugin, err := core.NewPlugin(config)
	expect.NoError(err)

	formatter, casted := plugin.(*Serialize)
	expect.True(casted)

	msg := core.NewMessage(nil, []byte("foo bar"),
		0, core.InvalidStreamID)

	err = formatter.ApplyFormatter(msg)
	expect.NoError(err)

	controlMsg, err := core.DeserializeMessage(msg.Data())
	expect.NoError(err)

	expect.Equal("foo bar", string(controlMsg.Data()))
}