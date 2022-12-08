package env

import (
	"embed"
	_ "embed"
	"log"
	"net/http"
	"sync"

	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/data"
	entitymodel "github.com/devgek/webskeleton/entity/model"
	genmodels "github.com/devgek/webskeleton/models/generated"
	"github.com/devgek/webskeleton/services"
	"github.com/devgek/webskeleton/web/app/msg"
	"github.com/devgek/webskeleton/web/app/template"
	"github.com/gobuffalo/packr/v2"
)

// AppEnv the environment
type AppEnv struct {
	DS             data.Datastore
	EF             entitymodel.EntityFactory
	Services       *services.Services
	TStore         template.TStore
	Templates      *packr.Box
	Assets         http.FileSystem
	MessageLocator *msg.MessageLocator
}

// load locale specific message file messages.yaml or messages-en.yaml
//
//go:embed msg/messages.yaml
var messages []byte

//go:embed assets/*
var assets embed.FS

var once sync.Once

// theEnv singleton instance for the app env
var theAppEnv *AppEnv

func GetAppEnv() *AppEnv {
	return theAppEnv
}

// GetWebEnv return new initialized environment
func GetWebEnv() *AppEnv {
	once.Do(func() {
		// ../templates important for packr2 to find files
		originalTemplateBox := packr.New("templates", "../templates")
		// templateBox := packrfix.New(origninalTemplateBox)

		//init TStore
		tStore := template.NewBoxBasedTemplateStore(originalTemplateBox)

		//load messages
		ml := msg.NewMessageLocator(messages)

		//here we create the datastore
		//?_foreign_keys=1, necessary for golang to respect foreign key constraints on sqlite3 db
		var ds data.Datastore
		var err error
		if config.DatastoreSystem() == "postgres" {
			ds, err = data.NewPostgres()
		} else {
			ds, err = data.NewSqlite(config.DatabaseName)
		}
		if err != nil {
			log.Panic(err)
		}

		ef := genmodels.EntityFactoryImpl{}
		s := services.NewServices(ef, ds)

		theAppEnv = &AppEnv{
			DS:             ds,
			EF:             &ef,
			Services:       s,
			TStore:         tStore,
			Templates:      originalTemplateBox,
			Assets:         http.FS(assets),
			MessageLocator: ml,
		}
	})

	return theAppEnv
}
