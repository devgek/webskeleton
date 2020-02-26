package web

import (
	"net/http"

	"github.com/stretchr/objx"
)

//AuthCookieName the name of the auth cookie
var AuthCookieName = "webskeleton-auth"

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

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//don't check auth cookie with this requests
		if r.URL.Path == "/login" || r.URL.Path == "/loginuser" || r.URL.Path == "/health" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie(AuthCookieName)
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
		ctx := ToContext(r.Context(), cData)
		// success - call the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
