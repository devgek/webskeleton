package handler

import (
	"encoding/json"
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web"
	"net/http"
)

//HandleHealth ...
func HandleHealth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vd := web.NewViewData()
		vd["Host"] = r.Host
		vd["ProjectName"] = config.ProjectName
		vd["VersionInfo"] = config.ProjectVersion
		vd["health"] = "ok"

		json.NewEncoder(w).Encode(vd)
	})
}

//HandleLogout ...
func HandleLogout() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   web.AuthCookieName,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
}

//HandleFavicon ...
func HandleFavicon() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, web.AssetRoot+"/favicon_kahrersoftware.png")
	})
}

//HandlePageDefault ...
func HandlePageDefault(name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		web.RenderTemplate(w, r, name, nil)
	})
}
