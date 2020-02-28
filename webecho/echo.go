package webecho

import (
	"io"
	"log"

	"github.com/labstack/echo"
	"kahrersoftware.at/webskeleton/web"
)

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
