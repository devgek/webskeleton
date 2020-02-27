package webecho

import (
	"log"
	"net/http"

	"kahrersoftware.at/webskeleton/config"
	"kahrersoftware.at/webskeleton/web"

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
func (c *EchoController) InitWeb(e *echo.Echo) {
	e.Renderer = &TemplateRenderer{}

	e.GET("/health", c.HandleHealth())

	e.Match([]string{"GET", "POST"}, "/:name", c.DefaultPageHandler())

	e.POST("/loginuser", c.HandleLogin())

	e.GET("/logout", c.HandleLogout())

	e.Static(web.AssetPattern, web.AssetRoot)

	e.Use(echo.WrapMiddleware(web.LoggingMiddleware))
	e.Use(echo.WrapMiddleware(web.AuthMiddleware))
}

//DefaultPageHandler ...
func (c *EchoController) DefaultPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		pageName := c.Param("name")
		c.Render(http.StatusOK, pageName+".html", web.NewViewData(EchoContextData(c)))
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

		user, err := c.Env.Services.LoginUser(theUser, thePass)

		if err != nil {
			viewData := web.NewViewData(EchoContextData(ec))
			viewData["LoginUser"] = theUser
			viewData["LoginPass"] = thePass
			viewData["ErrorMessage"] = "Login with this credentials not allowed!"
			return ec.Render(http.StatusOK, "login.html", viewData)
		}

		//login ok
		log.Println("User", user.Name, "logged in")

		contextData := web.NewContextData()
		web.ToContext(r.Context(), contextData)
		contextData.SetUserID(theUser)

		cookieData := web.NewCookieData(contextData)

		authCookieValue := objx.New(cookieData).MustBase64()

		ec.SetCookie(&http.Cookie{
			Name:  web.AuthCookieName,
			Value: authCookieValue,
			Path:  "/"})
		return ec.Redirect(http.StatusTemporaryRedirect, "/page1")
	}

}

//HandleLogout logout user
func (c *EchoController) HandleLogout() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.SetCookie(&http.Cookie{
			Name:   web.AuthCookieName,
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
		vd := web.NewViewData(EchoContextData(c))
		vd["status"] = "ok"
		c.JSON(http.StatusOK, vd)

		return nil
	}

}
