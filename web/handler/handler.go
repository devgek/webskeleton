package handler

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/global"
	"github.com/devgek/webskeleton/web"
	"github.com/labstack/echo"
	"net/http"
)

//HandleHealth ...
func HandleHealth(c echo.Context) error {
	vd := config.NewTemplateData()
	vd["Host"] = c.Request().Host
	vd["ProjectName"] = global.ProjectName
	vd["VersionInfo"] = global.ProjectVersion
	vd["health"] = "ok"

	return c.JSON(http.StatusOK, vd)
}

//HandleLogout ...
func HandleLogout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:   web.AuthCookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

//HandleFavicon ...
func HandleFavicon(c echo.Context) error {
	return c.File(web.AssetRoot + "/favicon_kahrersoftware.png")
}

//HandlePageDefault ...
func HandlePageDefault(c echo.Context) error {
	page := c.Param("page")

	ec := c.(*config.EnvContext)
	return c.Render(http.StatusOK, page, config.NewTemplateDataWithRequestData(ec.RequestData()))
}
