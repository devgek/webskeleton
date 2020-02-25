package services

import "kahrersoftware.at/webskeleton/data"

//Services the business serives
type Services struct {
	DS data.Datastore
}

//NewServices ...
func NewServices(ds data.Datastore) *Services {
	return &Services{ds}
}
