package handler

import (
	"github.com/devgek/webskeleton/web/app/env"
	"github.com/devgek/webskeleton/web/app/template"
	"log"
	"net/http"
	"strconv"

	"github.com/devgek/webskeleton/helper/helper"
	"github.com/devgek/webskeleton/types"
	"github.com/labstack/echo"
)

//HandlePage1 ...
func HandlePage1(c echo.Context) error {
	contactType := helper.ValueOrDefault(c.FormValue("page1FilterContactType"), "0")
	iContactType, err := strconv.Atoi(contactType)
	helper.PanicOnError(err)
	log.Println("page1FilterContactType", iContactType)

	ec := c.(*env.AppEnvContext)

	viewData := template.NewTemplateDataWithRequestData(ec.RequestData())
	viewData["FilterContactType"] = contactType
	viewData["ContactTypes"] = types.ContactTypes()

	return c.Render(http.StatusOK, "page1", viewData)
}
