package handler

import (
	"github.com/devgek/webskeleton/config"
	"github.com/labstack/echo"
	"net/http"
)

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
