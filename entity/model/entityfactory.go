package entitymodel

type DoWithEntityFunc func(Entity, ...interface{})

type EntityFactory interface {
	RegisterType(entityName string, entityType interface{})
	RegisterSliceType(entityName string, entitySliceType interface{})
	GetEntity(entityName string) (interface{}, error)
	GetEntitySlice(entityName string) (interface{}, error)
	DoWithAllEntities(entityList interface{}, entityFunc DoWithEntityFunc, params ...interface{})
}
type entityTypeRegistry map[string]interface{}

type DefaultEntityFactory struct {
	entityRegistry      entityTypeRegistry
	entitySliceRegistry entityTypeRegistry
}

func NewDefaultEntityFactory() *DefaultEntityFactory {
	return &DefaultEntityFactory{make(map[string]interface{}), make(map[string]interface{})}
}

func (ef DefaultEntityFactory) RegisterType(entityName string, entityType interface{}) {
	ef.entityRegistry[entityName] = entityType
}

func (ef DefaultEntityFactory) RegisterSliceType(entityName string, entitySliceType interface{}) {
	ef.entitySliceRegistry[entityName] = entitySliceType
}

//GetEntity return entity struct by name
func (ef DefaultEntityFactory) GetEntity(entityName string) (interface{}, error) {
	entity := ef.entityRegistry[entityName]
	return entity, nil
}

//GetEntitySlice return slice of entity struct by name
func (ef DefaultEntityFactory) GetEntitySlice(entityName string) (interface{}, error) {
	entitySlice := ef.entitySliceRegistry[entityName]
	return entitySlice, nil
}

//DoWithAllEntities
/*
	Method ranges over entities and calls entityFunc with each entity. You can serve parameters with each call to entityFunc.
    Attention! Maybe params should be pointers to change things outside entityFunc.
*/
func (ef DefaultEntityFactory) DoWithAllEntities(entityList interface{}, entityFunc DoWithEntityFunc, params ...interface{}) {
	if val, ok := entityList.([]interface{}); ok {
		for _, e := range val {
			if entity, ok := e.(Entity); ok {
				entityFunc(entity, params...)
			}
		}
	}
}
