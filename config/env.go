package config

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm for sqlite3
	"kahrersoftware.at/webskeleton/data"
	"kahrersoftware.at/webskeleton/services"
)

//Env the environment
type Env struct {
	DS       data.Datastore
	Services *services.Services
}

//InitEnv return new initialized environment
func InitEnv() *Env {
	//here we decide the database system
	ds, err := data.NewDatastore("sqlite3", "webskeleton.db")
	if err != nil {
		log.Panic(err)
	}

	services := services.NewServices(ds)

	return &Env{DS: ds, Services: services}
}
