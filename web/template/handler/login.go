package handler

import (
	"log"
	"net/http"

	"github.com/devgek/webskeleton/config"
	webcookie "github.com/devgek/webskeleton/web/cookie"
	webenv "github.com/devgek/webskeleton/web/env"
	"github.com/devgek/webskeleton/web/request"
	"github.com/labstack/echo"
	"github.com/stretchr/objx"
)

//HandleLogin ...
func HandleLogin(c echo.Context) error {
	//do the login
	theUser := c.FormValue("userid")
	thePass := c.FormValue("password")

	ec := c.(*webenv.EnvContext)
	user, err := ec.Env.Services.LoginUser(theUser, thePass)
	if err != nil {
		viewData := webenv.NewTemplateData()
		viewData["LoginUser"] = theUser
		viewData["LoginPass"] = thePass
		viewData["ErrorMessage"] = ec.Env.MessageLocator.GetMessageF(err.Error())
		return c.Render(http.StatusOK, "login", viewData)
	}

	//login ok
	log.Println("User", user.Name, "logged in")

	//hold userID and admin flag in request data
	requestData := request.NewRequestData()
	requestData.SetUserID(theUser)
	requestData.SetRole(user.Role)

	cookieData := webcookie.NewCookieData(requestData)

	authCookieValue := objx.New(cookieData).MustBase64()

	c.SetCookie(&http.Cookie{
		Name:  webcookie.AuthCookieName,
		Value: authCookieValue,
		Path:  "/"})

	return c.Redirect(http.StatusTemporaryRedirect, config.StartPage)
}
