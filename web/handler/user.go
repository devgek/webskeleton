package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"kahrersoftware.at/webskeleton/config"
	"kahrersoftware.at/webskeleton/models"
	"kahrersoftware.at/webskeleton/web/viewmodel"
)

//HandleCreateUser own handler instead of using HandleEntityNew
func HandleCreateUser(c echo.Context) error {
	ec := c.(*config.EnvContext)
	entity := ec.Param("entity")
	oEntityObject := ec.Env.EF.Get(entity)

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

	var err error
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
