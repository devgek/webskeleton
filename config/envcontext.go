package config

import (
	"log"

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
	i := ec.Get(request.ContextKeyRequestData)
	if i == nil {
		return request.NewRequestData()
	}

	return i.(request.RData)
}

//FormValueRequired ...
func (ec *EnvContext) FormValueRequired(formValue string) string {
	v := ec.Context.FormValue(formValue)
	if v == "" {
		log.Panic("Missing required form value:", formValue)
	}

	return v
}
