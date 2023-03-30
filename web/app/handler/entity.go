package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/devgek/webskeleton/models"
	genmodels "github.com/devgek/webskeleton/models/generated"
	"github.com/devgek/webskeleton/web/app/env"
	"github.com/devgek/webskeleton/web/app/template"
	viewmodel2 "github.com/devgek/webskeleton/web/app/viewmodel"
	"github.com/labstack/echo"
)

// HandleEntityList ...
func HandleEntityList(c echo.Context) error {
	//show entity list
	entityName := c.Param("entity")

	ec := c.(*env.AppEnvContext)
	entities, err := ec.Env.EF.GetEntitySlice(entityName)
	if err != nil {
		return err
	}
	entity, err := ec.Env.EF.GetEntity(entityName)
	if err != nil {
		return err
	}

	err = ec.Env.DS.GetAllEntities(entity, entities)

	viewData := template.NewTemplateDataWithRequestData(ec.RequestData())
	viewData["Entities"] = entities
	viewData["EditEntityType"] = ec.Env.MessageLocator.GetString("entity." + entityName)
	if entity == "consumptiongroup" {
		viewData["EmbeddedEntityType"] = ec.Env.MessageLocator.GetString("entity." + "energymetermapping")
	}
	if err != nil {
		viewData["ErrorMessage"] = err.Error()
	}
	return c.Render(http.StatusOK, entityName, viewData)
}

// HandleEntityDelete ...
func HandleEntityDelete(c echo.Context) error {
	ec := c.(*env.AppEnvContext)

	entity := c.Param("entity")
	entityName := ec.Env.MessageLocator.GetString("entity." + entity)

	oID := c.FormValue("gkvObjId")
	ioID, _ := strconv.Atoi(oID)
	if ioID <= 0 {
		return errors.New(ec.Env.MessageLocator.GetMessageF("msg.error.entity.general", entity, "ID <= 0!"))
	}

	entityModel, err := ec.Env.EF.GetEntity(entity)
	if err != nil {
		return err
	}

	err = ec.Env.DS.DeleteEntityByID(entityModel, uint(ioID))

	baseResponse := &viewmodel2.BaseResponse{}
	if err != nil {
		baseResponse.IsError = true
		baseResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.error.entity.delete", entityName)
	} else {
		baseResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.success.entity.delete", entityName)
	}

	return c.JSON(http.StatusOK, baseResponse)
}

// HandleEntityEdit ...
func HandleEntityEdit(c echo.Context) error {
	oID := c.FormValue("gkvObjId")
	ioID, _ := strconv.Atoi(oID)

	ec := c.(*env.AppEnvContext)
	entity := ec.Param("entity")
	oEntityObject, err := ec.Env.EF.GetEntity(entity)
	if err != nil {
		return err
	}

	entityResponse := viewmodel2.NewEntityResponse(oEntityObject)
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

// HandleEntityNew ...
func HandleEntityNew(c echo.Context) error {
	ec := c.(*env.AppEnvContext)
	entity := ec.Param("entity")

	oEntityObject, err := ec.Env.EF.GetEntity(entity)
	if err != nil {
		return err
	}

	entityResponse := viewmodel2.NewEntityResponse(oEntityObject)
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

// HandleOptionListAjax ...
func HandleOptionListAjax(c echo.Context) error {
	//show entity list
	entity := c.Param("entity")
	entityType := genmodels.ParseEntityType(strings.ToLower(entity))

	ec := c.(*env.AppEnvContext)

	entityName := ec.Env.MessageLocator.GetString("entity." + entity)

	entityResponse := viewmodel2.NewEntityOptionsResponse(nil)
	var err error
	entityResponse.EntityOptions, err = ec.Env.Services.GetEntityOptions(entityType)
	if err == nil {
		entityResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.success.entity.optionlist", entityName)
	} else {

		entityResponse.IsError = true
		entityResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.error.entity.optionlist", entityName)
	}

	//on client entityResponse is received as javascript object, no JSON.parse is needed
	return c.JSON(http.StatusOK, entityResponse)
}

// HandleEntityListAjax ...
func HandleEntityListAjax(c echo.Context) error {
	//show entity list
	entityName := c.Param("entity")

	ec := c.(*env.AppEnvContext)
	entities, err := ec.Env.EF.GetEntitySlice(entityName)
	if err != nil {
		return err
	}
	entity, err := ec.Env.EF.GetEntity(entityName)
	if err != nil {
		return err
	}

	entityResponse := viewmodel2.NewEntityResponse(entities)
	entityDesc := ec.Env.MessageLocator.GetString("entity." + entityName)

	err = ec.Env.DS.GetAllEntities(entity, entities)

	if err != nil {
		entityResponse.IsError = true
		entityResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.error.entity.list", entityDesc)
	} else {
		entityResponse.Message = ec.Env.MessageLocator.GetMessageF("msg.success.entity.list", entityDesc)
	}

	//on client entityResponse is received as javascript object, no JSON.parse is needed
	return c.JSON(http.StatusOK, entityResponse)
}
