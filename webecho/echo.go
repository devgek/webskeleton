package webecho

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web"
	"github.com/devgek/webskeleton/web/handler"
	"github.com/labstack/echo"
	"io"
	"log"
)

//InitWeb initialize the web framework
func InitWeb(env *config.Env) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Renderer = &TemplateRenderer{}

	e.GET("/health", handler.HandleHealth)

	e.POST("/loginuser", handler.HandleLogin)

	e.GET("/logout", handler.HandleLogout)

	e.GET("/favicon.ico", handler.HandleFavicon)

	e.Static(web.AssetPattern, web.AssetRoot)

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

// TemplateRenderer is a custom html/template renderer for Echo framework
// damit man echo.Context.Render aufrufen kann
type TemplateRenderer struct {
}

// Render renders a template document
func (r *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	log.Println("render", name)
	th := web.NewTemplateHandler(name)
	return th.Templ.Execute(w, data)
}
