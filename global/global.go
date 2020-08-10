package global

import "github.com/spf13/viper"

// ...
var (
	ProjectName    = "webskeleton" //do not change
	ProjectTitle   = "go-webskeleton"
	ProjectVersion = "V1.0"
	DatabaseName   = "webskeleton.db"
	Debug          = false
	StartPage      = "/"
)

//IsDev ...
func IsDev() bool {
	return viper.GetString("mode") == "DEV"
}

//IsServerDebug ...
func IsServerDebug() bool {
	return viper.GetBool("server.debug")
}

//ServerPort ...
func ServerPort() string {
	return viper.GetString("server.port")
}

//IsServerSecure ...
func IsServerSecure() bool {
	return viper.GetBool("server.tls")
}

//ServerCert ...
func ServerCert() string {
	return viper.GetString("server.cert")
}

//ServerKey ...
func ServerKey() string {
	return viper.GetString("server.key")
}

//DatastoreSystem ...
func DatastoreSystem() string {
	return viper.GetString("datastore.system")
}

//DatastoreHost ...
func DatastoreHost() string {
	return viper.GetString("datastore.host")
}

//DatastorePort ...
func DatastorePort() string {
	return viper.GetString("datastore.port")
}

//DatastoreUser ...
func DatastoreUser() string {
	return viper.GetString("datastore.user")
}

//DatastorePassword ...
func DatastorePassword() string {
	return viper.GetString("datastore.password")
}

//DatastoreDb ...
func DatastoreDb() string {
	return viper.GetString("datastore.db")
}
