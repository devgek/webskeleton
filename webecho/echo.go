package webecho

import (
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"kahrersoftware.at/webskeleton/config"
	"kahrersoftware.at/webskeleton/web"
)

//EchoController front controller using echo framework
type EchoController struct {
	Env *config.Env
}

//NewEchoController ...
func NewEchoController(env *config.Env) *EchoController {
	return &EchoController{env}
}

//InitWeb initialize the web framework
func (c *EchoController) InitWeb(e *echo.Echo) {
	e.Renderer = &TemplateRenderer{}

	e.GET("/health", echo.WrapHandler(web.HandleHealth()))

	e.Match([]string{"GET", "POST"}, "/:name", c.DefaultPageHandler())

	e.POST("/loginuser", echo.WrapHandler(web.HandleLogin(c.Env)))

	e.GET("/logout", echo.WrapHandler(web.HandleLogout()))

	e.Static(web.AssetPattern, web.AssetRoot)

	e.Use(echo.WrapMiddleware(web.LoggingMiddleware))
	e.Use(echo.WrapMiddleware(web.AuthMiddleware))
}

//DefaultPageHandler ...
func (c *EchoController) DefaultPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		pageName := c.Param("name")
		c.Render(http.StatusOK, pageName+".html", web.NewViewData(EchoContextData(c)))
		return nil
	}
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
}

// Render renders a template document
func (r *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	log.Println("render", name)
	th := web.NewTemplateHandler(name)
	return th.Templ.Execute(w, data)
}

//EchoContextData get ContextData from context
func EchoContextData(c echo.Context) web.ContextData {
	return web.FromContext(c.Request().Context())
}
