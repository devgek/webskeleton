package common

import (
	"github.com/labstack/echo"
	"log"
)

//RequestLoggingMiddleware ...
func RequestLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("r:", c.Request().RequestURI, c.Request().Method)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		return next(c)
	}
}

//TokenLoggingMiddleware ...
func TokenLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("t:", c.Request().RequestURI, "token:", c.Get("token"))
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		return next(c)
	}
}

// JWTAuthSkipper returns true for URL's, that do not need token authentication
func JWTAuthSkipper(c echo.Context) bool {
	r := c.Request()
	//don't check token for this URL's
	// method OPTIONS because of CORS Preflight requests from axios, they do not have Authorization token
	if r.URL.Path == "/api/login" || r.URL.Path == "/api/health" || r.Method == "OPTIONS" {
		log.Println("JWTAuthSkipper: skipped because of url or method exception")
		return true
	}
	return false
}
