package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"kahrersoftware.at/webskeleton/config"
	"kahrersoftware.at/webskeleton/web"
	"kahrersoftware.at/webskeleton/web/request"
)

//EnvContextMiddleware this is a custom echo context, representing the environment context
//it must be the first middleware, which is registered
func EnvContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &config.EnvContext{Context: c, Env: config.GetWebEnv()}
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
		if r.URL.Path == "/login" || r.URL.Path == "/loginuser" || r.URL.Path == "/health" || strings.Contains(r.URL.Path, "api") || strings.Contains(r.URL.Path, web.AssetPattern) {
			return next(c)
		}

		cookie, err := r.Cookie(web.AuthCookieName)

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
