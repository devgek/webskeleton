package web

import (
	"encoding/json"
	"net/http"
	"text/template"

	"log"

	"github.com/gorilla/mux"
	"github.com/stretchr/objx"
	"kahrersoftware.at/webskeleton/services"
)

//Controller ...
type Controller struct {
	Services           *services.Services
	TemplateHandlerMap map[string]*TemplateHandler
}

//NewController ...
func NewController(services *services.Services) *Controller {
	return &Controller{Services: services, TemplateHandlerMap: make(map[string]*TemplateHandler)}
}

//InitRoutes ...
func (c *Controller) InitRoutes(r *mux.Router) {
	loginPageHandler := c.NewTemplateHandler("login.html")
	r.Handle("/", loginPageHandler)
	r.Handle("/login", loginPageHandler)
	r.Handle("/loginuser", c.HandleLoginUser())
	//MustAuth secures the following site to be authenticated (auth cookie web-skeleton)
	r.Handle("/page1", c.NewTemplateHandler("page1.html"))
	r.Handle("/page2", c.NewTemplateHandler("page2.html"))

	r.Handle("/logout", c.HandleLogout())
	r.Handle("/health", c.HandleHealth())

	// r.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	r.Use(loggingMiddleware)
	r.Use(authMiddleware)
}

//NewTemplateHandler create templateHandler and parse template
func (c *Controller) NewTemplateHandler(fileName string) *TemplateHandler {
	th := &TemplateHandler{filename: fileName}
	c.TemplateHandlerMap[fileName] = th

	if th.filename == "login.html" {
		th.templ = template.Must(template.ParseFiles("./templates/" + fileName))
	} else {
		th.templ = template.Must(template.ParseFiles("./templates/layout.html", "./templates/menu.html", "./templates/"+fileName))
	}

	return th
}

//HandleView ...
func (c *Controller) HandleView(w http.ResponseWriter, r *http.Request, templateName string, viewData interface{}) {
	th := c.TemplateHandlerMap[templateName]

	th.templ.Execute(w, viewData)
}

//NewViewData return view data map
func (c *Controller) NewViewData(r *http.Request) map[string]interface{} {
	vd := make(map[string]interface{})
	vd["Host"] = r.Host
	vd["VersionInfo"] = "V1.0"
	if contextData, ok := FromContext(r.Context()); ok {
		vd["UserID"] = contextData.UserID()
	}

	return vd
}

//HandleHealth ...
func (c *Controller) HandleHealth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vd := c.NewViewData(r)
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
			viewData := c.NewViewData(r)
			viewData["LoginUser"] = theUser
			viewData["LoginPass"] = thePass
			viewData["ErrorMessage"] = "Login with this credentials not allowed!"
			c.HandleView(w, r.WithContext(ctx), "login.html", viewData)
			return
		}

		//login ok
		log.Println("User", user.Name, "logged in")
		contextData.SetUserID(theUser)
		ToContext(r.Context(), contextData)

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

//HandleLogout ...
func (c *Controller) HandleLogout() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "webskeleton-auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println("r:", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func xMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println("x:", r.Referer())
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func yMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println("y:", r.RemoteAddr)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
