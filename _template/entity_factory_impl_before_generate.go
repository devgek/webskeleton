/*
	Type EntityFactoryImpl must exist even before first generating models, because they are used
	in appenv.go and apienv.go
*/
package generated_models

import (
	"errors"

	entitymodel "github.com/devgek/webskeleton/entity/model"
)

//EntityFactoryImpl create Entities by name
type EntityFactoryImpl struct {
}

//Get return entity struct by name
func (ef EntityFactoryImpl) Get(entityName string) (interface{}, error) {
	return nil, errors.New("Unknown entity '" + entityName + "'")
}

//GetSlice return slice of entity struct by name
func (ef EntityFactoryImpl) GetSlice(entityName string) (interface{}, error) {
	return nil, errors.New("Unknown entity '" + entityName + "'")
}

//DoWithAll
/*
	Method ranges over entities and calls entityFunc with each entity. You can serve parameters with each call to entityFunc.
    Attention! Maybe params should be pointers to change things outside entityFunc.
*/
func (ef EntityFactoryImpl) DoWithAll(entityList interface{}, entityFunc entitymodel.DoWithEntityFunc, params ...interface{}) {
}
