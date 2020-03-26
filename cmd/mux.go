package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/webmux"
	"github.com/spf13/cobra"
)

// muxcmd represents the serve command
var muxCmd = &cobra.Command{
	Use:   "mux",
	Short: "start web server and serve html with mux router",
	Long:  `webskeleton serve; a typical go web app`,
	Run: func(cmd *cobra.Command, args []string) {
		runMux(cmd)
	},
}

func init() {
	rootCmd.AddCommand(muxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	muxCmd.Flags().String("port", "8080", "The port this app listens")
	muxCmd.Flags().Bool("debug", false, "debug mode on/off")
}

func runMux(cmd *cobra.Command) {
	// start the web server
	port, _ := cmd.Flags().GetString("port")
	config.Debug, _ = cmd.Flags().GetBool("debug")
	log.Println("Starting webskeleton with mux on port ", port)
	if config.Debug {
		log.Println("Debug mode is on")
	}

	//the app env, containes pointer to db and services
	env := config.GetWebEnv()
	router := webmux.InitWeb(env)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
