package cmd

import (
	"log"

	"kahrersoftware.at/webskeleton/global"
	"kahrersoftware.at/webskeleton/web/echo"

	"github.com/spf13/cobra"
	"kahrersoftware.at/webskeleton/config"
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
	log.Println("Starting ", global.ProjectName, " server on port ", global.ServerPort())
	if global.IsDev() {
		log.Println("Development mode is on")
	}
	if global.IsServerDebug() {
		log.Println("Debug mode is on")
	}

	env := config.GetWebEnv()

	e := echo.InitEcho(env)

	if global.IsServerSecure() {
		log.Fatal(e.StartTLS(":"+global.ServerPort(), global.ServerCert(), global.ServerKey()))
	} else {
		log.Fatal(e.Start(":" + global.ServerPort()))
	}
}
