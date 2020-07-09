package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "JHZ697heza"
	dbname   = "proyect_encyclopedia"
	//APIKey is the WakaTime key
	APIKey = "502f3c9e-67d4-48ce-a6b9-77dbe3887e7c"
)

var files = []string{"templates/base.html", "templates/footer.html", "templates/header.html", "templates/welcome-template.html", "templates/error.html", "templates/proyect.html", "templates/card.html", "templates/crud.html", "templates/login.html"}

//ProyectFiles is the complete info about certain proyect or a general vieo of every project
var ProyectFiles = map[string]string{
	"General":  "static/Database/Projects.json",
	"Elements": "static/Database/ElementsData.json",
}

//Templates is where the templates in "files "are parsed
var Templates *template.Template

func newRouter() *mux.Router {
	r := mux.NewRouter()
	//		NOTA toma el strip prefix y lo pasa al fileserver, si el html dice 'miCasa' en stripprefix, y el codigo dice 'monclova' en fileserver pues lo manda a monclova
	staticFileDirectory := http.Dir("./templates")
	staticFileHandler := http.StripPrefix("/templates/", http.FileServer(staticFileDirectory))
	staticFileSheet := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	r.PathPrefix("/templates/").Handler(staticFileHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(staticFileSheet).Methods("GET")
	TemplateFunctions := template.FuncMap{
		"Session": IsAuthenticated,
		"Admin":   IsAdmin,
		"Links":   Links,
		"Join": Join,
	}
	Templates = template.New("")
	Templates = template.Must(Templates.Funcs(TemplateFunctions).ParseFiles(files...))

	r.HandleFunc("/", welcomeHandler).Methods("GET")
	r.HandleFunc("/Home", homeHandler).Methods("GET")
	r.HandleFunc("/LogIn", logInHandler).Methods("GET")
	r.HandleFunc("/LogIn", sessionHandler).Methods("POST")
	r.HandleFunc("/LogOut", logOut)
	r.HandleFunc("/Proyect/{name}", proyectHandler)
	r.HandleFunc("/Error", errorHandler)
	r.HandleFunc("/Edit", handler)
	r.HandleFunc("/Edit/{id}", updateProject)
	r.HandleFunc("/Data", getProyectHandler).Methods("GET")
	r.HandleFunc("/Data", createProject).Methods("POST")
	r.HandleFunc("/Data/{id}", getOneProjectHandler).Methods("GET")
	r.HandleFunc("/Data/{id}", updateProjectHandler).Methods("PUT")
	r.HandleFunc("/Data/{id}", deleteProjectHandler).Methods("DELETE")

	r.HandleFunc("/Crud/Modelo/{proyecto}", createModelHandler).Methods("POST")
	r.HandleFunc("/Crud/Modelo", updateModelHandler).Methods("PUT")
	r.HandleFunc("/Crud/Modelo/{id}", deleteModelHandler).Methods("DELETE")

	r.HandleFunc("/Crud/Funcion/{proyecto}", createFunctionHandler).Methods("POST")
	r.HandleFunc("/Crud/Funcion", updateFunctionHandler).Methods("PUT")
	r.HandleFunc("/Crud/Funcion/{id}", deleteFunctionHandler).Methods("DELETE")

	r.HandleFunc("/Crud/Notas/{proyecto}", createNotasHandler).Methods("POST")
	r.HandleFunc("/Crud/Notas", updateNotasHandler).Methods("PUT")
	r.HandleFunc("/Crud/Notas/{id}", deleteNotasHandler).Methods("DELETE")

	r.HandleFunc("/Crud/Tarea/{proyecto}", createTareasHandler).Methods("POST")
	r.HandleFunc("/Crud/Tarea", updateTareasHandler).Methods("PUT")
	r.HandleFunc("/Crud/Tarea/{id}", deleteTareasHandler).Methods("DELETE")

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
	r := newRouter()
	log.Println("listening http://127.0.0.1:8080/Home")
	http.ListenAndServe(":8080", r)
	//welcome()
}

// IfErr is to check the error and redirect to error page if necessary
func IfErr(err error, w http.ResponseWriter, r *http.Request) {
	fmt.Println("IfErr", err)
	if err != nil {
		http.Redirect(w, r, "/Error/", http.StatusFound)
	}
}
