package main

import (
	"context"
	"net/http"

	"github.com/stretchr/objx"
)

//CookieData ...
type CookieData interface {
	Data() interface{}
}

type cookieData struct {
	CData ContextData
}

func (c cookieData) Data() interface{} {
	return c.CData
}

func (c cookieData) MSI() map[string]interface{} {
	ctxMap := objx.New(c.Data())

	return objx.New(map[string]interface{}{
		"cookie-data": ctxMap,
	})
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("webskeleton-auth")
	if err == http.ErrNoCookie {
		// not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	if err != nil {
		// some other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//set cookie data to context
	cData, ok := FromCookie(cookie)
	if !ok {
		http.Error(w, "illegal auth data", http.StatusInternalServerError)
		return
	}
	ctx := context.WithValue(r.Context(), contextKeyContextData, cData)
	// success - call the next handler
	h.next.ServeHTTP(w, r.WithContext(ctx))
}

// MustAuth adapts handler to ensure authentication has occurred.
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//todo do the login
	contextData := &contextData{"123456"}
	cookieData := &cookieData{contextData}

	authCookieValue := objx.New(cookieData).MustBase64()

	http.SetCookie(w, &http.Cookie{
		Name:  "webskeleton-auth",
		Value: authCookieValue,
		Path:  "/"})

	w.Header().Set("Location", "/page1")
	w.WriteHeader(http.StatusTemporaryRedirect)

}
