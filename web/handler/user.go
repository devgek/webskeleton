package handler

import (
	"log"
	"net/http"

	"github.com/devgek/webskeleton/models"
	webenv "github.com/devgek/webskeleton/web/env"
	"github.com/devgek/webskeleton/web/viewmodel"
	"github.com/labstack/echo"
)

//HandleCreateUser own handler instead of using HandleEntityNew
func HandleCreateUser(c echo.Context) error {
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
