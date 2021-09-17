package handler

import (
	"github.com/devgek/webskeleton/web/app/cookie"
	"github.com/devgek/webskeleton/web/app/env"
	"github.com/devgek/webskeleton/web/app/template"
	"net/http"

	"github.com/devgek/webskeleton/config"
	"github.com/labstack/echo"
)

//HandleHealth ...
func HandleHealth(c echo.Context) error {
	vd := template.NewTemplateData()
	vd["Host"] = c.Request().Host
	vd["ProjectName"] = config.ProjectName
	vd["VersionInfo"] = config.ProjectVersion
	vd["health"] = "ok"

	return c.JSON(http.StatusOK, vd)
}

//HandleStartApp ...
func HandleStartApp(c echo.Context) error {
	startPage := config.StartPage

	return c.Redirect(http.StatusTemporaryRedirect, startPage)
}

//HandleLogout ...
func HandleLogout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:   webcookie.AuthCookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

//HandlePageDefault ...
func HandlePageDefault(c echo.Context) error {
	page := c.Param("page")

	ec := c.(*env.AppEnvContext)
	return c.Render(http.StatusOK, page, template.NewTemplateDataWithRequestData(ec.RequestData()))
}
