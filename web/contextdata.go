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
}
type contextData struct {
	userID string
	host   string
}

//NewContextData create ContextData
func NewContextData() ContextData {
	return &contextData{"", ""}
}

func (c contextData) UserID() string {
	return c.userID
}

func (c contextData) Host() string {
	return c.userID
}

func (c *contextData) SetUserID(userID string) {
	c.userID = userID
}

func (c *contextData) SetHost(host string) {
	c.host = host
}

func (c contextData) MSI() map[string]interface{} {
	ctxMap := objx.New(map[string]interface{}{
		"userID": c.UserID(),
		"host":   c.Host(),
	})
	return objx.New(map[string]interface{}{
		"context-data": ctxMap,
	})
}

type contextKey struct {
	name string
}

var contextKeyContextData = &contextKey{"context-data"}

//ToContext set ContextData to context
func ToContext(ctx context.Context, cData ContextData) context.Context {
	return context.WithValue(ctx, contextKeyContextData, cData)
}

//FromContext get ContextData out of context
func FromContext(ctx context.Context) ContextData {
	val := ctx.Value(contextKeyContextData)
	if val == nil {
		return NewContextData()
	}
	return val.(ContextData)
}

//FromCookie get ContextData from cookie value
func FromCookie(cookie *http.Cookie) (ContextData, bool) {
	c := objx.MustFromBase64(cookie.Value)

	val := c.Get("cookie-data.context-data.userID")
	if val != nil {
		return &contextData{userID: val.Str()}, true
	}
	return &contextData{}, false
}
