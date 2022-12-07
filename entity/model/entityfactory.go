package entitymodel

type DoWithEntityFunc func(Entity, ...interface{})

type EntityFactory interface {
	Get(entityName string) (interface{}, error)
	GetSlice(entityName string) (interface{}, error)
	DoWithAll(entityList interface{}, entityFunc DoWithEntityFunc, params ...interface{})
}
