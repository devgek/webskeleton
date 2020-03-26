package webmux

import (
	"github.com/devgek/webskeleton/web/handler"
	"net/http"

	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web"

	"github.com/gorilla/mux"
)

//InitWeb ...
func InitWeb(env *config.Env) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/health", handler.HandleHealth())

	r.Handle("/loginuser", handler.HandleLogin(env))

	r.Handle("/users", handler.HandleUsers(env))

	r.Handle("/logout", handler.HandleLogout())

	r.PathPrefix(web.AssetPattern).Handler(http.StripPrefix(web.AssetPattern, http.FileServer(http.Dir(web.AssetRoot))))

	r.Handle("/{page}", DefaultPageHandler())

	r.Use(handler.RequestLoggingMiddleware)
	r.Use(handler.AuthMiddleware)

	return r
}

//DefaultPageHandler ...
func DefaultPageHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		page := vars["page"]
		if page == "" {
			page = "login"
		}
		web.RenderTemplate(w, r, page+".html", nil)
	})
}
