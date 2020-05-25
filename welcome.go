package main

//https://medium.com/google-cloud/building-a-go-web-app-from-scratch-to-deploying-on-google-cloud-part-1-building-a-simple-go-aee452a2e654
//?name=gjjggj

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

//Welcome struct
type Welcome struct {
	Name string
	Time string
}

func welcome() {
	welcome := Welcome{"Juan", time.Now().Format(time.Stamp)}

	template := template.Must(template.ParseFiles("templates/welcome-template.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := template.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
