package common

import (
	"github.com/devgek/webskeleton/web/app/env"
	"github.com/labstack/echo"
)

//HandleFavicon ...
func HandleFavicon(c echo.Context) error {
	return c.File(env.AppAssetRoot + "/favicon_kahrersoftware.png")
}
