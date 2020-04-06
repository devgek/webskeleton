package viewmodel

//EntityResponse ...
type EntityResponse struct {
	*BaseResponse
	EntityObject interface{}
}

//NewEntityResponse ...
func NewEntityResponse(entity interface{}) *EntityResponse {
	return &EntityResponse{&BaseResponse{false, ""}, entity}
}
