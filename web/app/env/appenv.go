package env

import (
	"log"
	"net/http"
	"sync"

	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/data"
	entitymodel "github.com/devgek/webskeleton/entity/model"
	generated_models "github.com/devgek/webskeleton/models/generated"
	"github.com/devgek/webskeleton/services"
	"github.com/devgek/webskeleton/web/app/msg"
	"github.com/devgek/webskeleton/web/app/template"
	"github.com/gobuffalo/packr/v2"
)

//AppEnv the environment
type AppEnv struct {
	DS             data.Datastore
	EF             entitymodel.EntityFactory
	Services       *services.Services
	TStore         template.TStore
	Templates      *packr.Box
	Assets         http.FileSystem
	MessageLocator *msg.MessageLocator
}

var once sync.Once

//theEnv singleton instance for the app env
var theAppEnv *AppEnv

func GetAppEnv() *AppEnv {
	return theAppEnv
}

//GetWebEnv return new initialized environment
func GetWebEnv() *AppEnv {
	once.Do(func() {
		// ../templates important for packr2 to find files
		originalTemplateBox := packr.New("templates", "../templates")
		// templateBox := packrfix.New(origninalTemplateBox)

		//init TStore
		tStore := template.NewBoxBasedTemplateStore(originalTemplateBox)

		// ../assets important for packr2 to find files
		originalAssetBox := packr.New("assets", "../assets")
		// assetBox := packrfix.New(origninalAssetBox)

		//load locale specific message file, if not default
		// messages, err := assetBox.Find("msg/messages-en.yaml")
		messages, err := originalAssetBox.Find("msg/messages.yaml")
		if err != nil {
			log.Panic(err)
		}

		//load messages
		ml := msg.NewMessageLocator(messages)

		//here we create the datastore
		//?_foreign_keys=1, necessary for golang to respect foreign key constraints on sqlite3 db
		var ds data.Datastore
		if config.DatastoreSystem() == "postgres" {
			ds, err = data.NewPostgres()
		} else {
			ds, err = data.NewSqlite(config.DatabaseName)
		}
		if err != nil {
			log.Panic(err)
		}

		ef := generated_models.EntityFactoryImpl{}
		s := services.NewServices(ef, ds)

		theAppEnv = &AppEnv{
			DS:             ds,
			EF:             &ef,
			Services:       s,
			TStore:         tStore,
			Templates:      originalTemplateBox,
			Assets:         originalAssetBox,
			MessageLocator: ml,
		}
	})

	return theAppEnv
}
