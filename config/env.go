package config

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm for sqlite3
	"kahrersoftware.at/webskeleton/models"
)

//Env the environment
type Env struct {
	ds models.Datastore
}

//InitEnv return new initialized environment
func InitEnv() *Env {
	//here we decide the database system
	ds, err := models.NewDS("sqlite3", "webskeleton.db")
	if err != nil {
		log.Panic(err)
	}

	return &Env{ds}
}
