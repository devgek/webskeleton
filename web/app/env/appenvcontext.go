package env

import (
	"github.com/devgek/webskeleton/web/app/request"
	"github.com/labstack/echo"
	"log"
)

//AppEnvContext extends echo.Context to provide the application environment
type AppEnvContext struct {
	echo.Context
	Env *AppEnv
}

//RequestData get RequestData from context
func (ec *AppEnvContext) RequestData() request.RData {
	i := ec.Get(request.ContextKeyRequestData)
	if i == nil {
		return request.NewRequestData()
	}

	return i.(request.RData)
}

//FormValueRequired ...
func (ec *AppEnvContext) FormValueRequired(formValue string) string {
	v := ec.Context.FormValue(formValue)
	if v == "" {
		log.Panic("Missing required form value:", formValue)
	}

	return v
}
