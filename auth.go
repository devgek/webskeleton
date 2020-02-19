package main

import (
	"bytes"
	"net/http"

	"kahrersoftware.at/webskeleton/config"

	"github.com/stretchr/objx"
)

//CookieData ...
type CookieData interface {
	Data() interface{}
}

type cookieData struct {
	CData config.ContextData
}

func (c cookieData) Data() interface{} {
	return c.CData
}

func (c cookieData) MSI() map[string]interface{} {
	ctxMap := objx.New(c.Data())

	return objx.New(map[string]interface{}{
		"cookie-data": ctxMap,
	})
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("webskeleton-auth")
	if err == http.ErrNoCookie {
		// not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	if err != nil {
		// some other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//set cookie data to context
	cData, ok := config.FromCookie(cookie)
	if !ok {
		http.Error(w, "illegal auth data", http.StatusInternalServerError)
		return
	}
	ctx := config.ToContext(r.Context(), cData)
	// success - call the next handler
	h.next.ServeHTTP(w, r.WithContext(ctx))
}

// MustAuth adapts handler to ensure authentication has occurred.
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

func handleLoginUser(env *config.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do the login
		theUser := r.FormValue("userid")
		thePass := r.FormValue("password")

		contextData := config.NewContextData()
		ctx := config.ToContext(r.Context(), contextData)

		user, err := env.DS.GetUser(theUser)
		if err != nil || !bytes.Equal(user.Pass, []byte(thePass)) {
			viewData := env.NewViewData(r)
			viewData["LoginUser"] = theUser
			viewData["LoginPass"] = thePass
			viewData["ErrorMessage"] = "Login with this credentials not allowed!"
			env.HandleView(w, r.WithContext(ctx), "login.html", viewData)
			return
		}

		//login ok
		contextData.SetUserID(theUser)
		cookieData := &cookieData{contextData}

		authCookieValue := objx.New(cookieData).MustBase64()

		http.SetCookie(w, &http.Cookie{
			Name:  "webskeleton-auth",
			Value: authCookieValue,
			Path:  "/"})

		w.Header().Set("Location", "/page1")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

}
