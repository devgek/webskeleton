package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/devgek/webskeleton/models"
	webenv "github.com/devgek/webskeleton/web/env"
	"github.com/labstack/echo"
)

//HandleAPICreate ...
func HandleAPICreate(c echo.Context) error {
	ec := c.(*webenv.EnvContext)
	entity := ec.Param("entity")
	oEntityObject := ec.Env.EF.Get(entity)

	if err := c.Bind(oEntityObject); err != nil {
		return err
	}

	var err error
	if entity == "User" {
		user := oEntityObject.(*models.User)
		oEntityObject, err = ec.Env.Services.CreateUser(user.Name, user.Pass, user.Email, user.Role)
	} else {
		err = ec.Env.DS.CreateEntity(oEntityObject)
	}

	if err == nil {
		return c.JSON(http.StatusOK, oEntityObject)
	}

	msg := ec.Env.MessageLocator.GetMessageF(err.Error())
	log.Println("HandleAPICreate::", http.StatusInternalServerError, msg)
	return c.JSON(http.StatusInternalServerError, msg)
}

//HandleAPIEdit ...
func HandleAPIEdit(c echo.Context) error {
	ec := c.(*webenv.EnvContext)
	entity := ec.Param("entity")
	oEntityObject := ec.Env.EF.Get(entity)

	if err := c.Bind(oEntityObject); err != nil {
		return err
	}

	err := ec.Env.DS.SaveEntity(oEntityObject)

	if err == nil {
		return c.JSON(http.StatusOK, oEntityObject)
	}

	msg := ec.Env.MessageLocator.GetMessageF(err.Error())
	log.Println("HandleAPIEdit::", http.StatusInternalServerError, msg)
	return c.JSON(http.StatusInternalServerError, msg)
}

//HandleAPIDelete ...
func HandleAPIDelete(c echo.Context) error {
	ec := c.(*webenv.EnvContext)
	entity := ec.Param("entity")
	id := ec.Param("id")
	ioID, _ := strconv.Atoi(id)
	entityModel := ec.Env.EF.Get(entity)
	err := ec.Env.DS.DeleteEntityByID(entityModel, uint(ioID))

	if err == nil {
		return c.JSON(http.StatusOK, "Entity deleted")
	}

	msg := ec.Env.MessageLocator.GetMessageF(err.Error())
	log.Println("HandleAPIDelete::", http.StatusInternalServerError, msg)
	return c.JSON(http.StatusInternalServerError, msg)
}

//HandleAPICreateAll ...
func HandleAPICreateAll(c echo.Context) error {
	ec := c.(*webenv.EnvContext)
	entity := ec.Param("entity")
	oEntityObjects := ec.Env.EF.GetSlice(entity)

	if err := c.Bind(oEntityObjects); err != nil {
		return err
	}

	var err error
	switch entityType := oEntityObjects.(type) {
	case *[]models.Contact:
		for idx := range *entityType {
			err = ec.Env.DS.CreateEntity(&((*entityType)[idx]))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
	}

	if err == nil {
		return c.JSON(http.StatusOK, "Entities created.")
	}

	msg := ec.Env.MessageLocator.GetMessageF(err.Error())
	log.Println("HandleAPICreateAll::", http.StatusInternalServerError, msg)
	return c.JSON(http.StatusInternalServerError, msg)
}
