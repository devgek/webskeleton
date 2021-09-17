package apihandler

import (
	"github.com/devgek/webskeleton/web/api/env"
	"github.com/labstack/echo"
)

//ApiEnvContextMiddleware this is a custom echo context, representing the environment context
//it must be the first middleware, which is registered
func ApiEnvContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiEnv := env.GetApiEnv(false)

		ec := &env.ApiEnvContext{Context: c, ApiEnv: apiEnv}
		return next(ec)
	}
}
