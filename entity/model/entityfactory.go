package entitymodel

type DoWithEntityFunc func(Entity, ...interface{})

type EntityFactory interface {
	GetEntity(entityName string) (interface{}, error)
	GetEntitySlice(entityName string) (interface{}, error)
	DoWithAllEntities(entityList interface{}, entityFunc DoWithEntityFunc, params ...interface{})
}
