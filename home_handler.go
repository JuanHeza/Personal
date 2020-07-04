package main

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"net/http"
	"os"
	"strings"
)

// Projects General data
type Projects struct {
	Name         string   `json:"name,omitempty"`
	Language     []string `json:"language,omitempty"`
	Introduccion string   `json:"introduccion,omitempty"`
	Description  string   `json:"description,omitempty"`
	Icon         string   `json:"icon,omitempty"`
	Banner       string   `json:"banner,omitempty"`
	Progress     int      `json:"progress,omitempty"`
	Time         []string
	Side         int
	Models       []Model    `json:"models,omitempty"`
	Functions    []Function `json:"functions,omitempty"`
	Tasks        []Task     `json:"tasks,omitempty"`
	Notes        []Note     `json:"notes,omitempty"`
	Images       []Image    `json:"images,omitempty"`
}

//Model of the data in the proyect
type Model struct {
	Title string `json:"title,omitempty"`
	Data  []Data `json:"data,omitempty"`
}

//Data is the info of each field in the model
type Data struct {
	Name        string `json:"name,omitempty"`
	DataType    string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

//Function is something
type Function struct {
	ID int
	Call        string `json:"call,omitempty"`
	Return      string `json:"return,omitempty"`
	Description string `json:"description,omitempty"`
	Codigo      string
}

//Task is something
type Task struct {
	ID int
	Done bool   `json:"done,omitempty"`
	Text string `json:"text,omitempty"`
}

//Note is something
type Note struct {
	ID int
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}

//Image is something
type Image struct {
	Src   string `json:"src,omitempty"`
	Title string `json:"title,omitempty"`
}

//ProyectData stores the data of each proyect to be accesible everytime its nedeed
var ProyectData map[string]Projects

//General data of te proyects
var General []Projects

func homeHandler(w http.ResponseWriter, r *http.Request) {
	home := Templates.Lookup("home")
	data, err := store.GetProyect()
	IfErr(err, w, r)
	if err := home.ExecuteTemplate(w, "home", data); err != nil {
		//	if err := home.ExecuteTemplate(w, "home", General); err != nil {
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

//ReadJSONProyect is something
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

//Join convert the slice into an array
func (pr *Projects) Join() string {
	return strings.Join(pr.Language, ", ")
}
