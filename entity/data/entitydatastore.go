package entitydata

//EntityDatastore CRUD operations with abstract entity type using gorm.DB
type EntityDatastore interface {
	GetOneEntityBy(entity interface{}, key string, val interface{}) error
	GetEntityByID(entity interface{}, id uint) error
	GetAllEntities(entitySlice interface{}) error
	CreateEntity(entity interface{}) error
	SaveEntity(entity interface{}) error
	DeleteEntityByID(entity interface{}, id uint) error
}
