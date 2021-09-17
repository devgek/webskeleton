package cmd

import (
	"github.com/devgek/webskeleton/web/api"
	"github.com/devgek/webskeleton/web/api/env"
	"log"

	"github.com/devgek/webskeleton/config"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var apiCmd = &cobra.Command{
	Use:   "api --config=configfile.yaml",
	Short: "start web server and serve json API with echo http server",
	Long:  `bla bla`,
	Run: func(cmd *cobra.Command, args []string) {
		runApi(cmd)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().String("port", "8080", "The port this app listens")
	// serveCmd.Flags().Bool("debug", false, "debug mode on/off")
}

func runApi(cmd *cobra.Command) {
	// start the web server
	log.Println("Starting ", config.ProjectName, " API on port ", config.ServerPort())
	if config.IsServerSecure() {
		log.Println("Secure (TLS) mode is on")
	}
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

	apiEnv := env.GetApiEnv(false)

	e := api.InitEchoApi(apiEnv)

	if config.IsServerSecure() {
		log.Fatal(e.StartTLS(":"+config.ServerPort(), config.ServerCert(), config.ServerKey()))
	} else {
		log.Fatal(e.Start(":" + config.ServerPort()))
	}
}
