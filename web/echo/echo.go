package echo

import (
	"net/http"

	"github.com/labstack/echo"
	"kahrersoftware.at/webskeleton/config"
	"kahrersoftware.at/webskeleton/global"
	"kahrersoftware.at/webskeleton/web"
	"kahrersoftware.at/webskeleton/web/handler"
	"kahrersoftware.at/webskeleton/web/template"
)

//InitEcho initialize the echo web framework
func InitEcho(env *config.Env) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	if global.IsServerDebug() {
		e.Debug = true
		// e.Use(middleware.Recover())
	}

	e.Renderer = template.NewRenderer(env.TStore)

	//api tests
	e.PUT("/apinew:entity", handler.HandleAPICreate)
	e.PUT("/apiallnew:entity", handler.HandleAPICreateAll)

	e.GET("/health", handler.HandleHealth)

	e.POST("/loginuser", handler.HandleLogin)

	e.GET("/logout", handler.HandleLogout)

	e.GET("/favicon.ico", handler.HandleFavicon)

	assetHandler := http.FileServer(env.Assets)
	// e.GET(web.AssetHandlerPattern, echo.WrapHandler(http.StripPrefix(web.AssetPattern, assetHandler)))
	e.GET(web.AssetHandlerPattern, handler.AssetHandlerFunc(http.StripPrefix(web.AssetPattern, assetHandler)))
	// e.Static(web.AssetPattern, web.AssetRoot)

	e.POST("/apientitylist:entity", handler.HandleEntityListAjax)
	e.POST("/apioptionlist:entity", handler.HandleOptionListAjax)

	e.Match([]string{"GET", "POST"}, "/entitylist:entity", handler.HandleEntityList)
	e.POST("/entityedit:entity", handler.HandleEntityEdit)
	e.POST("/entitynew:entity", handler.HandleEntityNew)
	e.POST("/entitydelete:entity", handler.HandleEntityDelete)

	e.Match([]string{"GET", "POST"}, "/", handler.HandleStartApp)
	e.Match([]string{"GET", "POST"}, "/page1", handler.HandlePage1)
	e.Match([]string{"GET", "POST"}, "/:page", handler.HandlePageDefault)

	e.Use(handler.EnvContextMiddleware)
	e.Use(handler.RequestLoggingMiddleware)
	e.Use(handler.AuthMiddleware)

	return e
}
