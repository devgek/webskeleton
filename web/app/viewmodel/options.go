package viewmodel

import (
	"github.com/devgek/webskeleton/entity/dto"
)

//EntityOptionsResponse ...
type EntityOptionsResponse struct {
	*BaseResponse
	EntityOptions []entitydto.EntityOption
}

//NewEntityOptionsResponse ...
func NewEntityOptionsResponse(entityOptions []entitydto.EntityOption) *EntityOptionsResponse {
	return &EntityOptionsResponse{&BaseResponse{false, ""}, entityOptions}
}
