package webecho

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web"
	"github.com/devgek/webskeleton/web/handler"
	"github.com/devgek/webskeleton/web/template"
	"github.com/labstack/echo"
	"io"
	"log"
	"net/http"
)

//InitWeb initialize the web framework
func InitWeb(env *config.Env) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Renderer = NewTemplateRenderer(env)

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

// TemplateRenderer is a custom html/template renderer for Echo framework
// damit man echo.Context.Render aufrufen kann
type TemplateRenderer struct {
	TStore template.TStore
}

//NewTemplateRenderer ...
func NewTemplateRenderer(env *config.Env) *TemplateRenderer {
	return &TemplateRenderer{env.TStore}
}

// Render renders a template document
func (r *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	log.Println("render", name)
	templ := r.TStore.GetTemplate(name)
	//important templ.Execute not templ.ExecuteTemplate(w, name, data)
	return templ.Execute(w, data)
}
