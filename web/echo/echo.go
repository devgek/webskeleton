package echo

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web"
	"github.com/devgek/webskeleton/web/handler"
	"github.com/devgek/webskeleton/web/template"
	"github.com/labstack/echo"
	"net/http"
)

//InitEcho initialize the echo web framework
func InitEcho(env *config.Env) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Renderer = template.NewRenderer(env.TStore)

	e.GET("/health", handler.HandleHealth)

	e.POST("/loginuser", handler.HandleLogin)

	e.GET("/logout", handler.HandleLogout)

	e.GET("/favicon.ico", handler.HandleFavicon)

	assetHandler := http.FileServer(env.Assets)
	e.GET(web.AssetHandlerPattern, echo.WrapHandler(http.StripPrefix(web.AssetPattern, assetHandler)))
	// e.Static(web.AssetPattern, web.AssetRoot)

	e.Match([]string{"GET", "POST"}, "/users", handler.HandleUsers)
	e.POST("/useredit", handler.HandleUserEdit)
	e.POST("/usernew", handler.HandleUserNew)
	e.POST("/userdelete", handler.HandleUserDelete)

	e.Match([]string{"GET", "POST"}, "/:page", handler.HandlePageDefault)

	e.Use(handler.EnvContextMiddleware)
	e.Use(handler.RequestLoggingMiddleware)
	e.Use(handler.AuthMiddleware)

	return e
}
