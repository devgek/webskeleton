package entitydata

import (
	"errors"

	entitymodel "github.com/devgek/webskeleton/entity/model"
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

func NewGormEntityDatastore(db *gorm.DB) *GormEntityDatastore {
	return &GormEntityDatastore{db}
}

// GetOneEntityBy select * from table where key = value
func (ds *GormEntityDatastore) GetOneEntityBy(entity interface{}, key string, val interface{}) error {
	db := ds.LoadRelatedEntities(entity)
	if err := db.Where(key+" = ?", val).First(entity).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrorEntityNotFountBy
		}

		return err
	}

	return nil
}

// GetEntityByID ...
func (ds *GormEntityDatastore) GetEntityByID(entity interface{}, id uint) error {
	db := ds.LoadRelatedEntities(entity)
	if err := db.First(entity, id).Error; err != nil {
		return err
	}

	return nil
}

// GetAllEntities select * from table
func (ds *GormEntityDatastore) GetAllEntities(entity interface{}, entitySlice interface{}) error {
	db := ds.LoadRelatedEntities(entity)
	if err := db.Order("id").Find(entitySlice).Error; err != nil {
		return err
	}

	return nil
}

// CreateEntity insert into entity table
func (ds *GormEntityDatastore) CreateEntity(entity interface{}) error {
	// return ds.Create(entity).Error
	if err := ds.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

// SaveEntity update entity table
func (ds *GormEntityDatastore) SaveEntity(entity interface{}) error {
	// return ds.Save(entity).Error
	if err := ds.Save(entity).Error; err != nil {
		return err
	}
	return nil
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
func (ds *GormEntityDatastore) LoadRelatedEntities(i interface{}) *gorm.DB {
	db := ds.DB
	if val, ok := i.(entitymodel.Entity); ok {
		if val != nil {
			embeds := val.MustEmbed()
			for _, embed := range embeds {
				db = ds.Preload(embed)
			}
		}
	}

	return db
}
