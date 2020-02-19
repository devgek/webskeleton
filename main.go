package main

import (
	"flag"
	"log"
	"net/http"

	"kahrersoftware.at/webskeleton/config"
)

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags

	env := config.InitEnv()

	http.Handle("/", env.NewTemplateHandler("login.html"))
	http.Handle("/loginuser", handleLoginUser(env))
	//MustAuth secures the following site to be authenticated (auth cookie web-skeleton)
	http.Handle("/page1", MustAuth(env.NewTemplateHandler("page1.html")))
	http.Handle("/page2", env.NewTemplateHandler("page2.html"))
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "webskeleton-auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	// start the web server
	log.Println("Starting webskeleton on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
