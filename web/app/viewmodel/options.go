package viewmodel

import (
	"github.com/devgek/webskeleton/dtos"
)

//EntityOptionsResponse ...
type EntityOptionsResponse struct {
	*BaseResponse
	EntityOptions []dtos.EntityOption
}

//NewEntityOptionsResponse ...
func NewEntityOptionsResponse(entityOptions []dtos.EntityOption) *EntityOptionsResponse {
	return &EntityOptionsResponse{&BaseResponse{false, ""}, entityOptions}
}
