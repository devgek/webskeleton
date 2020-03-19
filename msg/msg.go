package msg

import (
	"github.com/devgek/webskeleton/helper"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

//MessageLocator the message locator
type MessageLocator struct {
	*viper.Viper
}

//Messages the app messages
var Messages = NewMessageLocator()

//NewMessageLocator create MessageLocator and load the message file
func NewMessageLocator() *MessageLocator {
	runtime.Caller()
	currDir, err := os.Getwd()
	helper.PanicOnError(err)
	ml := &MessageLocator{viper.New()}
	//load locale specific message file, if not default
	// Messages.SetConfigFile(filepath.Join(currDir, "msg", "messages-en.yaml"))
	ml.SetConfigFile(filepath.Join(currDir, "msg", "messages.yaml"))
	err = ml.ReadInConfig()
	helper.PanicOnError(err)

	return ml
}
