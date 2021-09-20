package apihandler

import (
	"errors"
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/dtos"
	"github.com/devgek/webskeleton/types"
	"github.com/devgek/webskeleton/web/api/env"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/devgek/webskeleton/models"
	"github.com/labstack/echo"
)

//HandleAPIHealth ...
func HandleAPIHealth(c echo.Context) error {
	vd := make(map[string]interface{})
	vd["Host"] = c.Request().Host
	vd["ProjectName"] = config.ProjectName
	vd["VersionInfo"] = config.ProjectVersion
	vd["API"] = true
	vd["health"] = "ok"

	return c.JSON(http.StatusOK, vd)
}

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

//HandleAPILogin handles login to api and returns a JWT token
func HandleAPILogin(c echo.Context) error {
	log.Println("HandleApiLogin")
	//do the login
	loginData := dtos.LoginData{}
	if err := c.Bind(&loginData); err != nil {
		return err
	}

	ec := c.(*env.ApiEnvContext)
	user, err := ec.ApiEnv.Services.LoginUser(loginData.User, loginData.Pass)
	if err != nil {
		log.Println("HandleApiLogin return:", http.StatusUnauthorized, err.Error())
		return c.JSON(http.StatusUnauthorized, "Login not allowed.")
	}

	//login ok
	log.Println("User", user.Name, "logged in for api")

	// Set claims; content can be checked on further requests with token
	isAdmin := user.Role == types.RoleTypeAdmin
	claims := &jwtCustomClaims{
		loginData.User,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create token with custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tSigned, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": tSigned,
		"name":  loginData.User,
		"admin": isAdmin,
	})
}

//HandleAPICreateEntity ...
func HandleAPICreateEntity(c echo.Context) error {
	ec := c.(*env.ApiEnvContext)
	entity := ec.Param("entity")

	oEntityObject, origError := ec.ApiEnv.EF.Get(entity)
	if origError == nil {
		origError = c.Bind(oEntityObject)
		if origError == nil {
			if entity == "User" {
				user := oEntityObject.(*models.User)
				oEntityObject, origError = ec.ApiEnv.Services.CreateUser(user.Name, user.Pass, user.Email, user.Role)
			} else {
				origError = ec.ApiEnv.DS.CreateEntity(oEntityObject)
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
	ec := c.(*env.ApiEnvContext)
	entity := ec.Param("entity")
	oEntityObjects, origError := ec.ApiEnv.EF.GetSlice(entity)
	if origError == nil {
		origError = c.Bind(oEntityObjects)
		if origError == nil {
			switch entityType := oEntityObjects.(type) {
			case *[]models.Contact:
				for idx := range *entityType {
					origError = ec.ApiEnv.DS.CreateEntity(&((*entityType)[idx]))
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
	ec := c.(*env.ApiEnvContext)
	entity := ec.Param("entity")
	oEntityObject, origError := ec.ApiEnv.EF.Get(entity)
	if origError == nil {
		origError = c.Bind(oEntityObject)
		if origError == nil {
			origError = ec.ApiEnv.DS.SaveEntity(oEntityObject)
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
	ec := c.(*env.ApiEnvContext)
	entity := ec.Param("entity")
	id := ec.Param("id")
	ioID, origError := strconv.Atoi(id)
	if origError == nil {
		entityModel, origError := ec.ApiEnv.EF.Get(entity)
		if origError == nil {
			origError = ec.ApiEnv.DS.DeleteEntityByID(entityModel, uint(ioID))
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
	entityType := models.ParseEntityType(strings.ToLower(entity))

	ec := c.(*env.ApiEnvContext)

	entityOptions, origError := ec.ApiEnv.Services.GetEntityOptions(entityType)
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

	ec := c.(*env.ApiEnvContext)
	entities, origError := ec.ApiEnv.EF.GetSlice(entity)

	if origError == nil {
		origError = ec.ApiEnv.DS.GetAllEntities(entities)
		if origError == nil {
			return c.JSON(http.StatusOK, entities)
		}
	}

	apiError := &dtos.ApiError{Nr: 5000, Msg: "No entity list"}
	log.Println("HandleAPIEntityList::", http.StatusInternalServerError, apiError, origError)
	return c.JSON(http.StatusInternalServerError, apiError)
}
