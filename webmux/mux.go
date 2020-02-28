package webmux

import (
	"net/http"

	"kahrersoftware.at/webskeleton/config"
	"kahrersoftware.at/webskeleton/web"

	"github.com/gorilla/mux"
)

//InitWeb ...
func InitWeb(env *config.Env) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/health", web.HandleHealth())

	r.Handle("/loginuser", web.HandleLogin(env))

	r.Handle("/logout", web.HandleLogout())

	r.PathPrefix(web.AssetPattern).Handler(http.StripPrefix(web.AssetPattern, http.FileServer(http.Dir(web.AssetRoot))))

	r.Handle("/{page}", DefaultPageHandler())

	r.Use(web.LoggingMiddleware)
	r.Use(web.AuthMiddleware)

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
