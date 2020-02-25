package web

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"kahrersoftware.at/webskeleton/services"
)

//Controller ...
type Controller struct {
	Services           *services.Services
	TemplateHandlerMap map[string]*TemplateHandler
}

//NewController ...
func NewController(services *services.Services) *Controller {
	return &Controller{Services: services, TemplateHandlerMap: make(map[string]*TemplateHandler)}
}

//InitRoutes ...
func (c *Controller) InitRoutes(r mux.Router) {
	r.Handle("/", c.NewTemplateHandler("login.html"))
	r.Handle("/loginuser", HandleLoginUser(c))
	//MustAuth secures the following site to be authenticated (auth cookie web-skeleton)
	r.Handle("/page1", MustAuth(c.NewTemplateHandler("page1.html")))
	r.Handle("/page2", c.NewTemplateHandler("page2.html"))
	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "webskeleton-auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	// r.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

}

//NewTemplateHandler create templateHandler and parse template
func (c *Controller) NewTemplateHandler(fileName string) *TemplateHandler {
	th := &TemplateHandler{filename: fileName}
	c.TemplateHandlerMap[fileName] = th

	if th.filename == "login.html" {
		th.templ = template.Must(template.ParseFiles("./templates/" + fileName))
	} else {
		th.templ = template.Must(template.ParseFiles("./templates/layout.html", "./templates/menu.html", "./templates/"+fileName))
	}

	return th
}

//HandleView ...
func (c *Controller) HandleView(w http.ResponseWriter, r *http.Request, templateName string, viewData interface{}) {
	th := c.TemplateHandlerMap[templateName]

	th.templ.Execute(w, viewData)
}

//NewViewData return view data map
func (c *Controller) NewViewData(r *http.Request) map[string]interface{} {
	vd := make(map[string]interface{})
	vd["Host"] = r.Host
	vd["VersionInfo"] = "V1.0"
	if contextData, ok := FromContext(r.Context()); ok {
		vd["UserID"] = contextData.UserID()
	}

	return vd
}
