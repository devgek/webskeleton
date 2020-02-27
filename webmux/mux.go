package webmux

import (
	"encoding/json"
	"net/http"

	"kahrersoftware.at/webskeleton/web"

	"log"

	"github.com/gorilla/mux"
	"github.com/stretchr/objx"
	"kahrersoftware.at/webskeleton/services"
)

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
	loginPageHandler := web.NewTemplateHandler("login.html")
	r.Handle("/health", c.HandleHealth())

	r.Handle("/", loginPageHandler)
	r.Handle("/login", loginPageHandler)
	r.Handle("/loginuser", c.HandleLoginUser())

	r.Handle("/page1", web.NewTemplateHandler("page1.html"))
	r.Handle("/page2", web.NewTemplateHandler("page2.html"))

	r.Handle("/logout", c.HandleLogout())

	r.PathPrefix(web.AssetPattern).Handler(http.StripPrefix(web.AssetPattern, http.FileServer(http.Dir(web.AssetRoot))))

	r.Use(web.LoggingMiddleware)
	r.Use(web.AuthMiddleware)
}

//RenderView ...
func (c *Controller) RenderView(w http.ResponseWriter, r *http.Request, templateName string, viewData interface{}) {
	th := web.TemplateHandlerMap[templateName]

	th.Templ.Execute(w, viewData)
}

//HandleHealth ...
func (c *Controller) HandleHealth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vd := web.NewViewData(web.FromContext(r.Context()))
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

		contextData := web.NewContextData()
		ctx := web.ToContext(r.Context(), contextData)

		user, err := c.Services.LoginUser(theUser, thePass)
		if err != nil {
			viewData := web.NewViewData(web.FromContext(r.Context()))
			viewData["LoginUser"] = theUser
			viewData["LoginPass"] = thePass
			viewData["ErrorMessage"] = "Login with this credentials not allowed!"
			c.RenderView(w, r.WithContext(ctx), "login.html", viewData)
			return
		}

		//login ok
		log.Println("User", user.Name, "logged in")
		contextData.SetUserID(theUser)
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

//HandleLogout ...
func (c *Controller) HandleLogout() http.Handler {
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
