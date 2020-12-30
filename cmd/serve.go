package cmd

import (
	"log"

	"github.com/devgek/webskeleton/web/echo"
	webenv "github.com/devgek/webskeleton/web/env"

	"github.com/devgek/webskeleton/config"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve --config=configfile.yaml",
	Short: "start web server and serve html with echo http server",
	Long:  `bla bla`,
	Run: func(cmd *cobra.Command, args []string) {
		runServe(cmd)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().String("port", "8080", "The port this app listens")
	// serveCmd.Flags().Bool("debug", false, "debug mode on/off")
}

func runServe(cmd *cobra.Command) {
	// start the web server
	log.Println("Starting ", config.ProjectName, " server on port ", config.ServerPort())
	if config.IsDev() {
		log.Println("Development mode is on")
	}
	if config.IsAssetsCache() {
		log.Println("Assets cache mode is on")
	}
	if config.IsServerDebug() {
		log.Println("Server debug mode is on")
	}
	if config.IsDatastoreLog() {
		log.Println("Datastore log mode is on")
	}

	env := webenv.GetWebEnv()

	e := echo.InitEcho(env)

	if config.IsServerSecure() {
		log.Fatal(e.StartTLS(":"+config.ServerPort(), config.ServerCert(), config.ServerKey()))
	} else {
		log.Fatal(e.Start(":" + config.ServerPort()))
	}
}
