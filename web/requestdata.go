package web

import (
	"context"
	"net/http"

	"github.com/stretchr/objx"
)

//ContextData the data hold in request context
type ContextData interface {
	UserID() string
	SetUserID(userID string)
	Host() string
	SetHost(host string)
	Admin() bool
	SetAdmin(admin bool)
}
type contextData struct {
	userID string
	host   string
	admin  bool
}

//NewContextData create ContextData
func NewContextData() ContextData {
	return &contextData{"", "", false}
}

func (c contextData) UserID() string {
	return c.userID
}

func (c contextData) Host() string {
	return c.userID
}

func (c contextData) Admin() bool {
	return c.admin
}

func (c *contextData) SetUserID(userID string) {
	c.userID = userID
}

func (c *contextData) SetHost(host string) {
	c.host = host
}

func (c *contextData) SetAdmin(admin bool) {
	c.admin = admin
}

func (c contextData) MSI() map[string]interface{} {
	ctxMap := objx.New(map[string]interface{}{
		"userID": c.UserID(),
		"host":   c.Host(),
		"admin":  c.Admin(),
	})
	return objx.New(map[string]interface{}{
		"context-data": ctxMap,
	})
}

type contextKey struct {
	name string
}

//ContextKeyContextData ...
var ContextKeyContextData = &contextKey{"context-data"}

//ToContext set ContextData to context
func ToContext(ctx context.Context, cData ContextData) context.Context {
	return context.WithValue(ctx, ContextKeyContextData, cData)
}

//FromContext get ContextData out of context
func FromContext(ctx context.Context) ContextData {
	val := ctx.Value(ContextKeyContextData)
	if val == nil {
		return NewContextData()
	}
	return val.(ContextData)
}

//FromCookie get ContextData from cookie value
func FromCookie(cookie *http.Cookie) (ContextData, bool) {
	c := objx.MustFromBase64(cookie.Value)

	cData := NewContextData()
	val := c.Get("cookie-data.context-data.userID")
	if val != nil {
		cData.SetUserID(val.Str())
		val = c.Get("cookie-data.context-data.host")
		if val != nil {
			cData.SetHost(val.Str())
			val = c.Get("cookie-data.context-data.admin")
			if val != nil {
				cData.SetAdmin(val.Bool())
				return cData, true
			}
		}
	}

	return &contextData{}, false
}
