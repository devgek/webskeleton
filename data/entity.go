package data

import (
	"errors"
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

	return nil
}

//GetAllEntities select * from table
func (ds *DatastoreImpl) GetAllEntities(entitySlice interface{}) error {
	return ds.Find(entitySlice).Error
}

//CreateEntity insert into entity table
func (ds *DatastoreImpl) CreateEntity(entity interface{}) error {
	return ds.Create(entity).Error
}

//SaveEntity update entity table
func (ds *DatastoreImpl) SaveEntity(entity interface{}) error {
	return ds.Save(entity).Error
}

//DeleteEntityByID delete entity by id (primary key)
//ID must be provided
//Attention ds is not the same as db!
func (ds *DatastoreImpl) DeleteEntityByID(entity interface{}) error {
	db := ds.Unscoped().Delete(entity)

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected != 1 {
		return ErrorEntityNotDeleted
	}

	return nil
}
