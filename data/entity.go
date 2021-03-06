package data

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

//GetOneEntityBy select * from table where key = value
func (ds *DatastoreImpl) GetOneEntityBy(entity interface{}, key string, val interface{}) error {
	if err := ds.Where(key+" = ?", val).First(entity).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrorEntityNotFountBy
		}

		return err
	}

	return ds.LoadRelated(entity)
}

//GetEntityByID ...
func (ds *DatastoreImpl) GetEntityByID(entity interface{}, id uint) error {
	if err := ds.First(entity, id).Error; err != nil {
		return err
	}

	return ds.LoadRelated(entity)
}

//GetAllEntities select * from table
func (ds *DatastoreImpl) GetAllEntities(entitySlice interface{}) error {
	if err := ds.Order("id").Find(entitySlice).Error; err != nil {
		return err
	}

	switch entityType := entitySlice.(type) {
	case *[]models.User:
		for idx := range *entityType {
			ds.LoadRelated(&((*entityType)[idx]))
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
func (ds *DatastoreImpl) CreateEntity(entity interface{}) error {
	// return ds.Create(entity).Error
	if err := ds.Create(entity).Error; err != nil {
		return err
	}
	return ds.LoadRelated(entity)
}

//SaveEntity update entity table
func (ds *DatastoreImpl) SaveEntity(entity interface{}) error {
	// return ds.Save(entity).Error
	if err := ds.Save(entity).Error; err != nil {
		return err
	}
	return ds.LoadRelated(entity)
}

//DeleteEntityByID delete entity by id (primary key)
//ID must be provided
//Attention ds is not the same as db!
func (ds *DatastoreImpl) DeleteEntityByID(entity interface{}, id uint) error {
	db := ds.Unscoped().Delete(entity, id)

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected != 1 {
		return ErrorEntityNotDeleted
	}

	return nil
}

//LoadRelated load embedded entities
func (ds *DatastoreImpl) LoadRelated(i interface{}) error {
	if val, ok := i.(models.EntityHolder); ok {
		if val != nil {
			return val.LoadRelated(ds.DB)
		}
	}

	return nil
}
