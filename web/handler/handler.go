package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"kahrersoftware.at/webskeleton/config"
	"kahrersoftware.at/webskeleton/global"
	"kahrersoftware.at/webskeleton/web"
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

//HandleStartApp ...
func HandleStartApp(c echo.Context) error {
	startPage := global.StartPage

	return c.Redirect(http.StatusTemporaryRedirect, startPage)
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

//AssetHandlerFunc handles asset files
func AssetHandlerFunc(h http.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		//cache assets in browser for one day
		if global.IsAssetsCache() {
			c.Response().Header().Set("Cache-Control", "public, max-age=86400")
		}
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
