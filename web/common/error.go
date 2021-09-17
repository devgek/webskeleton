package common

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func SimpleLoggingHTTPErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		log.Println("eh: HTTPError", he.Code, he.Message, he.Internal)
	} else {
		log.Println("eh: Error:", err.Error())
	}
}

func LoggingDefaultHTTPErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		if he.Internal != nil {
			err = fmt.Errorf("%v, %v", err, he.Internal)
		}
	} else if c.Echo().Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}
	if _, ok := msg.(string); ok {
		msg = echo.Map{"message": msg}
	}

	c.Echo().Logger.Error(err)

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
