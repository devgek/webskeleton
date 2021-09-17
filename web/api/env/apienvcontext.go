package env

import (
	"github.com/labstack/echo"
)

//ApiEnvContext extends echo.Context to provide the application environment
type ApiEnvContext struct {
	echo.Context
	ApiEnv *ApiEnv
}
