package webmux

import (
	"net/http"

	"kahrersoftware.at/webskeleton/config"
	"kahrersoftware.at/webskeleton/web"

	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	Env *config.Env
}

//NewController ...
func NewController(env *config.Env) *Controller {
	return &Controller{env}
}

//InitWeb ...
func (c *Controller) InitWeb(r *mux.Router) {
	loginPageHandler := web.NewTemplateHandler("login.html")
	r.Handle("/health", web.HandleHealth())

	r.Handle("/", loginPageHandler)
	r.Handle("/login", loginPageHandler)
	r.Handle("/loginuser", web.HandleLogin(c.Env))

	r.Handle("/page1", web.NewTemplateHandler("page1.html"))
	r.Handle("/page2", web.NewTemplateHandler("page2.html"))

	r.Handle("/logout", web.HandleLogout())

	r.PathPrefix(web.AssetPattern).Handler(http.StripPrefix(web.AssetPattern, http.FileServer(http.Dir(web.AssetRoot))))

	r.Use(web.LoggingMiddleware)
	r.Use(web.AuthMiddleware)
}
