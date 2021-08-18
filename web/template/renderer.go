package template

import (
	"io"
	"log"

	"github.com/labstack/echo"
)

// Renderer is a custom html/template renderer for Echo framework
// damit man echo.Context.Render aufrufen kann
type Renderer struct {
	TStore TStore
}

//NewRenderer ...
func NewRenderer(store TStore) *Renderer {
	return &Renderer{store}
}

// Render renders a template document
func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	log.Println("Renderer::Render: template", name)
	if templ, err := r.TStore.GetTemplate(name); err == nil {
		//important templ.Execute not templ.ExecuteTemplate(w, name, data)
		return templ.Execute(w, data)
	} else {
		return err
	}
}
