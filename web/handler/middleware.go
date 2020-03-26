package handler

import (
	"github.com/devgek/webskeleton/web"
	"log"
	"net/http"
	"strings"
)

//RequestLoggingMiddleware ...
func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("r:", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

//AuthMiddleware middleware handler for cookie authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//don't check auth cookie with this requests
		if r.URL.Path == "/login" || r.URL.Path == "/loginuser" || r.URL.Path == "/health" || strings.Contains(r.URL.Path, web.AssetPattern) {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie(web.AuthCookieName)

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
		cData, ok := web.FromCookie(cookie)
		if !ok {
			http.Error(w, "illegal auth data", http.StatusInternalServerError)
			return
		}
		ctx := web.ToContext(r.Context(), cData)
		// success - call the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
