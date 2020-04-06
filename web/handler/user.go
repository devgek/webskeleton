package handler

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web/viewmodel"
	"github.com/labstack/echo"
	"net/http"
)

//HandleUserEdit ...
func HandleUserEdit(c echo.Context) error {
	oName := c.FormValue("gkvName")
	oEmail := c.FormValue("gkvEmail")
	oAdmin := c.FormValue("gkvAdmin")

	ec := c.(*config.EnvContext)
	u, err := ec.Env.Services.UpdateUser(oName, oEmail, oAdmin == "true")

	userEditResponse := viewmodel.NewUserEditResponse()
	if err != nil {
		userEditResponse.IsError = true
		userEditResponse.Message = ec.Env.MessageLocator.GetString("msg.error.user.edit")
		userEditResponse.Name = oName
		userEditResponse.Email = oEmail
		userEditResponse.Admin = (oAdmin == "true")
	} else {
		userEditResponse.Message = ec.Env.MessageLocator.GetString("msg.success.user.edit")
		userEditResponse.Name = u.Name
		userEditResponse.Email = u.Email
		userEditResponse.Admin = u.Admin
	}

	//on client userEditResponse is received as javascript object, no JSON.parse is needed
	return c.JSON(http.StatusOK, userEditResponse)
}

//HandleUserNew ...
func HandleUserNew(c echo.Context) error {
	oName := c.FormValue("gkvName")
	oPass := c.FormValue("gkvPass")
	oEmail := c.FormValue("gkvEmail")
	oAdmin := c.FormValue("gkvAdmin")

	ec := c.(*config.EnvContext)
	u, err := ec.Env.Services.CreateUser(oName, oPass, oEmail, oAdmin == "true")

	userEditResponse := viewmodel.NewUserEditResponse()
	if err != nil {
		userEditResponse.IsError = true
		userEditResponse.Message = ec.Env.MessageLocator.GetString("msg.error.user.create")
		userEditResponse.Name = oName
		userEditResponse.Pass = oPass
		userEditResponse.Email = oEmail
		userEditResponse.Admin = (oAdmin == "true")
	} else {
		userEditResponse.Message = ec.Env.MessageLocator.GetString("msg.success.user.create")
		userEditResponse.Name = u.Name
		userEditResponse.Pass = string(u.Pass)
		userEditResponse.Email = u.Email
		userEditResponse.Admin = u.Admin
	}

	return c.JSON(http.StatusOK, userEditResponse)
}
