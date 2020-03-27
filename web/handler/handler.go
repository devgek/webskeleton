package handler

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web"
	"github.com/labstack/echo"
	"net/http"
)

//HandleHealth ...
func HandleHealth(c echo.Context) error {
	vd := web.NewViewData()
	vd["Host"] = c.Request().Host
	vd["ProjectName"] = config.ProjectName
	vd["VersionInfo"] = config.ProjectVersion
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

	return c.Render(http.StatusOK, page+".html", web.NewViewData())
}
