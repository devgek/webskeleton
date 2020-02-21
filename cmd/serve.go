package cmd

import (
	"log"
	"net/http"
	"time"

	"kahrersoftware.at/webskeleton/auth"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"kahrersoftware.at/webskeleton/config"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start web server and serve html",
	Long:  `webskeleton serve; a typical go web app`,
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
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().String("port", "8080", "The port this app listens")
}

func runServe(cmd *cobra.Command) {
	env := config.InitEnv()
	r := mux.NewRouter()

	r.Handle("/", env.NewTemplateHandler("login.html"))
	r.Handle("/loginuser", auth.HandleLoginUser(env))
	//MustAuth secures the following site to be authenticated (auth cookie web-skeleton)
	r.Handle("/page1", auth.MustAuth(env.NewTemplateHandler("page1.html")))
	r.Handle("/page2", env.NewTemplateHandler("page2.html"))
	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "webskeleton-auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	// r.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	// start the web server
	port, _ := cmd.Flags().GetString("port")
	log.Println("Starting webskeleton on port ", port)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
