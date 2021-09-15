package echo

import (
	apihandler "github.com/devgek/webskeleton/web/api/handler"
	templatehandler "github.com/devgek/webskeleton/web/template/handler"
	"net/http"

	"github.com/devgek/webskeleton/config"
	webenv "github.com/devgek/webskeleton/web/env"
	"github.com/devgek/webskeleton/web/handler"
	"github.com/devgek/webskeleton/web/template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//InitEchoApi initialize the echo web framework for serving an api
func InitEchoApi(env *webenv.Env) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	if config.IsServerDebug() {
		e.Debug = true
		// e.Use(middleware.Recover())
	}

	//e.Renderer = template.NewRenderer(env.TStore)
	e.HTTPErrorHandler = handler.LoggingDefaultHTTPErrorHandler

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
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/login", apihandler.HandleAPILogin)

	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entitylist:entity", apihandler.HandleAPIEntityList)
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/optionlist:entity", apihandler.HandleAPIOptionList)

	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entitynew:entity", apihandler.HandleAPICreateEntity)
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entityedit:entity", apihandler.HandleAPIUpdateEntity)
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entitydelete:entity/:id", apihandler.HandleAPIDeleteEntity)

	apiGroup.PUT("/allnew:entity", apihandler.HandleAPICreateAll)
	apiGroup.GET("/health", apihandler.HandleAPIHealth)

	// resoure files
	//assetHandler := http.FileServer(env.Assets)
	//e.GET(webenv.AssetHandlerPattern, handler.AssetHandlerFunc(http.StripPrefix(webenv.AssetPattern, assetHandler)))

	e.GET("/favicon.ico", templatehandler.HandleFavicon)

	e.Use(handler.EnvContextMiddleware)
	e.Use(handler.RequestLoggingMiddleware)

	return e
}

//InitEchoWebApp initialize the echo web framework for serving a web application
func InitEchoWebApp(env *webenv.Env) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	if config.IsServerDebug() {
		e.Debug = true
		// e.Use(middleware.Recover())
	}

	e.Renderer = template.NewRenderer(env.TStore)
	e.HTTPErrorHandler = handler.LoggingDefaultHTTPErrorHandler

	// resoure files
	assetHandler := http.FileServer(env.Assets)
	e.GET(webenv.AssetHandlerPattern, handler.AssetHandlerFunc(http.StripPrefix(webenv.AssetPattern, assetHandler)))
	//
	e.GET("/health", templatehandler.HandleHealth)

	e.POST("/loginuser", templatehandler.HandleLogin)

	e.GET("/logout", templatehandler.HandleLogout)

	e.GET("/favicon.ico", templatehandler.HandleFavicon)

	e.Match([]string{"GET", "POST"}, "/entitylist:entity", handler.HandleEntityList)

	e.Match([]string{"OPTIONS", "POST"}, "/entitylist:entity", handler.HandleEntityListAjax)
	e.Match([]string{"OPTIONS", "POST"}, "/optionlist:entity", handler.HandleOptionListAjax)

	e.POST("/entityedit:entity", handler.HandleEntityEdit)
	e.POST("/entitynew:entity", handler.HandleEntityNew)
	e.POST("/entitydelete:entity", handler.HandleEntityDelete)

	e.Match([]string{"GET", "POST"}, "/", templatehandler.HandleStartApp)
	e.Match([]string{"GET", "POST"}, "/page1", templatehandler.HandlePage1)
	e.Match([]string{"GET", "POST"}, "/:page", templatehandler.HandlePageDefault)

	e.Use(handler.EnvContextMiddleware)
	e.Use(handler.RequestLoggingMiddleware)
	e.Use(handler.CookieAuthMiddleware)

	return e
}
