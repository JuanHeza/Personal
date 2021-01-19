package main

import (
	// "fmt"
	// "net/http"
	"os"
	// "github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

//Usuario valores de la sesion
type Usuario struct {
	ID       string
	Password string
	Active   bool
	Admin    bool
}

//User es la variable que almacena la variable de la sesion
var (
	UserSession  *sessions.Session
	key          = []byte(os.Getenv("COOCKIE_KEY"))
	StoreSession = sessions.NewCookieStore(key)
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	if cr := IsAdminAutenticathed(); cr.Admin && cr.Auth {
// 		errorPage := Templates.Lookup("CRUD")
// 		if err := errorPage.ExecuteTemplate(w, "CRUD", nil); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 	} else {
// 		http.Redirect(w, r, "/Error", http.StatusFound)
// 	}
// }

// func updateProject(w http.ResponseWriter, r *http.Request) {
// 	if cr := IsAdminAutenticathed(); cr.Admin && cr.Auth {
// 		vars := mux.Vars(r)
// 		id := vars["id"]
// 		// fmt.Fprintf(w, Hi)
// 		pr, err := store.GetProyect(id)
// 		IfErr(err, w, r)
// 		errorPage := Templates.Lookup("CRUD")
// 		if err := errorPage.ExecuteTemplate(w, "CRUD", pr[0]); err != nil {
// 			fmt.Println(len(pr))
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 	} else {
// 		http.Redirect(w, r, "/Error", http.StatusFound)
// 	}
// }

// func errorHandler(w http.ResponseWriter, r *http.Request) {
// 	errorPage := Templates.Lookup("error")
// 	if err := errorPage.ExecuteTemplate(w, "error", nil); err != nil {
// 		fmt.Println("Error/", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func logInHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("init login")
// 	login := Templates.Lookup("login")
// 	err := login.ExecuteTemplate(w, "login", nil)
// 	IfErr(err, w, r)
// 	fmt.Println("close login")
// }
