package main

import (
	// "fmt"

	"net/http"

	"github.com/JuanHeza/Personal/controllers"
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	//		NOTA toma el strip prefix y lo pasa al fileserver, si el html dice 'miCasa' en stripprefix, y el codigo dice 'monclova' en fileserver pues lo manda a monclova
	staticFileDirectory := http.Dir("./templates")
	staticFileHandler := http.StripPrefix("/Views/", http.FileServer(staticFileDirectory))
	staticFileSheet := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	r.PathPrefix("/Views/").Handler(staticFileHandler).Methods(http.MethodGet)
	r.PathPrefix("/static/").Handler(staticFileSheet).Methods(http.MethodGet)
	controllers.InitTemplates()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/Home/", http.StatusFound) }) // WORKING
	r.HandleFunc("/Home/", controllers.HomeHandler).Methods(http.MethodGet)                                             // WORKING
	r.HandleFunc("/Error/", controllers.ErrorHandler)                                                                   // WORKING                              // WORKING

	r.HandleFunc("/LogIn", controllers.LogInHandler).Methods(http.MethodGet)    // WORKING
	r.HandleFunc("/LogIn", controllers.SessionHandler).Methods(http.MethodPost) //
	r.HandleFunc("/LogOut", controllers.LogOutHandler)                          //
	r.HandleFunc("/Project/{id}", controllers.ProjectHandler)                   // WORKING
	r.HandleFunc("/ProjectList/", controllers.ProjectListHandler)               // WORKING
	r.HandleFunc("/ProjectList/{leng}", controllers.ProjectListHandler)         // WORKING
	r.HandleFunc("/PostList/", controllers.PostListHandler)                     // WORKING
	r.HandleFunc("/Post/{id}", controllers.PostHandler)                         // WORKING
	r.HandleFunc("/Edit/", controllers.EditHandler)                             // WORKING
	r.HandleFunc("/Edit/Post/", controllers.PostEditHandler)                    // WORKING
	r.HandleFunc("/Edit/Project/", controllers.ProjectEditHandler)              // WORKING
	r.HandleFunc("/Edit/Post/{id}", controllers.PostEditHandler)                // WORKING
	r.HandleFunc("/Edit/Project/{id}", controllers.ProjectEditHandler)          // WORKING
	r.HandleFunc("/Create/{model}", controllers.CreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/Update/{model}/{id}", controllers.UpdateHandler).Methods(http.MethodPut)
	r.HandleFunc("/Delete/{model}/{id}", controllers.DeleteHandler).Methods(http.MethodDelete)

	r.HandleFunc("/api/Home/", controllers.ApiHomeHandler).Methods("GET", "OPTIONS")               // WORKING
	r.HandleFunc("/api/Image/{folder}/{title}/", controllers.IconHandler).Methods("GET", "OPTIONS") // WORKING
	r.HandleFunc("/api/Project/{id}/", controllers.ApiProjectHandler).Methods("GET", "OPTIONS")    // WORKING
	r.HandleFunc("/api/Post/{id}/", controllers.ApiPostHandler).Methods("GET", "OPTIONS")          // WORKING
	r.HandleFunc("/api/Lista/{filtro}/", controllers.ApiListHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/Lenguage/{id}/", controllers.ApiListHandler).Methods("GET", "OPTIONS")

	// r.HandleFunc("/Crud/Links/{id}", controllers.UpdateLinkHandler).Methods(http.MethodPost)   //
	// r.HandleFunc("/Crud/Links/{id}", controllers.DeleteLinkHandler).Methods(http.MethodDelete) //
	// r.HandleFunc("/Crud/Notas", updateNotasHandler).Methods(http.MethodPut) //
	// r.HandleFunc("/Crud/Notas/{id}", deleteNotasHandler).Methods(http.MethodDelete) //
	// r.HandleFunc("/Data/{id}", updateProjectHandler).Methods(http.MethodPut) //
	// r.HandleFunc("/Data/{id}", deleteProjectHandler).Methods(http.MethodDelete) //

	return r
}
