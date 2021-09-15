package webenv

import (
	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/services"
	"log"
	"net/http"
	"sync"

	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/models"
	"github.com/devgek/webskeleton/msg"
	"github.com/devgek/webskeleton/web/template"
	"github.com/gobuffalo/packr/v2"

	_ "github.com/jinzhu/gorm/dialects/postgres" // gorm for postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // gorm for sqlite3
)

//Env the environment
type Env struct {
	Api            bool
	TStore         template.TStore
	Templates      *packr.Box
	Assets         http.FileSystem
	DS             data.Datastore
	Services       *services.Services
	MessageLocator *msg.MessageLocator
	EF             models.EntityFactory
}

var once sync.Once

//theEnv singleton instance for the app env
var theEnv *Env

func GetEnv() *Env {
	return theEnv
}

//GetApiEnv return new initialized environment for serving api
func GetApiEnv(isTest bool) *Env {
	once.Do(func() {
		originalAssetBox := packr.New("assets", "../assets")
		// assetBox := packrfix.New(origninalAssetBox)

		//load locale specific message file, if not default
		// messages, err := assetBox.Find("msg/messages-en.yaml")
		messages, err := originalAssetBox.Find("msg/messages.yaml")

		//load messages
		ml := msg.NewMessageLocator(messages)

		//here we create the datastore
		//?_foreign_keys=1, neccessary for golang to respect foreign key constraints on sqlite3 db
		var ds data.Datastore
		if isTest {
			ds, err = data.NewInMemoryDatastore()
		} else {
			if config.DatastoreSystem() == "postgres" {
				ds, err = data.NewPostgres()
			} else {
				ds, err = data.NewSqlite(config.DatabaseName)
			}
		}
		if err != nil {
			log.Panic(err)
		}

		s := services.NewServices(models.EntityFactory{}, ds)

		theEnv = &Env{Api: true, TStore: nil, Templates: nil, Assets: originalAssetBox, DS: ds, Services: s, MessageLocator: ml, EF: models.EntityFactory{}}
	})

	return GetEnv()
}

//GetWebEnv return new initialized environment
func GetWebEnv() *Env {
	once.Do(func() {
		// ../web/templates important for packr2 to find files
		originalTemplateBox := packr.New("templates", "../templates")
		// templateBox := packrfix.New(origninalTemplateBox)

		//init TStore
		tStore := template.NewBoxBasedTemplateStore(originalTemplateBox)

		originalAssetBox := packr.New("assets", "../assets")
		// assetBox := packrfix.New(origninalAssetBox)

		//load locale specific message file, if not default
		// messages, err := assetBox.Find("msg/messages-en.yaml")
		messages, err := originalAssetBox.Find("msg/messages.yaml")

		//load messages
		ml := msg.NewMessageLocator(messages)

		//here we create the datastore
		//?_foreign_keys=1, neccessary for golang to respect foreign key constraints on sqlite3 db
		var ds data.Datastore
		if config.DatastoreSystem() == "postgres" {
			ds, err = data.NewPostgres()
		} else {
			ds, err = data.NewSqlite(config.DatabaseName)
		}
		if err != nil {
			log.Panic(err)
		}

		s := services.NewServices(models.EntityFactory{}, ds)

		theEnv = &Env{Api: false, TStore: tStore, Templates: originalTemplateBox, Assets: originalAssetBox, DS: ds, Services: s, MessageLocator: ml, EF: models.EntityFactory{}}
	})

	return GetEnv()
}
