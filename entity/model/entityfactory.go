package entitymodel

type DoWithEntityFunc func(EntityOptionBuilder, ...interface{})

type EntityFactory interface {
	Get(entityName string) (interface{}, error)
	GetSlice(entityName string) (interface{}, error)
	DoWithAll(entityList interface{}, entityFunc DoWithEntityFunc, params ...interface{})
}
