package msg

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/devgek/webskeleton/helper"
	"github.com/spf13/viper"
)

//MessageLocator the message locator
type MessageLocator struct {
	*viper.Viper
}

var once sync.Once

//Messages singleton instance for the app messages
// var Messages *MessageLocator

//NewMessageLocator create MessageLocator and load the message file
func NewMessageLocator(messages []byte) *MessageLocator {
	msgReader := bytes.NewReader(messages)
	ml := &MessageLocator{viper.New()}
	ml.SetConfigType("yaml")
	err := ml.ReadConfig(msgReader)
	helper.PanicOnError(err)
	return ml
}

//GetMessageF get the formatted message
func (ml *MessageLocator) GetMessageF(msgKey string, a ...interface{}) string {
	format := ml.GetString(msgKey)
	if len(a) == 0 {
		return format
	}
	return fmt.Sprintf(format, a...)
}
