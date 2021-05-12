package controllers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JuanHeza/Personal/models"
)

var errorChannel = make(chan string, 1)

// Credentials struct defines if the user is logged in and if is admin
type Credentials struct {
	Admin bool
	Auth  bool
}

// IfErr is to check the error and redirect to error page if necessary True if is free of errors and false & redirect if an error exist
func IfErr(err error, w http.ResponseWriter, r *http.Request) bool {
	if err != nil {
		errorChannel <- err.Error()
		fmt.Println("IfErr:", err.Error())
		http.Redirect(w, r, "/Error/", http.StatusFound)
		return false
	}
	return true
}

//tryTemplate execute the given template in an auxiliar buffer, if there is an error is sended to the Error route
func tryTemplate(template string, data map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	var auxBuffer = &bytes.Buffer{}
	if IfErr(Templates.Lookup(template).ExecuteTemplate(auxBuffer, template, data), w, r) {
		auxBuffer.WriteTo(w)
	}
}

// IsAdminAutenticathed verify the credentials of the user is logged in
func IsAdminAutenticathed() (cr Credentials) {
	var ok bool
	if UserSession == nil {
		return Credentials{false, false}
	}
	if cr.Auth, ok = UserSession.Values["authenticated"].(bool); !ok {
		cr.Auth = false
	}
	if cr.Admin, ok = UserSession.Values["Admin"].(bool); !ok {
		cr.Admin = false
	}
	return
}

//Fecha returns the date in format DD/MM/YYYY
func Fecha(t time.Time, input ...bool) string {
	y, m, d := t.Date()
	if len(input) == 0 {
		return fmt.Sprintf("%02d/%02d/%4d", d, m, y)
	}
	return fmt.Sprintf("%4d-%02d-%02d", y, m, d)
}

// Links get the social media links
func Links() []*models.LinkModel {
	if models.StaticDataCollection["link"] == nil {
		return []*models.LinkModel{{Link: "mailto:juanehza@hotmail.com", Icon: "envelope.png", ID: 1}, {Link: "https://repl.it/@JuanHeza/", Icon: "Repl.it.png", ID: 2}, {Link: "https://github.com/JuanHeza", Icon: "GitHub-Mark-32px.png", ID: 3}}
	}
	return models.StaticDataCollection["link"].([]*models.LinkModel)
}

func folderExist(id string) string {
	folder := fmt.Sprint(id)
	_, err := os.Stat("./static/images/" + folder)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("./static/images/"+folder+"/", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
	return "./static/images/" + folder + "/"
}

//Horas is a  function return the hours of the given time, counting just days not monts
func Horas(t time.Time) int {
	d := t.Day()
	if m := t.Month(); m == time.February {
		d += 31
	}
	h := t.Hour()
	return h + (d * 24)
}
