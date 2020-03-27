package handler

import (
	"github.com/devgek/webskeleton/config"
	"github.com/labstack/echo"
)

//EnvContextMiddleware this is a custom echo context, representing the environment context
//it must be the first middleware, which is registered
func EnvContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &config.EnvContext{Context: c, Env: config.GetWebEnv()}
		return next(cc)
	}
}
