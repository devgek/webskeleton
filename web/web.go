package web

import (
	"net/http"

	"github.com/devgek/webskeleton/config"
)

//AssetPattern the pattern for the static file rout
var AssetPattern = "/assets"

//AssetRoot the root dir of the static asset files
var AssetRoot = "web/assets"

//RenderTemplate ...
func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, viewData interface{}) {
	if viewData == nil {
		viewData = NewViewDataWithContextData(FromContext(r.Context()))
	}
	th := TemplateHandlerMap[templateName]
	//reload template in debug mode
	if config.Debug || th == nil {
		th = NewTemplateHandler(templateName)
	}

	th.Templ.Execute(w, viewData)
}
