package echo

import (
	"net/http"

	"github.com/devgek/webskeleton/config"
	webenv "github.com/devgek/webskeleton/web/env"
	"github.com/devgek/webskeleton/web/handler"
	"github.com/devgek/webskeleton/web/template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//InitEcho initialize the echo web framework
func InitEcho(env *webenv.Env) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	if config.IsServerDebug() {
		e.Debug = true
		// e.Use(middleware.Recover())
	}

	e.Renderer = template.NewRenderer(env.TStore)

	// api
	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		ContextKey: "token",
		Skipper:    handler.JWTAuthSkipper,
	}))
	apiGroup.Use(handler.TokenLoggingMiddleware)

	apiGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodHead},
	}))

	// OPTIONS because of CORS Preflight requests sent from axios
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/login", handler.HandleAPILogin)

	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entitylist:entity", handler.HandleEntityListAjax)
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/optionlist:entity", handler.HandleOptionListAjax)

	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entitynew:entity", handler.HandleAPICreate)
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entityedit:entity", handler.HandleAPIEdit)
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entitydelete:entity/:id", handler.HandleAPIDelete)

	apiGroup.PUT("/allnew:entity", handler.HandleAPICreateAll)

	// resoure files
	assetHandler := http.FileServer(env.Assets)
	e.GET(webenv.AssetHandlerPattern, handler.AssetHandlerFunc(http.StripPrefix(webenv.AssetPattern, assetHandler)))
	//
	e.GET("/health", handler.HandleHealth)

	e.POST("/loginuser", handler.HandleLogin)

	e.GET("/logout", handler.HandleLogout)

	e.GET("/favicon.ico", handler.HandleFavicon)

	e.Match([]string{"GET", "POST"}, "/entitylist:entity", handler.HandleEntityList)
	e.POST("/entityedit:entity", handler.HandleEntityEdit)
	e.POST("/entitynew:entity", handler.HandleEntityNew)
	e.POST("/entitydelete:entity", handler.HandleEntityDelete)

	e.Match([]string{"GET", "POST"}, "/", handler.HandleStartApp)
	e.Match([]string{"GET", "POST"}, "/page1", handler.HandlePage1)
	e.Match([]string{"GET", "POST"}, "/:page", handler.HandlePageDefault)

	e.Use(handler.EnvContextMiddleware)
	e.Use(handler.RequestLoggingMiddleware)
	e.Use(handler.CookieAuthMiddleware)

	return e
}
