package msg_test

import (
	"github.com/devgek/webskeleton/web/app/msg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	messages := []byte("test01: Dieser Text ist bla bla")
	ml := msg.NewMessageLocator(messages)

	theMsg := ml.GetString("test01")
	assert.Equal(t, "Dieser Text ist bla bla", theMsg, "Text was not expected")
}

func TestGetMessageF(t *testing.T) {
	messages := []byte("test02: Dieser Texxt beinhaltet %s")
	ml := msg.NewMessageLocator(messages)

	theMsg := ml.GetMessageF("test02", "blub")
	assert.Equal(t, "Dieser Texxt beinhaltet blub", theMsg, "Text was not expected")
}
