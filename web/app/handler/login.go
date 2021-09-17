package handler

import (
	"github.com/devgek/webskeleton/web/app/cookie"
	"github.com/devgek/webskeleton/web/app/env"
	"github.com/devgek/webskeleton/web/app/request"
	"github.com/devgek/webskeleton/web/app/template"
	"log"
	"net/http"

	"github.com/devgek/webskeleton/config"
	"github.com/labstack/echo"
	"github.com/stretchr/objx"
)

//HandleLogin ...
func HandleLogin(c echo.Context) error {
	//do the login
	theUser := c.FormValue("userid")
	thePass := c.FormValue("password")

	ec := c.(*env.AppEnvContext)
	user, err := ec.Env.Services.LoginUser(theUser, thePass)
	if err != nil {
		viewData := template.NewTemplateData()
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
