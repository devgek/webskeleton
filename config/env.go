package config

import (
	"log"

	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/services"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm for sqlite3
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
