package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/JuanHeza/Personal/models"
	"github.com/gorilla/mux"
)

//HomeHandler is the handler of the homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/Home/ @ Project.HomeHandler")
	_, data, err := models.Project.ReadProject()
	if IfErr(err, w, r) {
		var Data = map[string]interface{}{
			"Data": data,
		}
		tryTemplate("home", Data, w, r)
	}
}

// ProjectHandler is the handler of the project detail page
func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	var actual *models.ProjectModel
	var Data = map[string]interface{}{
		"CSS": []string{"/static/stylesheets/project.css"},
		"JS":  []string{"/static/scripts/script.js", "https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"},
	}
	vars := mux.Vars(r)
	log.Printf("/Project/%v @ Project.ProjectHandler", vars["id"])
	id, err := strconv.Atoi(vars["id"])
	if IfErr(err, w, r) {
		actual, _, err = models.Project.ReadProject(id)
		actual.Posts = models.Post.ReadProjectPost(id)
		if IfErr(err, w, r) {
			// actual.Time = getWakaTime(vars["name"])
			Data["Data"] = actual
			tryTemplate("project", Data, w, r)
		}
	}
}

//ProjectListHandler is the handler to the list of all projects and when a language is selected
func ProjectListHandler(w http.ResponseWriter, r *http.Request) {
	var list interface{}
	var err error
	var Data = map[string]interface{}{
		"CSS": []string{"/static/stylesheets/lists.css"},
	}
	vars := mux.Vars(r)
	log.Printf("/ProjectList/%v @ Project.ProjectListHandler", vars["leng"])
	if vars["leng"] != "" {
		id, err := strconv.Atoi(vars["leng"])
		IfErr(err, w, r)
		list = models.Lenguage.ReadAllByLenguage(id)
	} else {
		_, list, err = models.Project.ReadProject()
		IfErr(err, w, r)
	}
	Data["Data"] = list
	// actual.Time = getWakaTime(vars["name"])
	tryTemplate("projectList", Data, w, r)
}

// ProjectEditHandler if for the form of update and create
func ProjectEditHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var Data = map[string]interface{}{
		"JS":  []string{"https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js", "http://malsup.github.com/jquery.form.js", "https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.3/js/select2.min.js", "/static/scripts/dragNdrop.js", "/static/scripts/script.js"},
		"CSS": []string{"https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.3/css/select2.css", "/static/stylesheets/dragNdrop.css"},
	}
	vars := mux.Vars(r)
	if val, ok := vars["id"]; ok {
		id, err := strconv.Atoi(val)
		log.Printf("/Edit/Project/%v @ Project.ProjectEditHandler", id)
		if IfErr(err, w, r) {
			if id != 0 {
				Data["Data"], _, err = models.Project.ReadProject(id)
			}
		}
	} else {
		log.Printf("/Edit/Project/ @ Project.ProjectEditHandler")
	}
	if IfErr(err, w, r) {
		fmt.Println("Data: ", Data["Data"])
		tryTemplate("projectForm", Data, w, r)
	}
}

type data struct {
	One  *models.ProjectModel
	Many []*models.ProjectModel
}

func createProjectHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/Create/Project/ @ Project.CreateProjectHandler")
	pr := &models.ProjectModel{}
	pr, err := parseProjectForm(r)
	if IfErr(err, w, r) {
		pr.CreateProject()
		fmt.Println(pr)
		handleMultipart(pr, r)
		pr.UpdateProject()
		http.Redirect(w, r, "/Edit/", http.StatusFound)
	}
}

func updateProjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	log.Printf("/Update/Project/%v @ Project.updateProjectHandler", id)
	pr := &models.ProjectModel{}
	pr, err = parseProjectForm(r)
	if IfErr(err, w, r) {
		pr.UpdateProject()
		fmt.Println(pr)
		http.Redirect(w, r, "/Edit/", http.StatusFound)
	}
}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	log.Printf("/Delete/Project/%v @ Project.deleteProjectHandler", id)
	pr := &models.ProjectModel{ID: id}
	pr.DeleteProject()
	os.Remove("./static/images/" + vars["id"])
	//DeleteFolder()
	if IfErr(err, w, r) {
		http.Redirect(w, r, "/Edit/", http.StatusFound)
	}
}

