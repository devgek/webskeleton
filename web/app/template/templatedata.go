package template

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/types"
	"github.com/devgek/webskeleton/web/app/msg"
	"github.com/devgek/webskeleton/web/app/request"
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
	vd["Messages"] = msg.GetMessageLocator()
	vd["ProjectName"] = config.ProjectName
	vd["ProjectTitle"] = config.ProjectTitle
	vd["VersionInfo"] = config.ProjectVersion
	vd["StartPage"] = config.StartPage
	//add types for handling in templates
	vd["EntityTypes"] = types.EntityTypes()
	vd["OrgTypes"] = types.OrgTypes()
	vd["RoleTypes"] = types.RoleTypes()
	vd["ContactTypes"] = types.ContactTypes()

	return vd
}
