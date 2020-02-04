package main

import (
	"context"
	"net/http"

	"github.com/stretchr/objx"
)

//ContextData the data hold in request context
type ContextData interface {
	UserID() string
}
type contextData struct {
	userID string
}

func (c contextData) UserID() string {
	return c.userID
}

func (c contextData) MSI() map[string]interface{} {
	ctxMap := objx.New(map[string]interface{}{
		"userID": c.UserID(),
	})
	return objx.New(map[string]interface{}{
		"context-data": ctxMap,
	})
}

type contextKey struct {
	name string
}

var contextKeyContextData = &contextKey{"context-data"}

//FromContext get ContextData out of context
func FromContext(ctx context.Context) (ContextData, bool) {
	key := ctx.Value(contextKeyContextData)
	if key == nil {
		return &contextData{""}, false
	}
	cData, ok := key.(ContextData)
	return cData, ok
}

//FromCookie get ContextData from cookie value
func FromCookie(cookie *http.Cookie) (ContextData, bool) {
	c := objx.MustFromBase64(cookie.Value)

	val := c.Get("cookie-data.context-data.userID")
	if val != nil {
		return &contextData{val.Str()}, true
	}
	return &contextData{""}, false
}
