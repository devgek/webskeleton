package config

import (
	"github.com/devgek/webskeleton/global"
	"github.com/devgek/webskeleton/types"
	"github.com/devgek/webskeleton/web/request"
)

//TData map holding data for page templates
type TData map[string]interface{}

//NewTemplateDataWithRequestData return view data map filled with context data
func NewTemplateDataWithRequestData(requestData request.RData) TData {
	vd := NewTemplateData()

	vd["UserID"] = requestData.UserID()
	vd["Admin"] = requestData.IsAdmin()

	return vd
}

//NewTemplateData ...
func NewTemplateData() TData {
	vd := make(map[string]interface{})
	vd["Messages"] = GetWebEnv().MessageLocator
	vd["ProjectName"] = global.ProjectName
	vd["VersionInfo"] = global.ProjectVersion
	vd["StartPage"] = global.StartPage
	//add types for handling in templates
	vd["EntityTypes"] = types.EntityTypes()
	vd["OrgTypes"] = types.OrgTypes()
	vd["RoleTypes"] = types.RoleTypes()
	vd["ContactTypes"] = types.ContactTypes()

	return vd
}
