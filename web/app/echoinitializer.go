package app

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web/app/env"
	apphandler "github.com/devgek/webskeleton/web/app/handler"
	"github.com/devgek/webskeleton/web/app/template"
	"github.com/devgek/webskeleton/web/common"
	"github.com/labstack/echo"
	"net/http"
)

//InitEchoWebApp initialize the echo web framework for serving a web application
func InitEchoWebApp(appEnv *env.AppEnv) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	if config.IsServerDebug() {
		e.Debug = true
		// e.Use(middleware.Recover())
	}
	e.Renderer = template.NewRenderer(appEnv.TStore)
	e.HTTPErrorHandler = common.LoggingDefaultHTTPErrorHandler

	// resoure files
	assetHandler := http.FileServer(appEnv.Assets)
	e.GET(env.AssetHandlerPattern, apphandler.AssetHandlerFunc(http.StripPrefix(env.AssetPattern, assetHandler)))
	//
	e.GET("/health", apphandler.HandleHealth)

	e.POST("/loginuser", apphandler.HandleLogin)

	e.GET("/logout", apphandler.HandleLogout)

	e.GET("/favicon.ico", common.HandleFavicon)

	e.Match([]string{"GET", "POST"}, "/entitylist:entity", apphandler.HandleEntityList)

	e.Match([]string{"OPTIONS", "POST"}, "/entitylist:entity", apphandler.HandleEntityListAjax)
	e.Match([]string{"OPTIONS", "POST"}, "/optionlist:entity", apphandler.HandleOptionListAjax)

	e.POST("/entityedit:entity", apphandler.HandleEntityEdit)
	e.POST("/entitynew:entity", apphandler.HandleEntityNew)
	e.POST("/entitydelete:entity", apphandler.HandleEntityDelete)

	e.Match([]string{"GET", "POST"}, "/", apphandler.HandleStartApp)
	e.Match([]string{"GET", "POST"}, "/page1", apphandler.HandlePage1)
	e.Match([]string{"GET", "POST"}, "/:page", apphandler.HandlePageDefault)

	e.Use(apphandler.AppEnvContextMiddleware)
	e.Use(common.RequestLoggingMiddleware)
	e.Use(common.CookieAuthMiddleware)

	return e
}
