package config

import (
	"log"
	"net/http"
	"sync"

	"github.com/gobuffalo/packr/v2"
	"kahrersoftware.at/webskeleton/global"
	"kahrersoftware.at/webskeleton/models"
	"kahrersoftware.at/webskeleton/msg"
	"kahrersoftware.at/webskeleton/web/template"

	_ "github.com/jinzhu/gorm/dialects/postgres" // gorm for postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // gorm for sqlite3
	"kahrersoftware.at/webskeleton/data"
	"kahrersoftware.at/webskeleton/services"
)

//Env the environment
type Env struct {
	TStore         template.TStore
	Templates      *packr.Box
	Assets         http.FileSystem
	DS             data.Datastore
	Services       *services.Services
	MessageLocator *msg.MessageLocator
	EF             models.EntityFactory
}

var once sync.Once

//WebEnv singleton instance for the app env
var webEnv *Env

//GetWebEnv return new initialized environment
func GetWebEnv() *Env {
	once.Do(func() {
		// ../web/templates important for packr2 to find files
		origninalTemplateBox := packr.New("templates", "../web/templates")
		// templateBox := packrfix.New(origninalTemplateBox)

		//init TStore
		tStore := template.NewBoxBasedTemplateStore(origninalTemplateBox)

		origninalAssetBox := packr.New("assets", "../web/assets")
		// assetBox := packrfix.New(origninalAssetBox)

		//load locale specific message file, if not default
		// messages, err := assetBox.Find("msg/messages-en.yaml")
		messages, err := origninalAssetBox.Find("msg/messages.yaml")

		//load messages
		ml := msg.NewMessageLocator(messages)

		//here we create the datastore
		//?_foreign_keys=1, neccessary for golang to respect foreign key constraints on sqlite3 db
		var ds data.Datastore
		if global.DatastoreSystem() == "postgres" {
			ds, err = data.NewPostgres()
		} else {
			ds, err = data.NewSqlite(global.DatabaseName)
		}
		if err != nil {
			log.Panic(err)
		}

		services := services.NewServices(ds)

		webEnv = &Env{TStore: tStore, Templates: origninalTemplateBox, Assets: origninalAssetBox, DS: ds, Services: services, MessageLocator: ml}
	})

	return webEnv
}
