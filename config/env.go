package config

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm for sqlite3
	"kahrersoftware.at/webskeleton/data"
)

//Env the environment
type Env struct {
	DS data.Datastore
}

//InitEnv return new initialized environment
func InitEnv() *Env {
	//here we decide the database system
	ds, err := data.NewDatastore("sqlite3", "webskeleton.db")
	if err != nil {
		log.Panic(err)
	}

	return &Env{DS: ds}
}
