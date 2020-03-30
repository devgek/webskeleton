package config

import (
	"github.com/devgek/webskeleton/web/request"
	"github.com/labstack/echo"
)

//EnvContext extends echo.Context to provide the application environment
type EnvContext struct {
	echo.Context
	Env *Env
}

//RequestData get RequestData from context
func (ec *EnvContext) RequestData() request.RData {
	return ec.Get(request.ContextKeyRequestData).(request.RData)
}
