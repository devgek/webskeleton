package handler

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)

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
		if r.URL.Path == "/login" || r.URL.Path == "/loginuser" || r.URL.Path == "/health" || strings.Contains(r.URL.Path, web.AssetPattern) {
			return next(c)
		}

		cookie, err := r.Cookie(web.AuthCookieName)

		if err == http.ErrNoCookie {
			// not authenticated
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		if err != nil {
			// some other error
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		//get request data from cookie and save to context
		rData, ok := config.FromCookie(cookie)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "illegal auth data")
		}

		c.Set(config.ContextKeyRequestData, rData)

		return next(c)
	}
}
