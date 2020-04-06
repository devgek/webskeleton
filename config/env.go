package config

import (
	"github.com/devgek/webskeleton/global"
	"github.com/devgek/webskeleton/models"
	"github.com/devgek/webskeleton/msg"
	"github.com/devgek/webskeleton/packrfix"
	"github.com/devgek/webskeleton/web/template"
	"github.com/gobuffalo/packr/v2"
	"log"
	"net/http"
	"sync"

	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/services"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm for sqlite3
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
		templateBox := packrfix.New(origninalTemplateBox)

		//init TStore
		tStore := template.NewBoxBasedTemplateStore(templateBox)

		origninalAssetBox := packr.New("assets", "../web/assets")
		assetBox := packrfix.New(origninalAssetBox)

		//load locale specific message file, if not default
		// messages, err := assetBox.Find("msg/messages-en.yaml")
		messages, err := assetBox.Find("msg/messages.yaml")

		//load messages
		ml := msg.NewMessageLocator(messages)

		//here we create the datastore
		ds, err := data.NewDatastore("sqlite3", global.DatabaseName)
		if err != nil {
			log.Panic(err)
		}

		services := services.NewServices(ds)

		webEnv = &Env{TStore: tStore, Templates: origninalTemplateBox, Assets: assetBox, DS: ds, Services: services, MessageLocator: ml}
	})

	return webEnv
}
