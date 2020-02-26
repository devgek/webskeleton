package web

import "net/http"

//NewViewData return view data map
func NewViewData(r *http.Request) map[string]interface{} {
	vd := make(map[string]interface{})
	vd["Host"] = r.Host
	vd["VersionInfo"] = "V1.0"
	if contextData, ok := FromContext(r.Context()); ok {
		vd["UserID"] = contextData.UserID()
	}

	return vd
}
