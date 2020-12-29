package handler

import (
	"log"
	"net/http"

	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/global"
	"github.com/devgek/webskeleton/web"
	"github.com/devgek/webskeleton/web/request"
	"github.com/labstack/echo"
	"github.com/stretchr/objx"
)

//HandleLogin ...
func HandleLogin(c echo.Context) error {
	//do the login
	theUser := c.FormValue("userid")
	thePass := c.FormValue("password")

	ec := c.(*config.EnvContext)
	user, err := ec.Env.Services.LoginUser(theUser, thePass)
	if err != nil {
		viewData := config.NewTemplateData()
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

	cookieData := web.NewCookieData(requestData)

	authCookieValue := objx.New(cookieData).MustBase64()

	c.SetCookie(&http.Cookie{
		Name:  web.AuthCookieName,
		Value: authCookieValue,
		Path:  "/"})

	return c.Redirect(http.StatusTemporaryRedirect, global.StartPage)
}
