package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Proyect struct
type Proyect struct {
	Data        string `json:"data"`
	Description string `json:"description"`
}

func getProyectHandler(w http.ResponseWriter, r *http.Request) {

	proyects, err := store.GetProyect()

	proyectListBytes, err := json.Marshal(proyects)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v",err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(proyectListBytes)
}

func createProyectHandler(w http.ResponseWriter, r *http.Request) {
	proyect := Proyect{}
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	proyect.Data = r.Form.Get("data")
	proyect.Description = r.Form.Get("description")

	err = store.CreateProyect(&proyect)
	if err != nil{
		fmt.Println(err)
	}

	http.Redirect(w, r, "/templates/", http.StatusFound)
}
