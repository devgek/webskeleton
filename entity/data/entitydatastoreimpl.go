package entitydata

import (
	"errors"
	entitymodel "github.com/devgek/webskeleton/entity/model"

	"github.com/devgek/webskeleton/models"
	"gorm.io/gorm"
)

var (
	ErrorEntityNotFountBy = errors.New("Entity with given where condition not found")
	ErrorEntityNotDeleted = errors.New("Entity not deleted")
)

// GormEntityDatastore the EntityDatastore implementation using gorm.DB for database operations
type GormEntityDatastore struct {
	*gorm.DB
}

// GetOneEntityBy select * from table where key = value
func (ds *GormEntityDatastore) GetOneEntityBy(entity interface{}, key string, val interface{}) error {
	if err := ds.Where(key+" = ?", val).First(entity).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrorEntityNotFountBy
		}

		return err
	}

	return ds.LoadRelatedEntities(entity)
}

// GetEntityByID ...
func (ds *GormEntityDatastore) GetEntityByID(entity interface{}, id uint) error {
	if err := ds.First(entity, id).Error; err != nil {
		return err
	}

	return ds.LoadRelatedEntities(entity)
}

// GetAllEntities select * from table
func (ds *GormEntityDatastore) GetAllEntities(entitySlice interface{}) error {
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

// CreateEntity insert into entity table
func (ds *GormEntityDatastore) CreateEntity(entity interface{}) error {
	// return ds.Create(entity).Error
	if err := ds.Create(entity).Error; err != nil {
		return err
	}
	return ds.LoadRelatedEntities(entity)
}

// SaveEntity update entity table
func (ds *GormEntityDatastore) SaveEntity(entity interface{}) error {
	// return ds.Save(entity).Error
	if err := ds.Save(entity).Error; err != nil {
		return err
	}
	return ds.LoadRelatedEntities(entity)
}

// DeleteEntityByID delete entity by id (primary key)
// ID must be provided
// Attention ds is not the same as db!
func (ds *GormEntityDatastore) DeleteEntityByID(entity interface{}, id uint) error {
	db := ds.Unscoped().Delete(entity, id)

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected != 1 {
		return ErrorEntityNotDeleted
	}

	return nil
}

// LoadRelatedEntities load embedded entities
func (ds *GormEntityDatastore) LoadRelatedEntities(i interface{}) error {
	if val, ok := i.(entitymodel.EntityHolder); ok {
		if val != nil {
			return val.LoadRelated(ds.DB)
		}
	}

	return nil
}
