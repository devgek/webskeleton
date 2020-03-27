package config

import (
	"net/http"

	"github.com/stretchr/objx"
)

//RequestData the data hold in request context
type RequestData interface {
	UserID() string
	SetUserID(userID string)
	Admin() bool
	SetAdmin(admin bool)
}
type requestData struct {
	userID string
	admin  bool
}

//NewRequestData create RequestData
func NewRequestData() RequestData {
	return &requestData{"", false}
}

func (c requestData) UserID() string {
	return c.userID
}

func (c requestData) Admin() bool {
	return c.admin
}

func (c *requestData) SetUserID(userID string) {
	c.userID = userID
}

func (c *requestData) SetAdmin(admin bool) {
	c.admin = admin
}

func (c requestData) MSI() map[string]interface{} {
	ctxMap := objx.New(map[string]interface{}{
		"userID": c.UserID(),
		"admin":  c.Admin(),
	})
	return objx.New(map[string]interface{}{
		"request-data": ctxMap,
	})
}

//ContextKeyRequestData ...
var ContextKeyRequestData = "request-data"

//FromCookie get RequestData from cookie value
func FromCookie(cookie *http.Cookie) (RequestData, bool) {
	c := objx.MustFromBase64(cookie.Value)

	cData := NewRequestData()
	val := c.Get("cookie-data.request-data.userID")
	if val != nil {
		cData.SetUserID(val.Str())
		val = c.Get("cookie-data.request-data.admin")
		if val != nil {
			cData.SetAdmin(val.Bool())
			return cData, true
		}
	}

	return &requestData{}, false
}
