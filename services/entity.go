package services

import (
	"log"
)

//GetEntities ...
func (s Services) GetEntities(entityName string) (interface{}, error) {
	entities := s.EF.GetSlice(entityName)

	err := s.DS.GetAllEntities(entities)
	if err == nil {
		return entities, err
	}

	log.Println("GetAllEntities:", err.Error())
	return entities, err
}
