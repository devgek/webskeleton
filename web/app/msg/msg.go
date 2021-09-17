package msg

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/devgek/webskeleton/helper/common"
	"github.com/spf13/viper"
)

//MessageLocator the message locator
type MessageLocator struct {
	*viper.Viper
}

var once sync.Once
var theMessageLocator *MessageLocator

func GetMessageLocator() *MessageLocator {
	return theMessageLocator
}

//NewMessageLocator create MessageLocator and load the message file
func NewMessageLocator(messages []byte) *MessageLocator {
	once.Do(func() {
		msgReader := bytes.NewReader(messages)
		ml := &MessageLocator{viper.New()}
		ml.SetConfigType("yaml")
		err := ml.ReadConfig(msgReader)
		common.PanicOnError(err)
		theMessageLocator = ml
	})

	return theMessageLocator
}

//GetMessageF get the formatted message
func (ml *MessageLocator) GetMessageF(msgKey string, a ...interface{}) string {
	format := ml.GetString(msgKey)
	if len(a) == 0 {
		return format
	}
	return fmt.Sprintf(format, a...)
}
