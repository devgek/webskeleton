package handler

import (
	webcookie "github.com/devgek/webskeleton/web/app/cookie"
	"github.com/devgek/webskeleton/web/app/env"
	"github.com/devgek/webskeleton/web/app/request"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)

//AppEnvContextMiddleware this is a custom echo context, representing the environment context
//it must be the first middleware, which is registered
func AppEnvContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		appEnv := env.GetAppEnv()
		log.Println("ec: template resolution dir:", appEnv.Templates.ResolutionDir)

		ec := &env.AppEnvContext{Context: c, Env: appEnv}
		return next(ec)
	}
}

//CookieAuthMiddleware middleware handler for cookie authentication
func CookieAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("ca:", c.Request().RequestURI, c.Request().Method)
		r := c.Request()
		//don't check auth cookie with this requests
		if r.URL.Path == "/favicon.ico" || r.URL.Path == "/login" || r.URL.Path == "/loginuser" || r.URL.Path == "/health" || strings.Contains(r.URL.Path, "api") || strings.Contains(r.URL.Path, env.AssetPattern) || strings.Contains(r.URL.Path, "/assets") {
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
