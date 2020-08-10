package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"kahrersoftware.at/webskeleton/config"
	"kahrersoftware.at/webskeleton/models"
)

//HandleAPICreate ...
func HandleAPICreate(c echo.Context) error {
	ec := c.(*config.EnvContext)
	entity := ec.Param("entity")
	oEntityObject := ec.Env.EF.Get(entity)

	if err := c.Bind(oEntityObject); err != nil {
		return err
	}

	err := ec.Env.DS.CreateEntity(oEntityObject)

	if err == nil {
		return c.JSON(http.StatusOK, oEntityObject)
	}

	return c.JSON(http.StatusInternalServerError, err.Error())
}

//HandleAPICreateAll ...
func HandleAPICreateAll(c echo.Context) error {
	ec := c.(*config.EnvContext)
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

	return c.JSON(http.StatusInternalServerError, err.Error())
}
