package web

import (
	"log"
	"net/http"

	"kahrersoftware.at/webskeleton/config"

	"github.com/labstack/echo"
	"github.com/stretchr/objx"
)

//EchoController front controller using echo framework
type EchoController struct {
	Env *config.Env
}

//NewEchoController ...
func NewEchoController(env *config.Env) *EchoController {
	return &EchoController{env}
}

//InitWeb initialize the web framework
func (c *EchoController) InitWeb(echo *echo.Echo) {
	echo.Renderer = &TemplateRenderer{}

	echo.GET("/health", c.HandleHealth())

	echo.Match([]string{"GET", "POST"}, "/:name", c.DefaultPageHandler())

	echo.POST("/loginuser", c.HandleLogin())

	echo.GET("/logout", c.HandleLogout())

	echo.Static(AssetPattern, AssetRoot)

	echo.Use(c.AuthMiddleware)
}

//AuthMiddleware ...
func (c *EchoController) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ec echo.Context) error {
		r := ec.Request()
		//don't check auth cookie with this requests
		if r.URL.Path == "/login" || r.URL.Path == "/loginuser" || r.URL.Path == "/health" {
			return next(ec)
		}

		cookie, err := r.Cookie(AuthCookieName)
		if err == http.ErrNoCookie {
			// not authenticated
			return ec.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		if err != nil {
			// some other error
			return ec.String(http.StatusInternalServerError, err.Error())
		}
		//set cookie data to context
		cData, ok := FromCookie(cookie)
		if !ok {
			return ec.String(http.StatusInternalServerError, "illegal auth data")
		}
		ec.Set("contextData", cData)
		// ToContext(r.Context(), cData)
		return next(ec)
	}
}

//DefaultPageHandler ...
func (c *EchoController) DefaultPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		pageName := c.Param("name")
		c.Render(http.StatusOK, pageName+".html", NewViewData(GetEchoContextData(c)))
		return nil
	}
}

//HandleLogin login user
func (c *EchoController) HandleLogin() echo.HandlerFunc {
	return func(ec echo.Context) error {
		//do the login
		r := ec.Request()
		theUser := r.FormValue("userid")
		thePass := r.FormValue("password")

		contextData := NewContextData()
		ToContext(r.Context(), contextData)

		user, err := c.Env.Services.LoginUser(theUser, thePass)

		if err != nil {
			viewData := NewViewData(GetEchoContextData(ec))
			viewData["LoginUser"] = theUser
			viewData["LoginPass"] = thePass
			viewData["ErrorMessage"] = "Login with this credentials not allowed!"
			return ec.Render(http.StatusOK, "login.html", viewData)
		}

		//login ok
		log.Println("User", user.Name, "logged in")
		contextData.SetUserID(theUser)
		ec.Set("contextData", contextData)
		// ToContext(r.Context(), contextData)

		cookieData := &cookieData{contextData}

		authCookieValue := objx.New(cookieData).MustBase64()

		ec.SetCookie(&http.Cookie{
			Name:  AuthCookieName,
			Value: authCookieValue,
			Path:  "/"})
		return ec.Redirect(http.StatusTemporaryRedirect, "/page1")
	}

}

//HandleLogout logout user
func (c *EchoController) HandleLogout() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.SetCookie(&http.Cookie{
			Name:   AuthCookieName,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}

//HandleHealth ...
func (c *EchoController) HandleHealth() echo.HandlerFunc {
	return func(c echo.Context) error {
		vd := NewViewData(GetEchoContextData(c))
		vd["status"] = "ok"
		c.JSON(http.StatusOK, vd)

		return nil
	}

}
