package handler

import (
	"github.com/devgek/webskeleton/config"
	"github.com/labstack/echo"
	"net/http"
)

//HandleEntityList ...
func HandleEntityList(c echo.Context) error {
	//show entity list
	entity := c.Param("entity")

	ec := c.(*config.EnvContext)
	entities, err := ec.Env.Services.GetEntities(entity)
	viewData := config.NewTemplateDataWithRequestData(ec.RequestData())
	viewData["Entities"] = entities
	viewData["EditEntityType"] = ec.Env.MessageLocator.GetString("entity." + entity)
	if err != nil {
		viewData["ErrorMessage"] = err.Error()
	}
	return c.Render(http.StatusOK, entity, viewData)
}
