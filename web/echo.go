package web

import (
	"io"
	"text/template"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
}

// Render renders a template document
func (r *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var t *template.Template
	if name == "login.html" {
		t = template.Must(template.ParseFiles("./templates/" + name))
	} else {
		t = template.Must(template.ParseFiles("./templates/layout.html", "./templates/menu.html", "./templates/"+name))
	}
	return t.ExecuteTemplate(w, name, data)
}
