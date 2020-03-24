package web

import (
	"encoding/json"
	"github.com/devgek/webskeleton/web/viewmodel"
	"log"
	"net/http"
	"strconv"

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
		viewData = NewViewDataWithContextData(FromContext(r.Context()))
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
		vd := NewViewData()
		vd["Host"] = r.Host
		vd["ProjectName"] = config.ProjectName
		vd["VersionInfo"] = config.ProjectVersion
		vd["health"] = "ok"

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
			viewData := NewViewDataWithContextData(FromContext(r.Context()))
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
		viewData := NewViewDataWithContextData(FromContext(r.Context()))
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
		oName := r.FormValue("gkvName")
		oEmail := r.FormValue("gkvEmail")
		oAdmin := r.FormValue("gkvAdmin")
		log.Println(oAdmin)

		contextData := NewContextData()
		ToContext(r.Context(), contextData)

		u, err := env.Services.UpdateUser(oName, oEmail, oAdmin == "true")

		vd := NewViewData()
		userEditResponse := viewmodel.NewUserEditResponse()
		if err != nil {
			userEditResponse.IsError = true
			userEditResponse.Message = err.Error()
			userEditResponse.Name = oName
			userEditResponse.Email = oEmail
			userEditResponse.Admin = (oAdmin == "true")
		} else {
			userEditResponse.Name = u.Name
			userEditResponse.Email = u.Email
			userEditResponse.Admin = u.Admin
		}

		vd["Response"] = userEditResponse
		json.NewEncoder(w).Encode(vd)
		return
	})

}

//HandleUserDelete ...
func HandleUserDelete(env *config.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		oID := r.FormValue("gkvObjId")
		ioID, _ := strconv.Atoi(oID)

		contextData := NewContextData()
		ToContext(r.Context(), contextData)

		err := env.Services.DeleteUser(uint(ioID))

		vd := NewViewData()
		baseResponse := &viewmodel.BaseResponse{}
		if err != nil {
			baseResponse.IsError = true
			baseResponse.Message = err.Error()
		}

		vd["Response"] = baseResponse
		json.NewEncoder(w).Encode(vd)
		return
	})

}

//HandlePageDefault ...
func HandlePageDefault(name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, r, name, nil)
	})
}
