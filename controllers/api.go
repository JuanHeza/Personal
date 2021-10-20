package controllers

import (
	"encoding/json"

	// "fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JuanHeza/Personal/models"
	"github.com/gorilla/mux"
)

//ApiHomeHandler is the handler of the homepage
func ApiHomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/api/Home/ @ Project.ApiHomeHandler")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	_, data, err := models.Project.ReadProject()
	st := models.StaticDataCollection["es"].(models.StaticData)
	// log.Println(st)
	if IfErr(err, w, r) {
		var Data = map[string]interface{}{
			"Projects":  data,
			"Posts":     "nil",
			"Statics":   st,
			"Links":     models.StaticDataCollection["link"],
			"Lenguages": models.Lenguage.ReadAll(),
		}
		json.NewEncoder(w).Encode(Data)
	}
}

func IconHandler(w http.ResponseWriter, r *http.Request) {
	var dir string
	vars := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if vars["folder"] == "Icons" || vars["folder"] == "icons" {
		dir = "./static/stylesheets/icons/" + vars["title"] + ".png"
	} else {
		dir = "./static/images/" + vars["folder"] + "/" + vars["title"] + ".png"
	}
	log.Println(dir)
	http.ServeFile(w, r, dir)
}

// ApiProjectHandler is the handler of the project detail page
func ApiProjectHandler(w http.ResponseWriter, r *http.Request) {
	var actual *models.ProjectModel
	var Data = map[string]interface{}{
		"Data": &models.ProjectModel{},
	}
	vars := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	log.Printf("/Project/%v @ Project.ApiProjectHandler", vars["id"])
	id, err := strconv.Atoi(vars["id"])
	if IfErr(err, w, r) {
		actual, _, err = models.Project.ReadProject(id)
		actual.Posts = models.Post.ReadProjectPost(id)
		if IfErr(err, w, r) {
			// actual.Time = getWakaTime(vars["name"])
			Data["Data"] = actual
			json.NewEncoder(w).Encode(Data)
		}
	}
}

// ApiPostHandler is the handler of the project detail page
func ApiPostHandler(w http.ResponseWriter, r *http.Request) {
	var Data = map[string]interface{}{
		"Data": &models.PostModel{},
	}
	vars := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	log.Printf("/Post/%v @ Post.ApiPostHandler", vars["id"])
	id, err := strconv.Atoi(vars["id"])
	IfErr(err, w, r)
	Data["Data"], _, err = models.Post.ReadPost(id)
	// actual.Time = getWakaTime(vars["name"])
	if IfErr(err, w, r) {
		json.NewEncoder(w).Encode(Data)
	}
}

func ApiListHandler(w http.ResponseWriter, r *http.Request) {
	var Data = make(map[string]interface{})
	vars := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch vars["filtro"] {
	case "Post":
		Data["data"], _ = ApiPostListHanler()
		// Handler(w,r)
		break
	case "Project":
		Data["data"], _ = ApiProjectListHandler()
		break
	default:
		// Handler(w,r)
		id, err := strconv.Atoi(vars["id"])
		if err == nil {
			data, _ := ApilenguageListHandler(id)
			Data["data"] = data
		}
		break
	}
	json.NewEncoder(w).Encode(Data)
}

func ApiProjectListHandler() (list interface{}, err error) {
	_, list, err = models.Project.ReadProject()
	return
}

func ApiPostListHanler() (list interface{}, err error) {
	_, list, err = models.Post.ReadPost()
	return
}
func ApilenguageListHandler(id int) (list map[string]interface{}, err error) {
	list = make(map[string]interface{})
	list["Project"] = models.Lenguage.ReadAllByLenguage(id)
	list["Post"] = models.Post.ReadAllByLenguage(id)
	return
}