func parseProjectForm(r *http.Request) (pr *models.ProjectModel, err error) {
	var lengList []*models.LenguageModel
	var noteList []*models.NoteModel
	var imageList []*models.ImageModel
	err = r.ParseMultipartForm(10 << 20) //r.ParseForm()
	if err != nil {
		err = r.ParseForm()
		fmt.Println("NO MULTIPART")
		if err != nil {
			return
		}
	}
	pr = &models.ProjectModel{}
	log.Println(r.Form) //
	pr.ID, err = strconv.Atoi(r.Form.Get("project_id"))
	if err != nil {
		panic(err)
	}
	pr.Titulo = r.Form.Get("nombre")
	pr.Detalle = r.Form.Get("detalle")
	pr.Github = r.Form.Get("github")
	pr.Link = r.Form.Get("link")
	pr.Descripcion = r.Form.Get("descripcion")
	pr.WakaLinks = r.Form["wakalinks"]
	pr.Progreso, err = strconv.Atoi(r.Form.Get("progreso"))
	if err != nil {
		panic(err)
	}

	lengs := r.Form["lenguajes[]"]
	for _, i := range lengs {
		lengList = append(lengList, &models.LenguageModel{Titulo: i})
	}
	pr.Lenguajes = lengList
	noteTitle := r.Form["note_title[]"]
	noteBody := r.Form["note_body[]"]
	noteDate := r.Form["note_date[]"]
	noteID := r.Form["note_id[]"]
	fmt.Println(len(noteTitle))
	if noteTitle[0] != "" {
		for i := range noteTitle {
			fecha, er := time.Parse("2006-01-02", noteDate[i])
			if er != nil {
				panic(er)
			}
			id, er := strconv.Atoi(noteID[i])
			if er != nil {
				panic(er)
			}
			note := &models.NoteModel{Titulo: noteTitle[i], Detalle: noteBody[i], Fecha: fecha, ID: id}
			fmt.Println(note)
			noteList = append(noteList, note)
		}
		pr.Notas = noteList
	}

	imgUpdateTitle := r.Form["update_title[]"]
	imgUpdateCaption := r.Form["update_caption[]"]
	imgUpdateID := r.Form["update_index[]"]

	for i := range imgUpdateID {
		id, er := strconv.Atoi(imgUpdateID[i])
		if er != nil {
			panic(er)
		}
		image := &models.ImageModel{Titulo: imgUpdateTitle[i], Detalle: imgUpdateCaption[i], ID: id}
		log.Println(&image)
		imageList = append(imageList, image)
	}
	pr.Galeria = imageList
	return
}

func handleMultipart(pr *models.ProjectModel, r *http.Request) {
	var imageList []*models.ImageModel
	folder := folderExist(pr.ID)

	icon, handler, err := r.FormFile("icon")
	if err == nil {
		_, err = uploadImage(folder, icon, handler, err, "Icon.png")
		if err != nil {
			panic(err)
		}
	}

	banner, handler, err := r.FormFile("banner")
	if err == nil {
		_, err = uploadImage(folder, banner, handler, err, "Banner.png")
		if err != nil {
			panic(err)
		}
	}

	galleryFiles, okFiles := r.MultipartForm.File["gallery[]"]
	captionList, okCaption := r.Form["caption[]"]
	//gallery, handler, err := r.FormFile("gallery")
	if okCaption && okFiles {
	}
	imageList = append(imageList, uploadGalleryImage(folder, galleryFiles, captionList)...)
	pr.Galeria = append(pr.Galeria, imageList...)
}
