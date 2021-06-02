package webenv

import (
	"log"
	"net/http"
	"sync"

	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/models"
	"github.com/devgek/webskeleton/msg"
	"github.com/devgek/webskeleton/web/template"
	"github.com/gobuffalo/packr/v2"

	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/services"
	_ "github.com/jinzhu/gorm/dialects/postgres" // gorm for postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // gorm for sqlite3
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

		services := services.NewServices(ds)

		webEnv = &Env{TStore: tStore, Templates: originalTemplateBox, Assets: originalAssetBox, DS: ds, Services: services, MessageLocator: ml}
	})

	return webEnv
}
