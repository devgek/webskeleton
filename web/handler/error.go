package handler

import (
	"log"

	"github.com/labstack/echo"
)

func SimpleLoggingHTTPErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		log.Println("eh: HTTPError", he.Code, he.Message, he.Internal)
	} else {
		log.Println("eh: Error:", err.Error())
	}
}
