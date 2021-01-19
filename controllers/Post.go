package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/JuanHeza/Personal/models"
	"github.com/gorilla/mux"
)

//PostHandler is the handler of the /Post/{ID} route
func PostHandler(w http.ResponseWriter, r *http.Request) { //	FUNCIONAL
	var Data = map[string]interface{}{
		"CSS": []string{"/static/stylesheets/post.css"},
		// "JS":  []string{"/static/scripts/script.js", "https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"},
	}
	vars := mux.Vars(r)
	log.Printf("/Post/%v @ Post.PostHandler", vars["id"])
	id, err := strconv.Atoi(vars["id"])
	IfErr(err, w, r)
	Data["Data"], _, err = models.Post.ReadPost(id)
	fmt.Println("Actual: ", Data["Data"], Data["Data"].(*models.PostModel).Fecha)
	// actual.Time = getWakaTime(vars["name"])
	if IfErr(err, w, r) {
		tryTemplate("post", Data, w, r)
	}
}

//PostListHandler is the habler of /Post/
func PostListHandler(w http.ResponseWriter, r *http.Request) { //	FUNCIONAL
	var Data = map[string]interface{}{
		"CSS": []string{"/static/stylesheets/lists.css"},
	}
	var err error
	log.Printf("/PostList/ @ Post.PostListHandler")
	_, Data["Data"], err = models.Post.ReadPost()
	if IfErr(err, w, r) {
		tryTemplate("postList", Data, w, r)
	}
}

// PostEditHandler if for the form of update and create
func PostEditHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var Data = map[string]interface{}{
		"JS":  []string{"https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js", "http://malsup.github.com/jquery.form.js", "https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.3/js/select2.min.js", "/static/scripts/forms.js"},
		"CSS": []string{"https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.3/css/select2.css", "/static/stylesheets/estilos.css", "/static/stylesheets/forms.css"},
	}
	vars := mux.Vars(r)
	if val, ok := vars["id"]; ok {
		id, err := strconv.Atoi(val)
		log.Printf("/Edit/Post/%v @ Post.PostEditHandler", id)
		if IfErr(err, w, r) {
			if id != 0 {
				Data["Data"], _, err = models.Post.ReadPost(id)
				IfErr(err, w, r)
			}
		}
	} else {
		log.Printf("/Edit/Post/ @ Post.PostEditHandler")
	}
	if IfErr(err, w, r) {
		tryTemplate("postEditor", Data, w, r)
	}
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/Create/Post/ @ Post.CreatePostHandler")
	pt := &models.PostModel{}
	pt, err := parsePostForm(r)
	IfErr(err, w, r)
	pt.CreatePost()
	http.Redirect(w, r, "/Edit/", http.StatusFound)
}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	log.Printf("/Update/Post/%v @ Post.updatePostHandler", id)
	pt := &models.PostModel{}
	pt, err = parsePostForm(r)
	if IfErr(err, w, r) {
		pt.UpdatePost()
		http.Redirect(w, r, "/Edit/", http.StatusFound)
	}
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	log.Printf("/Delete/Post/%v @ Post.deletePostHandler", id)
	pt := &models.PostModel{ID: id}
	if IfErr(err, w, r) {
		pt.DeletePost()
		//DeleteFolder()
		http.Redirect(w, r, "/Edit/", http.StatusFound)
	}
}

func parsePostForm(r *http.Request) (pt *models.PostModel, err error) {
	var lengList []*models.LenguageModel
	if err = r.ParseForm(); err != nil {
		return
	}
	pt = &models.PostModel{}
	pt.ID, err = strconv.Atoi(r.Form.Get("post_id"))
	if err != nil {
		panic(err)
	}
	pt.Titulo = r.Form.Get("nombre")
	pt.Detalle = r.Form.Get("detalle")
	pt.Cuerpo = r.Form.Get("cuerpo")
	pt.Replit = r.Form.Get("link")
	pt.Fecha, err = time.Parse("2006-01-02", r.Form.Get("fecha"))
	if err != nil {
		panic(err)
	}
	pt.ProjectID, err = strconv.Atoi(r.Form.Get("project"))
	if err != nil {
		panic(err)
	}
	lengs := r.Form["lenguajes[]"]
	for _, i := range lengs {
		lengList = append(lengList, &models.LenguageModel{Titulo: i})
	}
	pt.Lenguajes = lengList
	return
}
