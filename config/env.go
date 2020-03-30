package config

import (
	"github.com/devgek/webskeleton/global"
	"github.com/devgek/webskeleton/helper"
	"github.com/devgek/webskeleton/msg"
	"github.com/devgek/webskeleton/packrfix"
	"github.com/devgek/webskeleton/web/template"
	"github.com/gobuffalo/packr/v2"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
}

var once sync.Once

//WebEnv singleton instance for the app env
var webEnv *Env

//GetWebEnv return new initialized environment
func GetWebEnv() *Env {
	once.Do(func() {
		root, err := os.Getwd()
		helper.PanicOnError(err)
		//init template FileSystem
		templatePath := filepath.Join(root, "web", "templates")
		origninalTemplateBox := packr.New("templates", templatePath)
		templateBox := packrfix.New(origninalTemplateBox)

		//init TStore
		tStore := template.NewBoxBasedTemplateStore(templateBox)

		//init asset FileSystem
		path := filepath.Join(root, "web", "assets")
		origninalAssetBox := packr.New("assets", path)
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
