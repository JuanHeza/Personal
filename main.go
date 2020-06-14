package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

/*
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/Home", handler).Methods("GET")
	return r
}
*/
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "JHZ697heza"
	dbname   = "proyect_encyclopedia"
)

var files = []string{"templates/base.html", "templates/footer.html", "templates/header.html", "templates/welcome-template.html", "templates/error.html", "templates/proyect.html", "templates/card.html"}

//ProyectFiles is the complete info about certain proyect or a general vieo of every project
var ProyectFiles = map[string]string{
	"General":  "static/Database/Projects.json",
	"Elements": "static/Database/ElementsData.json",
}

//Templates is where the templates in "files "are parsed
var Templates *template.Template

func newRouter() *mux.Router {
	r := mux.NewRouter()
	staticFileDirectory := http.Dir("./templates/")
	staticFileHandler := http.StripPrefix("/templates/", http.FileServer(staticFileDirectory))
	staticFileSheet := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/templates/").Handler(staticFileHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(staticFileSheet).Methods("GET")

	Templates = template.Must(template.ParseFiles(files...))

	r.HandleFunc("/", welcomeHandler).Methods("GET")
	// r.HandleFunc("/proyect/{name}", handler)
	r.HandleFunc("/Proyect/{name}", proyectHandler)
	r.HandleFunc("/Home", homeHandler).Methods("GET")
	r.HandleFunc("/Error", errorHandler)
	r.HandleFunc("/Proyects", getProyectHandler).Methods("GET")
	r.HandleFunc("/Proyects", createProyectHandler).Methods("POST")
	return r
}

func main() {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	InitStore(&dbStore{db: db})
	ReadJSON()
	r := newRouter()
	//http.HandleFunc("/",handler)
	fmt.Println("listening http://127.0.0.1:8080/Home")
	http.ListenAndServe(":8080", r)
	// welcome()
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Hi := fmt.Sprintf("Hi %v", vars["name"])
	fmt.Fprintf(w, Hi)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	errorPage := Templates.Lookup("error")
	if err := errorPage.ExecuteTemplate(w, "error", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
