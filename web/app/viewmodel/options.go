package viewmodel

import (
	"github.com/devgek/webskeleton/entity/dto"
)

//EntityOptionsResponse ...
type EntityOptionsResponse struct {
	*BaseResponse
	EntityOptions []dto.EntityOption
}

//NewEntityOptionsResponse ...
func NewEntityOptionsResponse(entityOptions []dto.EntityOption) *EntityOptionsResponse {
	return &EntityOptionsResponse{&BaseResponse{false, ""}, entityOptions}
}
