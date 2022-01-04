package env

import (
	"log"
	"sync"

	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/data"
	entitymodel "github.com/devgek/webskeleton/entity/model"
	generated_models "github.com/devgek/webskeleton/models/generated"
	"github.com/devgek/webskeleton/services"
	_ "github.com/jinzhu/gorm/dialects/postgres" // gorm for postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // gorm for sqlite3
)

//ApiEnv the environment
type ApiEnv struct {
	DS       data.Datastore
	EF       entitymodel.EntityFactory
	Services *services.Services
}

var once sync.Once

//theEnv singleton instance for the app env
var theApiEnv *ApiEnv

func GetEnv() *ApiEnv {
	return theApiEnv
}

//GetApiEnv return new initialized environment for serving api
func GetApiEnv(isTest bool) *ApiEnv {
	once.Do(func() {
		//here we create the datastore
		//?_foreign_keys=1, neccessary for golang to respect foreign key constraints on sqlite3 db
		var ds data.Datastore
		var err error
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

		ef := generated_models.EntityFactoryImpl{}
		s := services.NewServices(ef, ds)

		theApiEnv = &ApiEnv{DS: ds, Services: s, EF: &ef}
	})

	return GetEnv()
}
