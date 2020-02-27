package web

import (
	"net/http"
	"text/template"
)

//TemplateHandlerMap ...
var TemplateHandlerMap map[string]*TemplateHandler = make(map[string]*TemplateHandler)

//TemplateRoot rootdir for template files
var TemplateRoot = "./web/templates/"

// TemplateHandler ...
type TemplateHandler struct {
	theMap   *map[string]*TemplateHandler
	filename string
	Templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cData := FromContext(r.Context())

	data := map[string]interface{}{
		"Host":        r.Host,
		"VersionInfo": "V1.0",
		"UserID":      cData.UserID(),
	}

	t.Templ.Execute(w, data)

}

//NewTemplateHandler create templateHandler and parse template
func NewTemplateHandler(fileName string) *TemplateHandler {
	th := &TemplateHandler{theMap: &TemplateHandlerMap, filename: fileName}
	TemplateHandlerMap[fileName] = th

	if th.filename == "login.html" {
		th.Templ = template.Must(template.ParseFiles(TemplateRoot + fileName))
	} else {
		th.Templ = template.Must(template.ParseFiles(TemplateRoot+"layout.html", TemplateRoot+"menu.html", TemplateRoot+fileName))
	}

	return th
}
