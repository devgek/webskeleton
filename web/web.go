package web

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/stretchr/objx"
	"kahrersoftware.at/webskeleton/services"
)

//AssetPattern the pattern for the static file rout
var AssetPattern = "/assets"

//AssetRoot the root dir of the static asset files
var AssetRoot = "web/assets"

//Controller ...
type Controller struct {
	Services *services.Services
}

//NewController ...
func NewController(services *services.Services) *Controller {
	return &Controller{Services: services}
}

//InitWeb ...
func (c *Controller) InitWeb(r *mux.Router) {
	loginPageHandler := NewTemplateHandler("login.html")
	r.Handle("/health", c.HandleHealth())

	r.Handle("/", loginPageHandler)
	r.Handle("/login", loginPageHandler)
	r.Handle("/loginuser", c.HandleLoginUser())

	r.Handle("/page1", NewTemplateHandler("page1.html"))
	r.Handle("/page2", NewTemplateHandler("page2.html"))

	r.Handle("/logout", c.HandleLogout())

	r.PathPrefix(AssetPattern).Handler(http.StripPrefix(AssetPattern, http.FileServer(http.Dir(AssetRoot))))

	r.Use(c.loggingMiddleware)
	r.Use(authMiddleware)
}

//RenderView ...
func (c *Controller) RenderView(w http.ResponseWriter, r *http.Request, templateName string, viewData interface{}) {
	th := TemplateHandlerMap[templateName]

	th.templ.Execute(w, viewData)
}

//HandleHealth ...
func (c *Controller) HandleHealth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vd := NewViewData(FromContext(r.Context()))
		vd["status"] = "ok"
		json.NewEncoder(w).Encode(vd)
	})
}

//HandleLoginUser wrap handler func for login user
func (c *Controller) HandleLoginUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do the login
		theUser := r.FormValue("userid")
		thePass := r.FormValue("password")

		contextData := NewContextData()
		ctx := ToContext(r.Context(), contextData)

		user, err := c.Services.LoginUser(theUser, thePass)
		if err != nil {
			viewData := NewViewData(FromContext(r.Context()))
			viewData["LoginUser"] = theUser
			viewData["LoginPass"] = thePass
			viewData["ErrorMessage"] = "Login with this credentials not allowed!"
			c.RenderView(w, r.WithContext(ctx), "login.html", viewData)
			return
		}

		//login ok
		log.Println("User", user.Name, "logged in")
		contextData.SetUserID(theUser)
		ToContext(r.Context(), contextData)

		cookieData := &cookieData{contextData}

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
func (c *Controller) HandleLogout() http.Handler {
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

func (c *Controller) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println("r:", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
