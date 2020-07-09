package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

//Usuario valores de la sesion
type Usuario struct {
	Id       string
	Password string
	Active   bool
	Admin    bool
}

//User es la variable que almacena la variable de la sesion
var (
	UserSession  *sessions.Session
	key          = []byte("key")
	StoreSession = sessions.NewCookieStore(key)
)

func handler(w http.ResponseWriter, r *http.Request) {
	if IsAuthenticated() && IsAdmin() {
		errorPage := Templates.Lookup("CRUD")
		if err := errorPage.ExecuteTemplate(w, "CRUD", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Redirect(w, r, "/Error", http.StatusFound)
	}
}

func updateProject(w http.ResponseWriter, r *http.Request) {
	if IsAuthenticated() && IsAdmin() {
	vars := mux.Vars(r)
	id := vars["id"]
	// fmt.Fprintf(w, Hi)
	pr, err := store.GetProyect(id)
	IfErr(err, w, r)
	errorPage := Templates.Lookup("CRUD")
	if err := errorPage.ExecuteTemplate(w, "CRUD", pr[0]); err != nil {
		fmt.Println(len(pr))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
		} else {
			http.Redirect(w, r, "/Error", http.StatusFound)
		}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	errorPage := Templates.Lookup("error")
	if err := errorPage.ExecuteTemplate(w, "error", nil); err != nil {
		fmt.Println("Error/", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func logInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("init login")
	login := Templates.Lookup("login")
	err := login.ExecuteTemplate(w, "login", nil)
	IfErr(err, w, r)
	fmt.Println("close login")
}

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("init sesion")

	user := Usuario{}
	r.ParseForm()
	user.Id = r.Form.Get("usuario")
	user.Password = r.Form.Get("contraseña")
	err := Store.LogIn(store, &user)
	if err != nil {
		login := Templates.Lookup("login")
		err := login.ExecuteTemplate(w, "login", "hi")
		IfErr(err, w, r)
	} else {
		// User.Active = true
		// User.Admin = true
		UserSession, _ = StoreSession.Get(r, user.Id)
		UserSession.Values["authenticated"] = true
		UserSession.Values["Admin"] = user.Admin
		fmt.Println(UserSession.Values)
		UserSession.Save(r, w)
		http.Redirect(w, r, "/Home", http.StatusFound)
	}
	fmt.Println("close sesion")
}

// Links get the social media links
func Links() []int {
	return []int{0, 1, 2}
}

//IsAuthenticated verify if the user is loged in
func IsAuthenticated() bool {
	//session, _ := StoreSession.Get(r, User.Id)
	if UserSession == nil {
		return false
	}
	if auth, ok := UserSession.Values["authenticated"].(bool); !ok || !auth {
		return false
	}
	return true
}

//IsAdmin verify if the user is loged in
func IsAdmin() bool {
	//session, _ := StoreSession.Get(r, User.Id)
	if UserSession == nil {
		return false
	}
	if auth, ok := UserSession.Values["Admin"].(bool); !ok || !auth {
		return false
	}
	return true
}

func logOut(w http.ResponseWriter, r *http.Request) {
	// session, _ := StoreSession.Get(r, User.Id)

	// Revoke users authentication
	UserSession.Values["authenticated"] = false
	UserSession.Save(r, w)
	http.Redirect(w, r, "/Home", http.StatusFound)
}
