package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stretchr/objx"
	"kahrersoftware.at/webskeleton/config"
)

//AssetPattern the pattern for the static file rout
var AssetPattern = "/assets"

//AssetRoot the root dir of the static asset files
var AssetRoot = "web/assets"

//RenderView ...
func RenderView(w http.ResponseWriter, r *http.Request, templateName string, viewData interface{}) {
	th := TemplateHandlerMap[templateName]

	th.Templ.Execute(w, viewData)
}

//HandleHealth ...
func HandleHealth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vd := NewViewData(FromContext(r.Context()))
		vd["status"] = "ok"
		json.NewEncoder(w).Encode(vd)
	})
}

//HandleLogin ...
func HandleLogin(env *config.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do the login
		theUser := r.FormValue("userid")
		thePass := r.FormValue("password")

		contextData := NewContextData()
		ctx := ToContext(r.Context(), contextData)

		user, err := env.Services.LoginUser(theUser, thePass)
		if err != nil {
			viewData := NewViewData(FromContext(r.Context()))
			viewData["LoginUser"] = theUser
			viewData["LoginPass"] = thePass
			viewData["ErrorMessage"] = "Login with this credentials not allowed!"
			RenderView(w, r.WithContext(ctx), "login.html", viewData)
			return
		}

		//login ok
		log.Println("User", user.Name, "logged in")
		contextData.SetUserID(theUser)
		ToContext(r.Context(), contextData)

		cookieData := NewCookieData(contextData)

		authCookieValue := objx.New(cookieData).MustBase64()

		http.SetCookie(w, &http.Cookie{
			Name:  AuthCookieName,
			Value: authCookieValue,
			Path:  "/"})

		w.Header().Set("Location", "/page1")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

}

//HandleLogout ...
func HandleLogout() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   AuthCookieName,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
}
