package web

import (
	"net/http"

	"kahrersoftware.at/bmdexport/logs"

	"github.com/stretchr/objx"
)

//CookieData ...
type CookieData interface {
	Data() interface{}
}

type cookieData struct {
	CData ContextData
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
	cData, ok := FromCookie(cookie)
	if !ok {
		http.Error(w, "illegal auth data", http.StatusInternalServerError)
		return
	}
	ctx := ToContext(r.Context(), cData)
	// success - call the next handler
	h.next.ServeHTTP(w, r.WithContext(ctx))
}

// MustAuth adapts handler to ensure authentication has occurred.
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

//HandleLoginUser wrap handler func for login user
func HandleLoginUser(c *Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do the login
		theUser := r.FormValue("userid")
		thePass := r.FormValue("password")

		contextData := NewContextData()
		ctx := ToContext(r.Context(), contextData)

		user, err := c.Services.LoginUser(theUser, thePass)
		if err != nil {
			viewData := c.NewViewData(r)
			viewData["LoginUser"] = theUser
			viewData["LoginPass"] = thePass
			viewData["ErrorMessage"] = "Login with this credentials not allowed!"
			c.HandleView(w, r.WithContext(ctx), "login.html", viewData)
			return
		}

		//login ok
		logs.Debug("User ", user.Name, " logged in")
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
