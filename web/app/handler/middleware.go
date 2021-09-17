package handler

import (
	"github.com/devgek/webskeleton/web/app/env"
	"github.com/labstack/echo"
	"log"
)

//AppEnvContextMiddleware this is a custom echo context, representing the environment context
//it must be the first middleware, which is registered
func AppEnvContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		appEnv := env.GetAppEnv()
		log.Println("ec: template resolution dir:", appEnv.Templates.ResolutionDir)

		ec := &env.AppEnvContext{Context: c, Env: appEnv}
		return next(ec)
	}
}
