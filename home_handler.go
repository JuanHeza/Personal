package main

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"net/http"
	"os"
)

// Projects General data
type Projects struct {
	Name        string   `json:"name,omitempty"`
	Language    string   `json:"language,omitempty"`
	Description string   `json:"description,omitempty"`
	Icon        string   `json:"icon,omitempty"`
	Progress    float32  `json:"progress,omitempty"`
	Model       []Models `json:"models,omitempty"`
	Time        []string
	Side        int
}

//Models of the data in the proyect
type Models struct {
	Title string `json:"title,omitempty"`
	Data  []Data `json:"data,omitempty"`
}

//Data is the info of each field in the model
type Data struct {
	Name        string `json:"name,omitempty"`
	DataType    string `json:"data_type,omitempty"`
	Description string `json:"description,omitempty"`
}

//ProyectData stores the data of each proyect to be accesible everytime its nedeed
var ProyectData map[string]Projects

//General data of te proyects
var General []Projects

func homeHandler(w http.ResponseWriter, r *http.Request) {
	home := Templates.Lookup("home")
	if err := home.ExecuteTemplate(w, "home", General); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//ReadJSON reads the given file name
func ReadJSON() { //w http.ResponseWriter){
	dat, err := os.Open(ProyectFiles["General"])
	if err != nil {
		panic(err) //http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(dat)

	err = json.Unmarshal(buf.Bytes(), &General)
	if err != nil {
		panic(err) //http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ProyectData = make(map[string]Projects)
	for a := range General {
		General[a].Side = a % 2
		ProyectData[General[a].Name] = General[a]
	}
}
func ReadJSONProyect(file string) Projects { //w http.ResponseWriter){
	dat, err := os.Open(ProyectFiles[file])
	if err != nil {
		panic(err) //http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(dat)

	var ProyectData Projects
	err = json.Unmarshal(buf.Bytes(), &ProyectData)
	if err != nil {
		panic(err) //http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return ProyectData
}
