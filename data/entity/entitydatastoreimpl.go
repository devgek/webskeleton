package entitydata

import (
	"errors"

	"github.com/devgek/webskeleton/models"
	"github.com/jinzhu/gorm"
)

//
var (
	ErrorEntityNotFountBy = errors.New("Entity with given where condition not found")
	ErrorEntityNotDeleted = errors.New("Entity not deleted")
)

//GormEntityDatastoreImpl the EntityDatastore implementation using gorm.DB for database operations
type GormEntityDatastoreImpl struct {
	*gorm.DB
}

//GetOneEntityBy select * from table where key = value
func (ds *GormEntityDatastoreImpl) GetOneEntityBy(entity interface{}, key string, val interface{}) error {
	if err := ds.Where(key+" = ?", val).First(entity).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrorEntityNotFountBy
		}

		return err
	}

	return ds.LoadRelatedEntities(entity)
}

//GetEntityByID ...
func (ds *GormEntityDatastoreImpl) GetEntityByID(entity interface{}, id uint) error {
	if err := ds.First(entity, id).Error; err != nil {
		return err
	}

	return ds.LoadRelatedEntities(entity)
}

//GetAllEntities select * from table
func (ds *GormEntityDatastoreImpl) GetAllEntities(entitySlice interface{}) error {
	if err := ds.Order("id").Find(entitySlice).Error; err != nil {
		return err
	}

	switch entityType := entitySlice.(type) {
	case *[]models.User:
		for idx := range *entityType {
			ds.LoadRelatedEntities(&((*entityType)[idx]))
		}
	case *[]models.Contact:
		for idx := range *entityType {
			ds.LoadRelatedEntities(&((*entityType)[idx]))
		}
	}

	/*
		es := entitySlice.(*[]models.ConsumptionGroup)
		//Attention!! use index instead of value because range makes a copy of the value
		for idx := range *es {
			ds.LoadRelated(&((*es)[idx]))
		}

					if reflect.TypeOf(entitySlice).Kind() == reflect.Ptr {
					s := reflect.Indirect(reflect.ValueOf(entitySlice))

					if s.Kind() == reflect.Slice {
						for i := 0; i < s.Len(); i++ {
							valPtr := s.Index(i)
							val := reflect.Indirect(valPtr)
							if err := ds.LoadRelated(&val); err != nil {
								return err
							}
						}
					}

				}

			if reflect.TypeOf(entitySlice).Kind() == reflect.Ptr {
				s := reflect.Indirect(reflect.ValueOf(entitySlice))

				if s.Kind() == reflect.Slice {
					for i := 0; i < s.Len(); i++ {
						valPtr := s.Index(i)
						val := reflect.Indirect(valPtr)
						t := val.Interface().(models.EntityHolder)
						ds.LoadRelated(t)
					}
				}

			}
	*/

	return nil
}

//CreateEntity insert into entity table
func (ds *GormEntityDatastoreImpl) CreateEntity(entity interface{}) error {
	// return ds.Create(entity).Error
	if err := ds.Create(entity).Error; err != nil {
		return err
	}
	return ds.LoadRelatedEntities(entity)
}

//SaveEntity update entity table
func (ds *GormEntityDatastoreImpl) SaveEntity(entity interface{}) error {
	// return ds.Save(entity).Error
	if err := ds.Save(entity).Error; err != nil {
		return err
	}
	return ds.LoadRelatedEntities(entity)
}

//DeleteEntityByID delete entity by id (primary key)
//ID must be provided
//Attention ds is not the same as db!
func (ds *GormEntityDatastoreImpl) DeleteEntityByID(entity interface{}, id uint) error {
	db := ds.Unscoped().Delete(entity, id)

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected != 1 {
		return ErrorEntityNotDeleted
	}

	return nil
}

//LoadRelatedEntities load embedded entities
func (ds *GormEntityDatastoreImpl) LoadRelatedEntities(i interface{}) error {
	if val, ok := i.(models.EntityHolder); ok {
		if val != nil {
			return val.LoadRelated(ds.DB)
		}
	}

	return nil
}
