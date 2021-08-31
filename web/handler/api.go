package handler

import (
	"errors"
	"github.com/devgek/webskeleton/dtos"
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

	oEntityObject, origError := ec.Env.EF.Get(entity)
	if origError == nil {
		origError = c.Bind(oEntityObject)
		if origError == nil {
			if entity == "User" {
				user := oEntityObject.(*models.User)
				oEntityObject, origError = ec.Env.Services.CreateUser(user.Name, user.Pass, user.Email, user.Role)
			} else {
				origError = ec.Env.DS.CreateEntity(oEntityObject)
			}

			if origError == nil {
				return c.JSON(http.StatusOK, oEntityObject)
			}
		}
	}

	apiError := &dtos.ApiError{Nr: 1000, Msg: "Entity not created"}
	log.Println("HandleAPICreate::", http.StatusInternalServerError, apiError, origError)
	return c.JSON(http.StatusInternalServerError, apiError)
}

//HandleAPICreateAll ...
func HandleAPICreateAll(c echo.Context) error {
	ec := c.(*webenv.EnvContext)
	entity := ec.Param("entity")
	oEntityObjects, origError := ec.Env.EF.GetSlice(entity)
	if origError == nil {
		origError = c.Bind(oEntityObjects)
		if origError == nil {
			switch entityType := oEntityObjects.(type) {
			case *[]models.Contact:
				for idx := range *entityType {
					origError = ec.Env.DS.CreateEntity(&((*entityType)[idx]))
					if origError != nil {
						goto errorReturn
					}
				}
				return c.JSON(http.StatusOK, "Entities created.")
			}
			origError = errors.New("HandleAPICreateAll not valid for entity " + entity)
		}
	}
errorReturn:
	apiError := &dtos.ApiError{Nr: 1100, Msg: "EntityList not created"}
	log.Println("HandleAPICreateAll::", http.StatusInternalServerError, apiError, origError)
	return c.JSON(http.StatusInternalServerError, apiError)
}

//HandleAPIUpdate ...
func HandleAPIUpdate(c echo.Context) error {
	ec := c.(*webenv.EnvContext)
	entity := ec.Param("entity")
	oEntityObject, origError := ec.Env.EF.Get(entity)
	if origError == nil {
		origError = c.Bind(oEntityObject)
		if origError == nil {
			origError = ec.Env.DS.SaveEntity(oEntityObject)
			if origError == nil {
				return c.JSON(http.StatusOK, oEntityObject)
			}
		}
	}

	apiError := &dtos.ApiError{Nr: 3000, Msg: "Entity not updated"}
	log.Println("HandleAPIUpdate::", http.StatusInternalServerError, apiError, origError)
	return c.JSON(http.StatusInternalServerError, apiError)
}

//HandleAPIDelete ...
func HandleAPIDelete(c echo.Context) error {
	ec := c.(*webenv.EnvContext)
	entity := ec.Param("entity")
	id := ec.Param("id")
	ioID, origError := strconv.Atoi(id)
	if origError == nil {
		entityModel, origError := ec.Env.EF.Get(entity)
		if origError == nil {
			origError = ec.Env.DS.DeleteEntityByID(entityModel, uint(ioID))
			if origError == nil {
				return c.JSON(http.StatusOK, "Entity deleted")
			}
		}
	}

	apiError := &dtos.ApiError{Nr: 4000, Msg: "Entity not deleted"}
	log.Println("HandleAPIDelete::", http.StatusInternalServerError, apiError, origError)
	return c.JSON(http.StatusInternalServerError, apiError)
}
