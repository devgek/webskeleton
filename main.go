package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
	"text/template"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

type templateData struct {
	Host   string
	Client string
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		if t.filename == "login.html" {
			t.templ = template.Must(template.ParseFiles("./templates/" + t.filename))
		} else {
			t.templ = template.Must(template.ParseFiles("./templates/layout.html", "./templates/menu.html", "./templates/"+t.filename))
		}
	})

	cData, _ := FromContext(r.Context())

	data := map[string]interface{}{
		"Host":        r.Host,
		"VersionInfo": "V1.0",
		"UserID":      cData.UserID(),
	}

	t.templ.Execute(w, data)

}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags

	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/loginuser", loginHandler)
	http.Handle("/page1", MustAuth(&templateHandler{filename: "page1.html"}))
	http.Handle("/page2", &templateHandler{filename: "page2.html"})
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
