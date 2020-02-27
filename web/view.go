package web

//NewViewData return view data map
func NewViewData(contextData ContextData) map[string]interface{} {
	vd := make(map[string]interface{})
	vd["Host"] = contextData.Host()
	vd["VersionInfo"] = "V1.0"
	vd["UserID"] = contextData.UserID()

	return vd
}
