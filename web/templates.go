package web

import (
	"log"
	"net/http"
	"sync"
	"text/template"
)

//TemplateHandlerMap ...
var TemplateHandlerMap = make(map[string]*TemplateHandler)

//TemplateRoot rootdir for template files
var TemplateRoot = "./web/templates/"

// TemplateHandler ...
type TemplateHandler struct {
	sync.Mutex
	theMap   *map[string]*TemplateHandler
	filename string
	Templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cData := FromContext(r.Context())

	data := NewTemplateData(cData)

	t.Templ.Execute(w, data)

}

//NewTemplateHandler create templateHandler and parse template
func NewTemplateHandler(fileName string) *TemplateHandler {
	th := &TemplateHandler{theMap: &TemplateHandlerMap, filename: fileName}
	th.Lock()
	defer th.Unlock()
	TemplateHandlerMap[fileName] = th
	log.Println("sync new template handler in map for", fileName)

	if th.filename == "login.html" {
		th.Templ = template.Must(template.ParseFiles(TemplateRoot + fileName))
	} else {
		th.Templ = template.Must(template.ParseFiles(TemplateRoot+"layout.html", TemplateRoot+"menu.html", TemplateRoot+fileName))
	}

	return th
}

//NewTemplateData return view data map
func NewTemplateData(contextData ContextData) map[string]interface{} {
	vd := make(map[string]interface{})
	vd["Host"] = contextData.Host()
	vd["VersionInfo"] = "V1.0"
	vd["UserID"] = contextData.UserID()

	return vd
}
