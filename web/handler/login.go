package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/stretchr/objx"
	"kahrersoftware.at/webskeleton/config"
	"kahrersoftware.at/webskeleton/global"
	"kahrersoftware.at/webskeleton/web"
	"kahrersoftware.at/webskeleton/web/request"
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
		viewData["ErrorMessage"] = ec.Env.MessageLocator.GetString(err.Error())
		return c.Render(http.StatusOK, "login", viewData)
	}

	//login ok
	log.Println("User", user.Name, "logged in")

	//hold userID and admin flag in request data
	requestData := request.NewRequestData()
	requestData.SetUserID(theUser)
	requestData.SetRole(user.Role)
	requestData.SetCustomerID(user.CustomerID)

	cookieData := web.NewCookieData(requestData)

	authCookieValue := objx.New(cookieData).MustBase64()

	c.SetCookie(&http.Cookie{
		Name:  web.AuthCookieName,
		Value: authCookieValue,
		Path:  "/"})

	return c.Redirect(http.StatusTemporaryRedirect, global.StartPage)
}
