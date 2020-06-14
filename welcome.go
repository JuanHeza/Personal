package main

//https://medium.com/google-cloud/building-a-go-web-app-from-scratch-to-deploying-on-google-cloud-part-1-building-a-simple-go-aee452a2e654
//?name=gjjggj

import (
	"fmt"
	"net/http"
	"time"
)

//Welcome struct
type Welcome struct {
	Name string
	Time string
}

func welcome() {

	http.HandleFunc("/", welcomeHandler)
	fmt.Println("listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	welcome := Welcome{"Juan", time.Now().Format(time.Stamp)}
	temp := Templates.Lookup("welcome") //template.Must(template.ParseFiles("templates/welcome-template.html"))
	if name := r.FormValue("name"); name != "" {
		welcome.Name = name
	}
	if err := temp.ExecuteTemplate(w, "welcome", welcome); err != nil { // "welcome-template.html", welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
