package common

import (
	"github.com/labstack/echo"
)

//HandleFavicon ...
func HandleFavicon(c echo.Context) error {
	return c.File("web/common/favicon_kahrersoftware.png")
}
