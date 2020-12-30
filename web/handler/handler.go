package handler

import (
	"net/http"

	"github.com/devgek/webskeleton/config"
	webcookie "github.com/devgek/webskeleton/web/cookie"
	webenv "github.com/devgek/webskeleton/web/env"

	"github.com/labstack/echo"
)

//HandleHealth ...
func HandleHealth(c echo.Context) error {
	vd := webenv.NewTemplateData()
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

//HandleFavicon ...
func HandleFavicon(c echo.Context) error {
	return c.File(webenv.AssetRoot + "/favicon_kahrersoftware.png")
}

//HandlePageDefault ...
func HandlePageDefault(c echo.Context) error {
	page := c.Param("page")

	ec := c.(*webenv.EnvContext)
	return c.Render(http.StatusOK, page, webenv.NewTemplateDataWithRequestData(ec.RequestData()))
}

//AssetHandlerFunc handles asset files
func AssetHandlerFunc(h http.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		//cache assets in browser for one day
		if config.IsAssetsCache() {
			c.Response().Header().Set("Cache-Control", "public, max-age=86400")
		}
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
