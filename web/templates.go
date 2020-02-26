package web

import (
	"net/http"
	"text/template"
)

// TemplateHandler ...
type TemplateHandler struct {
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cData, _ := FromContext(r.Context())

	data := map[string]interface{}{
		"Host":        r.Host,
		"VersionInfo": "V1.0",
		"UserID":      cData.UserID(),
	}

	t.templ.Execute(w, data)

}

//NewTemplateHandler create templateHandler and parse template
func NewTemplateHandler(fileName string) *TemplateHandler {
	th := &TemplateHandler{filename: fileName}

	if th.filename == "login.html" {
		th.templ = template.Must(template.ParseFiles("./templates/" + fileName))
	} else {
		th.templ = template.Must(template.ParseFiles("./templates/layout.html", "./templates/menu.html", "./templates/"+fileName))
	}

	return th
}
