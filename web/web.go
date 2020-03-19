package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devgek/webskeleton/config"
	"github.com/stretchr/objx"
)

//AssetPattern the pattern for the static file rout
var AssetPattern = "/assets"

//AssetRoot the root dir of the static asset files
var AssetRoot = "web/assets"

//RenderTemplate ...
func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, viewData interface{}) {
	if viewData == nil {
		viewData = NewTemplateData(FromContext(r.Context()))
	}
	th := TemplateHandlerMap[templateName]
	//reload template in debug mode
	if config.Debug || th == nil {
		th = NewTemplateHandler(templateName)
	}

	th.Templ.Execute(w, viewData)
}

//HandleHealth ...
func HandleHealth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vd := NewTemplateData(FromContext(r.Context()))
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
			viewData := NewTemplateData(FromContext(r.Context()))
			viewData["LoginUser"] = theUser
			viewData["LoginPass"] = thePass
			viewData["ErrorMessage"] = err.Error()
			RenderTemplate(w, r.WithContext(ctx), "login.html", viewData)
			return
		}

		//login ok
		log.Println("User", user.Name, "logged in")
		contextData.SetUserID(theUser)
		contextData.SetAdmin(user.Admin)
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

//HandleUsers ...
func HandleUsers(env *config.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//show user list
		contextData := NewContextData()
		ctx := ToContext(r.Context(), contextData)

		users, err := env.Services.GetAllUsers()
		viewData := NewTemplateData(FromContext(r.Context()))
		viewData["Users"] = users
		if err != nil {
			viewData["ErrorMessage"] = err.Error()
		}
		RenderTemplate(w, r.WithContext(ctx), "users.html", viewData)
		return
	})

}

//HandleUserEdit ...
func HandleUserEdit(env *config.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// uName := r.FormValue("gkvName")
		// uPass := r.FormValue("gkvPass")
		// uEmail := r.FormValue("gkvEmail")
		// uAdmin := r.FormValue("gkvAdmin")

		contextData := NewContextData()
		ctx := ToContext(r.Context(), contextData)

		users, err := env.Services.GetAllUsers()
		viewData := NewTemplateData(FromContext(r.Context()))
		viewData["Users"] = users
		if err != nil {
			viewData["ErrorMessage"] = err.Error()
		}
		RenderTemplate(w, r.WithContext(ctx), "users.html", viewData)
		return
	})

}

//HandlePageDefault ...
func HandlePageDefault(name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, r, name, nil)
	})
}
