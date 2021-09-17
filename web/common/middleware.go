package common

import (
	"github.com/devgek/webskeleton/web/app/cookie"
	webenv "github.com/devgek/webskeleton/web/app/env"
	"github.com/devgek/webskeleton/web/app/request"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

//RequestLoggingMiddleware ...
func RequestLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("r:", c.Request().RequestURI, c.Request().Method)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		return next(c)
	}
}

//TokenLoggingMiddleware ...
func TokenLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("t:", c.Request().RequestURI, "token:", c.Get("token"))
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		return next(c)
	}
}

//CookieAuthMiddleware middleware handler for cookie authentication
func CookieAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("ca:", c.Request().RequestURI, c.Request().Method)
		r := c.Request()
		//don't check auth cookie with this requests
		if r.URL.Path == "/favicon.ico" || r.URL.Path == "/login" || r.URL.Path == "/loginuser" || r.URL.Path == "/health" || strings.Contains(r.URL.Path, "api") || strings.Contains(r.URL.Path, webenv.AssetPattern) || strings.Contains(r.URL.Path, "/assets") {
			log.Println("ca: skipped because of url exception")
			return next(c)
		}

		cookie, err := r.Cookie(webcookie.AuthCookieName)

		if err == http.ErrNoCookie {
			// not authenticated
			log.Println("a: ", r.URL.Path, " not authenticated!")
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		if err != nil {
			// some other error
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		//get request data from cookie and save to context
		rData, ok := request.FromCookie(cookie)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "illegal auth data")
		}

		c.Set(request.ContextKeyRequestData, rData)

		return next(c)
	}
}

// JWTAuthSkipper returns true for URL's, that do not need token authentication
func JWTAuthSkipper(c echo.Context) bool {
	r := c.Request()
	//don't check token for this URL's
	// method OPTIONS because of CORS Preflight requests from axios, they do not have Authorization token
	if r.URL.Path == "/api/login" || r.URL.Path == "/api/health" || r.Method == "OPTIONS" {
		log.Println("JWTAuthSkipper: skipped because of url or method exception")
		return true
	}
	return false
}
