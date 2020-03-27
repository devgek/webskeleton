package web

import (
	"github.com/devgek/webskeleton/config"
	"github.com/stretchr/objx"
)

//AuthCookieName the name of the auth cookie
var AuthCookieName = "webskeleton-auth"

//CookieData ...
type CookieData interface {
	Data() interface{}
}

//NewCookieData ...
func NewCookieData(data config.RequestData) CookieData {
	return &CookieDataImpl{data}
}

//CookieDataImpl ...
type CookieDataImpl struct {
	CData config.RequestData
}

//Data ...
func (c CookieDataImpl) Data() interface{} {
	return c.CData
}

//MSI ...
func (c CookieDataImpl) MSI() map[string]interface{} {
	ctxMap := objx.New(c.Data())

	return objx.New(map[string]interface{}{
		"cookie-data": ctxMap,
	})
}
