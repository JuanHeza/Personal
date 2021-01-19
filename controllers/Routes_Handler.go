package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JuanHeza/Personal/models"
	"github.com/gorilla/sessions"
)

var (
	//UserSession store the data for the cookie
	UserSession *sessions.Session
	key         = []byte(os.Getenv("COOCKIE_KEY"))
	//StoreSession is a cookie
	StoreSession = sessions.NewCookieStore(key)
)

//EditHandler is the handler for the configuration menu
func EditHandler(w http.ResponseWriter, r *http.Request) {
	// var pr data
	var err error
	var Data = map[string]interface{}{
		"JS":  []string{"https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js", "http://malsup.github.com/jquery.form.js", "https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.3/js/select2.min.js", "/static/scripts/edit.js"},
		"CSS": []string{"https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.3/css/select2.css", "/static/stylesheets/dragNdrop.css", "/static/stylesheets/forms.css"},
	}
	// if cr := IsAdminAutenticathed(); cr.Admin && cr.Auth {
	// vars := mux.Vars(r)
	// id, err := strconv.Atoi(vars["id"])
	log.Printf("/Edit/ @ Project.EditHandler")
	var aux = make(map[string]interface{})
	_, aux["Projects"], err = models.Project.ReadProject()
	_, aux["Posts"], err = models.Post.ReadPost()
	Data["Data"] = aux
	if IfErr(err, w, r) {
		tryTemplate("Edit", Data, w, r)
	}
	// } else {
	// 	http.Redirect(w, r, "/Error/", http.StatusFound)
	// }
}

//SessionHandler IS #ERROR
func SessionHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/LogIn/ @ Routes_Handler.SessionHandler")
	r.ParseForm()
	var Data = map[string]interface{}{
		"CSS": []string{"/static/stylesheets/style.css"},
	}
	user := models.NewUser(r.Form.Get("usuario"), r.Form.Get("contraseña"))
	save := models.NewUser(os.Getenv("User"), os.Getenv("Pass"))

	if user.ID != save.ID || user.Password != save.Password {
		Data["Data"] = fmt.Errorf("no se pudo iniciar sesion")
		tryTemplate("login", Data, w, r)
	} else {
		// User.Active = true
		// User.Admin = true
		UserSession, _ = StoreSession.Get(r, user.ID)
		UserSession.Values["authenticated"] = true
		UserSession.Values["Admin"] = true
		UserSession.Values["leng"] = "es"
		fmt.Println(UserSession.Values)
		UserSession.Save(r, w)
		http.Redirect(w, r, "/Home/", http.StatusFound)
	}
}

//LogInHandler IS #ERROR
func LogInHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/LogIn/ @ Routes_Handler.LogInHandler")
	tryTemplate("login", nil, w, r)
}

//LogOutHandler IS #ERROR
func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/LogOut/ @ Routes_Handler.LogOutHandler")
	// session, _ := StoreSession.Get(r, User.Id)

	// Revoke users authentication
	UserSession.Values["authenticated"] = false
	UserSession.Save(r, w)
	http.Redirect(w, r, "/Home/", http.StatusFound)
}

//ErrorHandler IS #ERROR
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/Error/ @ Routes_Handler.ErrorHandler")
	var err = "Error Desconocido"
	if len(errorChannel) > 0 {
		err = <-errorChannel
	}
	errorPage := Templates.Lookup("error")
	var Data = map[string]interface{}{
		"Data": err,
		"CSS":  []string{},
		"JS":   []string{},
	}
	if err := errorPage.ExecuteTemplate(w, "error", Data); err != nil {
		fmt.Println("Error/", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
