package apihandler

import (
	"errors"
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/dtos"
	"github.com/devgek/webskeleton/types"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/devgek/webskeleton/models"
	webenv "github.com/devgek/webskeleton/web/env"
	"github.com/labstack/echo"
)

//HandleAPIHealth ...
func HandleAPIHealth(c echo.Context) error {
	ec := c.(*webenv.EnvContext)

	vd := webenv.TData{}
	vd["Host"] = c.Request().Host
	vd["ProjectName"] = config.ProjectName
	vd["VersionInfo"] = config.ProjectVersion
	vd["API"] = ec.Env.Api
	vd["health"] = "ok"

	return c.JSON(http.StatusOK, vd)
}

//HandleAPILogin handles login to api and returns a JWT token
func HandleAPILogin(c echo.Context) error {
	log.Println("HandleApiLogin")
	//do the login
	loginData := dtos.LoginData{}
	if err := c.Bind(&loginData); err != nil {
		return err
	}

	ec := c.(*webenv.EnvContext)
	user, err := ec.Env.Services.LoginUser(loginData.User, loginData.Pass)
	if err != nil {
		// return echo.NewHTTPError(http.StatusUnauthorized)
		msg := ec.Env.MessageLocator.GetMessageF(err.Error())
		log.Println("HandleApiLogin return:", http.StatusUnauthorized, msg)
		return c.JSON(http.StatusUnauthorized, msg)
	}

	//login ok
	log.Println("User", user.Name, "logged in for api")

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims into webtoken, content can be checked on further requests with token
	claims := token.Claims.(jwt.MapClaims)
	isAdmin := (user.Role == types.RoleTypeAdmin)
	claims["name"] = loginData.User
	claims["admin"] = isAdmin

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": t,
		"name":  loginData.User,
		"admin": isAdmin,
	})
}

//HandleAPICreateEntity ...
func HandleAPICreateEntity(c echo.Context) error {
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
	log.Println("HandleAPICreateEntity::", http.StatusInternalServerError, apiError, origError)
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

//HandleAPIUpdateEntity ...
func HandleAPIUpdateEntity(c echo.Context) error {
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
	log.Println("HandleAPIUpdateEntity::", http.StatusInternalServerError, apiError, origError)
	return c.JSON(http.StatusInternalServerError, apiError)
}

//HandleAPIDeleteEntity ...
func HandleAPIDeleteEntity(c echo.Context) error {
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
	log.Println("HandleAPIDeleteEntity::", http.StatusInternalServerError, apiError, origError)
	return c.JSON(http.StatusInternalServerError, apiError)
}

//HandleAPIOptionList ...
func HandleAPIOptionList(c echo.Context) error {
	//show entity list
	entity := c.Param("entity")
	entityType := types.ParseEntityType(strings.ToLower(entity))

	ec := c.(*webenv.EnvContext)

	entityOptions, origError := ec.Env.Services.GetEntityOptions(entityType)
	if origError == nil {
		return c.JSON(http.StatusOK, entityOptions)
	}

	apiError := &dtos.ApiError{Nr: 6000, Msg: "No entity options"}
	log.Println("HandleAPIOptionList::", http.StatusInternalServerError, apiError, origError)
	return c.JSON(http.StatusInternalServerError, apiError)
}

//HandleAPIEntityList ...
func HandleAPIEntityList(c echo.Context) error {
	//show entity list
	entity := c.Param("entity")

	ec := c.(*webenv.EnvContext)
	entities, origError := ec.Env.EF.GetSlice(entity)

	if origError == nil {
		origError = ec.Env.DS.GetAllEntities(entities)
		if origError == nil {
			return c.JSON(http.StatusOK, entities)
		}
	}

	apiError := &dtos.ApiError{Nr: 5000, Msg: "No entity list"}
	log.Println("HandleAPIEntityList::", http.StatusInternalServerError, apiError, origError)
	return c.JSON(http.StatusInternalServerError, apiError)
}
