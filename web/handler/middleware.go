package handler

import (
	"log"
	"net/http"
	"strings"

	webcookie "github.com/devgek/webskeleton/web/cookie"
	webenv "github.com/devgek/webskeleton/web/env"
	"github.com/devgek/webskeleton/web/request"
	"github.com/labstack/echo"
)

//EnvContextMiddleware this is a custom echo context, representing the environment context
//it must be the first middleware, which is registered
func EnvContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &webenv.EnvContext{Context: c, Env: webenv.GetWebEnv()}
		return next(cc)
	}
}

//RequestLoggingMiddleware ...
func RequestLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("r:", c.Request().RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		return next(c)
	}
}

//AuthMiddleware middleware handler for cookie authentication
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		//don't check auth cookie with this requests
		if r.URL.Path == "/favicon.ico" || r.URL.Path == "/login" || r.URL.Path == "/loginuser" || r.URL.Path == "/health" || strings.Contains(r.URL.Path, "api") || strings.Contains(r.URL.Path, webenv.AssetPattern) || strings.Contains(r.URL.Path, "/vue") {
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
