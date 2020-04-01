package data

import (
	"errors"
)

//
var (
	ErrorEntityNotDeleted = errors.New("Entity not deleted")
)

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
