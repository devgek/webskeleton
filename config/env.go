package config

import (
	"log"
	"net/http"
	"text/template"

	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm for sqlite3
	"kahrersoftware.at/webskeleton/models"
)

//Env the environment
type Env struct {
	DS                 models.Datastore
	TemplateHandlerMap map[string]*TemplateHandler
}

//InitEnv return new initialized environment
func InitEnv() *Env {
	//here we decide the database system
	ds, err := models.NewDS("sqlite3", "webskeleton.db")
	if err != nil {
		log.Panic(err)
	}

	return &Env{DS: ds, TemplateHandlerMap: make(map[string]*TemplateHandler)}
}

//NewViewData return view data map
func (env *Env) NewViewData(r *http.Request) map[string]interface{} {
	vd := make(map[string]interface{})
	vd["Host"] = r.Host
	vd["VersionInfo"] = "V1.0"
	if contextData, ok := FromContext(r.Context()); ok {
		vd["UserID"] = contextData.UserID()
	}

	return vd
}

//NewTemplateHandler create templateHandler and parse template
func (env *Env) NewTemplateHandler(fileName string) *TemplateHandler {
	th := &TemplateHandler{filename: fileName}
	env.TemplateHandlerMap[fileName] = th

	if th.filename == "login.html" {
		th.templ = template.Must(template.ParseFiles("./templates/" + fileName))
	} else {
		th.templ = template.Must(template.ParseFiles("./templates/layout.html", "./templates/menu.html", "./templates/"+fileName))
	}

	return th
}

//HandleView ...
func (env *Env) HandleView(w http.ResponseWriter, r *http.Request, templateName string, viewData interface{}) {
	th := env.TemplateHandlerMap[templateName]

	th.templ.Execute(w, viewData)
}
