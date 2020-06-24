package main

//https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/#serving-static-files
import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(proyectListBytes)
}

func getOneProjectHandler (w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)

	proyects, err := store.GetProyect(id["id"])

	proyectListBytes, err := json.Marshal(proyects)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(proyectListBytes)
}
func deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	proyect := mux.Vars(r)
	err := store.DeleteProject(proyect["id"])
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/Error", http.StatusFound)
	}
	http.Redirect(w, r, "/Crud", http.StatusFound)
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
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/templates/", http.StatusFound)
}

func createProject(w http.ResponseWriter, r *http.Request) {
	proyect := Projects{}
	err := r.ParseMultipartForm(10 << 20) //r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	proyect.Name = r.Form.Get("nombre")
	proyect.Language = r.Form["lenguaje[]"]
	proyect.Introduccion = r.Form.Get("introduccion")
	proyect.Description = r.Form.Get("descripcion")
	val, err := strconv.Atoi(r.Form.Get("progreso"))
	if err != nil {
		panic(err)
	}
	proyect.Progress = val

	folder := folderExist(proyect.Name)

	icon, handler, err := r.FormFile("icon")
	uploadImage(folder, "Icon.png", icon, handler, err)
	proyect.Icon = fmt.Sprintf("%sIcon.png", folder)

	banner, handler, err := r.FormFile("banner")
	uploadImage(folder, "Banner.png", banner, handler, err)
	proyect.Banner = fmt.Sprintf("%sBanner.png", folder)

	err = store.CreateProject(&proyect)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Crud", http.StatusFound)
}

func folderExist(folder string) string {
	_, err := os.Stat("./static/images/" + folder)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("./static/images/"+folder+"/", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
	return "./static/images/" + folder + "/"
}

func uploadGalleryImage(folder string, proyect string, file multipart.File, handler *multipart.FileHeader, err error) {
	// https://tutorialedge.net/golang/go-file-upload-tutorial/
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	// fmt.Printf("Uploaded File: %+v \t %+v \n", handler.Filename, handler.Size)
	//name = proyectto-*.png
	tempFile, err := ioutil.TempFile(folder, proyect+"-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
}

func uploadImage(folder string, name string, file multipart.File, handler *multipart.FileHeader, err error) (string, error) {
	if err != nil {
		return "", err
	}
	// fmt.Printf("Uploaded File: %+v \t %+v \n", handler.Filename, handler.Size)
	defer file.Close() //close the file when we finish
	//this is path which  we want to store the file
	// fmt.Println(folder)
	// fmt.Println(name)
	f, err := os.OpenFile(folder+name, os.O_WRONLY|os.O_CREATE, 0666)
	// fmt.Printf("%s\n", f.Name())
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)
	return name, nil
}
