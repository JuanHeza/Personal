package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/Home", handler).Methods("GET")
	return r
}
*/
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/Home", handler).Methods("GET")
	staticFileDirectory := http.Dir("./templates/")
	staticFileHandler := http.StripPrefix("/templates/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/templates/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/proyects", getProyectHandler).Methods("GET")
	r.HandleFunc("/proyects", createProyectHandler).Methods("POST")
	return r
}

func main() {
	r := newRouter()
	//http.HandleFunc("/",handler)
	http.ListenAndServe(":8080", r)
	// welcome()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}
