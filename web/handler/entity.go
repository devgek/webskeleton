package handler

import (
	"github.com/devgek/webskeleton/models"
	webenv "github.com/devgek/webskeleton/web/env"
	"github.com/devgek/webskeleton/web/viewmodel"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

//HandleEntityList ...
func HandleEntityList(c echo.Context) error {
	//show entity list
	entity := c.Param("entity")

	ec := c.(*webenv.EnvContext)
	entities, err := ec.Env.EF.GetSlice(entity)
	if err != nil {
		return err
	}

	err = ec.Env.DS.GetAllEntities(entities)

	viewData := webenv.NewTemplateDataWithRequestData(ec.RequestData())
	viewData["Entities"] = entities
	viewData["EditEntityType"] = ec.Env.MessageLocator.GetString("entity." + entity)
	if entity == "consumptiongroup" {
		viewData["EmbeddedEntityType"] = ec.Env.MessageLocator.GetString("entity." + "energymetermapping")
	}
	if err != nil {
		viewData["ErrorMessage"] = err.Error()
	}
	return c.Render(http.StatusOK, entity, viewData)
}

//HandleEntityDelete ...
func HandleEntityDelete(c echo.Context) error {
	ec := c.(*webenv.EnvContext)

	entity := c.Param("entity")
	entityName := ec.Env.MessageLocator.GetString("entity." + entity)

	oID := c.FormValue("gkvObjId")
	ioID, _ := strconv.Atoi(oID)

	entityModel, err := ec.Env.EF.Get(entity)
	if err != nil {
		return err
	}

	err = ec.Env.DS.DeleteEntityByID(entityModel, uint(ioID))

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

	ec := c.(*webenv.EnvContext)
	entity := ec.Param("entity")
	oEntityObject, err := ec.Env.EF.Get(entity)
	if err != nil {
		return err
	}

	entityResponse := viewmodel.NewEntityResponse(oEntityObject)
	entityName := ec.Env.MessageLocator.GetString("entity." + entity)

	err = ec.Env.DS.GetEntityByID(oEntityObject, uint(ioID))
	if err == nil {
		//Attention!! embedded structs with same field names are also populated with form values (e.g. consumptiongroup.name and customer.name)
		//TODO: find a solution for this problem
		if err := ec.Bind(oEntityObject); err != nil {
			return err
		}
		//embedded structs are not saved because of gorm:"association_autoupdate:false"
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
	ec := c.(*webenv.EnvContext)
	entity := ec.Param("entity")

	oEntityObject, err := ec.Env.EF.Get(entity)
	if err != nil {
		return err
	}

	entityResponse := viewmodel.NewEntityResponse(oEntityObject)
	entityName := ec.Env.MessageLocator.GetString("entity." + entity)

	// params, _ := ec.FormParams()
	// log.Println("params:", params)

	//Attention!! embedded structs with same field names are also populated with form values (e.g. consumptiongroup.name and customer.name)
	//TODO: find a solution for this problem
	if err := ec.Bind(oEntityObject); err != nil {
		log.Println("Error while binding entity ", entity, " ", err.Error())
		return err
	}

	if entity == "user" {
		user := oEntityObject.(*models.User)
		entityResponse.EntityObject, err = ec.Env.Services.CreateUser(user.Name, user.Pass, user.Email, user.Role)
		if err != nil {
			entityResponse.Message = ec.Env.MessageLocator.GetMessageF(err.Error())
		} else {
			entityResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.success.entity.create", entityName)
		}
	} else {
		err = ec.Env.DS.CreateEntity(oEntityObject)
		if err != nil {
			entityResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.error.entity.create", entityName)
		} else {
			entityResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.success.entity.create", entityName)
		}
	}

	if err != nil {
		entityResponse.IsError = true
	}

	//on client entityResponse is received as javascript object, no JSON.parse is needed
	return c.JSON(http.StatusOK, entityResponse)
}
