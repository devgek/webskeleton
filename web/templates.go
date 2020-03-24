package web

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/msg"
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
	data := NewViewDataWithContextData(FromContext(r.Context()))

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
		th.Templ = template.Must(template.ParseFiles(TemplateRoot+"layout.html", TemplateRoot+fileName, TemplateRoot+"user-edit.html", TemplateRoot+"confirm-delete.html"))
	}

	return th
}

//NewViewDataWithContextData return view data map filled with context data
func NewViewDataWithContextData(contextData ContextData) map[string]interface{} {
	vd := NewViewData()

	vd["Host"] = contextData.Host()
	vd["Messages"] = msg.Messages
	vd["ProjectName"] = config.ProjectName
	vd["VersionInfo"] = config.ProjectVersion
	vd["UserID"] = contextData.UserID()
	vd["Admin"] = contextData.Admin()

	return vd
}

//NewViewData ...
func NewViewData() map[string]interface{} {
	return make(map[string]interface{})
}
