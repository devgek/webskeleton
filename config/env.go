package config

import (
	"github.com/devgek/webskeleton/msg"
	"log"

	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/services"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm for sqlite3
)

//
var (
	ProjectName    = "webskeleton"
	ProjectTitle   = "go-webskeleton"
	ProjectVersion = "V1.0"
	DatabaseName   = "webskeleton.db"
)

//Env the environment
type Env struct {
	DS             data.Datastore
	Services       *services.Services
	MessageLocator *msg.MessageLocator
}

//InitEnv return new initialized environment
func InitEnv() *Env {
	//load messages
	ml := msg.NewMessageLocator()

	//here we decide the database system
	ds, err := data.NewDatastore("sqlite3", DatabaseName)
	if err != nil {
		log.Panic(err)
	}

	services := services.NewServices(ds)

	return &Env{DS: ds, Services: services, MessageLocator: ml}
}
