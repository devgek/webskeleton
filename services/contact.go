package services

import (
	"github.com/devgek/webskeleton/models"
	"log"
)

//GetAllContacts ...
func (s Services) GetAllContacts() ([]models.Contact, error) {
	var contacts = []models.Contact{}
	err := s.DS.GetAllEntities(&contacts)
	if err == nil {
		return contacts, err
	}

	log.Println("GetAllContacts:", err.Error())
	return contacts, err
}
