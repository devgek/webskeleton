package template

import (
	"io"
	"log"

	"github.com/labstack/echo"
)

// TStoreRenderer is a custom html/template renderer for Echo framework, it uses a template.TStore for rendering templates
// damit man echo.Context.Render aufrufen kann
type TStoreRenderer struct {
	TStore TStore
}

//NewRenderer ...
func NewRenderer(store TStore) *TStoreRenderer {
	return &TStoreRenderer{store}
}

// Render renders a template document
func (r *TStoreRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	log.Println("TStoreRenderer::Render: template", name)
	if templ, err := r.TStore.GetTemplate(name); err == nil {
		//important templ.Execute not templ.ExecuteTemplate(w, name, data)
		return templ.Execute(w, data)
	} else {
		return err
	}
}
