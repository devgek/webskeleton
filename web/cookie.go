package web

import (
	"github.com/stretchr/objx"
)

//AuthCookieName the name of the auth cookie
var AuthCookieName = "webskeleton-auth"

//CookieData ...
type CookieData interface {
	Data() interface{}
}

//NewCookieData ...
func NewCookieData(data ContextData) CookieData {
	return &CookieDataImpl{data}
}

//CookieDataImpl ...
type CookieDataImpl struct {
	CData ContextData
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
