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
		t = template.Must(template.ParseFiles(TemplateRoot + name))
	} else {
		t = template.Must(template.ParseFiles(TemplateRoot+"layout.html", TemplateRoot+"menu.html", TemplateRoot+name))
	}
	return t.Execute(w, data)
}

//GetEchoContextData ...
func GetEchoContextData(c echo.Context) ContextData {
	contextData := c.Get("contextData")
	if contextData == nil {
		return NewContextData()
	}

	return contextData.(ContextData)
}
