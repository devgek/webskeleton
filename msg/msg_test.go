package msg_test

import (
	"github.com/devgek/webskeleton/msg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetString(t *testing.T) {
	messages := []byte("test01: Dieser Text ist bla bla")
	ml := msg.NewMessageLocator(messages)

	theMsg := ml.GetString("test01")
	assert.Equal(t, "Dieser Text ist bla bla", theMsg, "Text was not expected")
}

func TestGetMessageF(t *testing.T) {
	messages := []byte("test01: Dieser Text beinhaltet %s")
	ml := msg.NewMessageLocator(messages)

	theMsg := ml.GetMessageF("test01", "blub")
	assert.Equal(t, "Dieser Text beinhaltet blub", theMsg, "Text was not expected")
}
