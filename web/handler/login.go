package handler

import (
	"log"
	"net/http"

	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/types"
	webcookie "github.com/devgek/webskeleton/web/cookie"
	webenv "github.com/devgek/webskeleton/web/env"
	"github.com/devgek/webskeleton/web/request"
	"github.com/dgrijalva/jwt-go"
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

type LoginData struct {
	User string
	Pass string
}

//HandleApiLogin handles login to api and returns a JWT token
func HandleAPILogin(c echo.Context) error {
	log.Println("HandleApiLogin")
	//do the login
	loginData := LoginData{}
	if err := c.Bind(&loginData); err != nil {
		return err
	}

	ec := c.(*webenv.EnvContext)
	user, err := ec.Env.Services.LoginUser(loginData.User, loginData.Pass)
	if err != nil {
		// return echo.NewHTTPError(http.StatusUnauthorized)
		msg := ec.Env.MessageLocator.GetMessageF(err.Error())
		log.Println("HandleApiLogin return:", http.StatusUnauthorized, msg)
		return c.JSON(http.StatusUnauthorized, msg)
	}

	//login ok
	log.Println("User", user.Name, "logged in for api")

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims into webtoken, content can be checked on further requests with token
	claims := token.Claims.(jwt.MapClaims)
	isAdmin := (user.Role == types.RoleTypeAdmin)
	claims["name"] = loginData.User
	claims["admin"] = isAdmin

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": t,
		"name":  loginData.User,
		"admin": isAdmin,
	})
}
