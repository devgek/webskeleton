package handler

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web"
	"github.com/labstack/echo"
	"github.com/stretchr/objx"
	"log"
	"net/http"
)

//HandleLogin ...
func HandleLogin(c echo.Context) error {
	//do the login
	theUser := c.FormValue("userid")
	thePass := c.FormValue("password")

	ec := c.(*config.EnvContext)
	user, err := ec.Env.Services.LoginUser(theUser, thePass)
	if err != nil {
		viewData := web.NewViewData()
		viewData["LoginUser"] = theUser
		viewData["LoginPass"] = thePass
		viewData["ErrorMessage"] = ec.Env.MessageLocator.GetString("msg.error.login")
		return c.Render(http.StatusOK, "login.html", viewData)
	}

	//login ok
	log.Println("User", user.Name, "logged in")

	//hold userID and admin flag in request data
	requestData := config.NewRequestData()
	requestData.SetUserID(theUser)
	requestData.SetAdmin(user.Admin)

	cookieData := web.NewCookieData(requestData)

	authCookieValue := objx.New(cookieData).MustBase64()

	c.SetCookie(&http.Cookie{
		Name:  web.AuthCookieName,
		Value: authCookieValue,
		Path:  "/"})

	return c.Redirect(http.StatusTemporaryRedirect, "/page1")
}
