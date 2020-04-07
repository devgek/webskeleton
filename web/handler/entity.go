package handler

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web/viewmodel"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

//HandleEntityList ...
func HandleEntityList(c echo.Context) error {
	//show entity list
	entity := c.Param("entity")

	ec := c.(*config.EnvContext)
	entities := ec.Env.EF.GetSlice(entity)
	err := ec.Env.DS.GetAllEntities(entities)

	viewData := config.NewTemplateDataWithRequestData(ec.RequestData())
	viewData["Entities"] = entities
	viewData["EditEntityType"] = ec.Env.MessageLocator.GetString("entity." + entity)
	if err != nil {
		viewData["ErrorMessage"] = err.Error()
	}
	return c.Render(http.StatusOK, entity, viewData)
}

//HandleEntityDelete ...
func HandleEntityDelete(c echo.Context) error {
	ec := c.(*config.EnvContext)

	entity := c.Param("entity")
	entityName := ec.Env.MessageLocator.GetString("entity." + entity)

	oID := c.FormValue("gkvObjId")
	ioID, _ := strconv.Atoi(oID)

	entityModel := ec.Env.EF.Get(entity)
	err := ec.Env.DS.DeleteEntityByID(entityModel, uint(ioID))

	baseResponse := &viewmodel.BaseResponse{}
	if err != nil {
		baseResponse.IsError = true
		baseResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.error.entity.delete", entityName)
	} else {
		baseResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.success.entity.delete", entityName)
	}

	return c.JSON(http.StatusOK, baseResponse)
}

//HandleEntityEdit ...
func HandleEntityEdit(c echo.Context) error {
	oID := c.FormValue("gkvObjId")
	ioID, _ := strconv.Atoi(oID)

	ec := c.(*config.EnvContext)
	entity := ec.Param("entity")
	oEntityObject := ec.Env.EF.Get(entity)

	entityResponse := viewmodel.NewEntityResponse(oEntityObject)
	entityName := ec.Env.MessageLocator.GetString("entity." + entity)

	err := ec.Env.DS.GetEntityByID(oEntityObject, uint(ioID))
	if err == nil {
		if err := ec.Bind(oEntityObject); err != nil {
			return err
		}

		err = ec.Env.DS.SaveEntity(oEntityObject)
	}

	if err != nil {
		entityResponse.IsError = true
		entityResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.error.entity.edit", entityName)
	} else {
		entityResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.success.entity.edit", entityName)
	}

	//on client entityResponse is received as javascript object, no JSON.parse is needed
	return c.JSON(http.StatusOK, entityResponse)
}

//HandleEntityNew ...
func HandleEntityNew(c echo.Context) error {
	ec := c.(*config.EnvContext)
	entity := ec.Param("entity")
	oEntityObject := ec.Env.EF.Get(entity)

	entityResponse := viewmodel.NewEntityResponse(oEntityObject)
	entityName := ec.Env.MessageLocator.GetString("entity." + entity)

	if err := ec.Bind(oEntityObject); err != nil {
		return err
	}

	err := ec.Env.Services.CreateEntity(oEntityObject, entity)

	if err != nil {
		entityResponse.IsError = true
		entityResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.error.entity.create", entityName)
	} else {
		entityResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.success.entity.create", entityName)
	}

	//on client entityResponse is received as javascript object, no JSON.parse is needed
	return c.JSON(http.StatusOK, entityResponse)
}
