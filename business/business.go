package business

import "kahrersoftware.at/webskeleton/data"

//Business the business logic
type Business struct {
	DS data.Datastore
}

//NewBusiness ...
func NewBusiness(ds data.Datastore) (*Business, error) {
	return &Business{ds}, nil
}

//NewContext create new business context
func (b *Business) NewContext() *Context {
	return &Context{b.DS}
}
