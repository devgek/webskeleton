package cmd

import (
	"github.com/devgek/webskeleton/global"
	"github.com/devgek/webskeleton/web/echo"
	"log"

	"github.com/devgek/webskeleton/config"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start web server and serve html with echo http server",
	Long:  `bla bla`,
	Run: func(cmd *cobra.Command, args []string) {
		runEcho(cmd)
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
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().String("port", "8080", "The port this app listens")
	serveCmd.Flags().Bool("debug", false, "debug mode on/off")
}

func runEcho(cmd *cobra.Command) {
	// start the web server
	port, _ := cmd.Flags().GetString("port")
	global.Debug, _ = cmd.Flags().GetBool("debug")
	log.Println("Starting webskeleton server with echo on port ", port)
	if global.Debug {
		log.Println("Debug mode is on")
	}

	env := config.GetWebEnv()

	e := echo.InitEcho(env)

	log.Fatal(e.Start(":" + port))
}
