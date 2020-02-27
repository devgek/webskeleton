package webecho

import (
	"io"
	"text/template"

	"github.com/labstack/echo"
	"kahrersoftware.at/webskeleton/web"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
}

// Render renders a template document
func (r *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var t *template.Template
	if name == "login.html" {
		t = template.Must(template.ParseFiles(web.TemplateRoot + name))
	} else {
		t = template.Must(template.ParseFiles(web.TemplateRoot+"layout.html", web.TemplateRoot+"menu.html", web.TemplateRoot+name))
	}
	return t.Execute(w, data)
}

//EchoContextData get ContextData from context
func EchoContextData(c echo.Context) web.ContextData {
	return web.FromContext(c.Request().Context())
}
