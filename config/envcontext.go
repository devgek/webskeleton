package config

import (
	"github.com/labstack/echo"
)

//EnvContext extends echo.Context to provide the application environment
type EnvContext struct {
	echo.Context
	Env *Env
}

//RequestData get RequestData from context
func (ec *EnvContext) RequestData() RequestData {
	return ec.Get(ContextKeyRequestData).(RequestData)
}
