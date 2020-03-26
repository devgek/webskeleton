package handler

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web"
	"github.com/stretchr/objx"
	"log"
	"net/http"
)

//HandleLogin ...
func HandleLogin(env *config.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do the login
		theUser := r.FormValue("userid")
		thePass := r.FormValue("password")

		contextData := web.NewContextData()
		ctx := web.ToContext(r.Context(), contextData)

		user, err := env.Services.LoginUser(theUser, thePass)
		if err != nil {
			viewData := web.NewViewDataWithContextData(web.FromContext(r.Context()))
			viewData["LoginUser"] = theUser
			viewData["LoginPass"] = thePass
			viewData["ErrorMessage"] = env.MessageLocator.GetString("msg.error.login")
			web.RenderTemplate(w, r.WithContext(ctx), "login.html", viewData)
			return
		}

		//login ok
		log.Println("User", user.Name, "logged in")
		contextData.SetUserID(theUser)
		contextData.SetAdmin(user.Admin)
		web.ToContext(r.Context(), contextData)

		cookieData := web.NewCookieData(contextData)

		authCookieValue := objx.New(cookieData).MustBase64()

		http.SetCookie(w, &http.Cookie{
			Name:  web.AuthCookieName,
			Value: authCookieValue,
			Path:  "/"})

		w.Header().Set("Location", "/page1")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

}
