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

func getOneProjectHandler(w http.ResponseWriter, r *http.Request) {
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

// func updateProjectHandler(w http.ResponseWriter, r *http.Request) {
// 	proyect := mux.Vars(r)
// 	err := store.UpdateProject(proyect["id"])
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Redirect(w, r, "/Error", http.StatusFound)
// 	}
// 	http.Redirect(w, r, "/Crud", http.StatusFound)
// }

func deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	proyect := mux.Vars(r)
	err := store.DeleteProject(proyect["id"])
	IfErr(err, w, r)
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
	return "/static/images/" + folder + "/"
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

func updateProjectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "updateProjectHandler")
}

func createModelHandler(w http.ResponseWriter, r *http.Request) {
	model := Model{}
	dat := Data{}
	proyect := mux.Vars(r)
	err := r.ParseForm()
	IfErr(err, w, r)
	model.Title = r.Form.Get("titulo")
	campos := r.Form["campo[]"]
	datos := r.Form["dato[]"]
	for i := range campos {
		dat.Name = campos[i]
		dat.DataType = datos[i]
		model.Data = append(model.Data, dat)
	}
	fmt.Println("Model: ", model, "Proyecto", proyect["proyecto"])
	err = store.CreateModel(&model, proyect["proyecto"])
	IfErr(err, w, r)
	fmt.Println("redirect", err)
	http.Redirect(w, r, "/Crud/"+proyect["proyecto"], http.StatusFound)
}
func updateModelHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Request: ", r.Body, r.URL)
	model := Model{}
	dat := Data{}
	proyect := mux.Vars(r)
	fmt.Println(proyect)
	err := r.ParseForm()
	IfErr(err, w, r)
	model.Title = r.FormValue("titulo")
	campos := r.Form["campo[]"]
	datos := r.Form["dato[]"]
	for i := range campos {
		dat.Name = campos[i]
		dat.DataType = datos[i]
		model.Data = append(model.Data, dat)
	}
	fmt.Println("Model: ", model, "Proyecto", proyect["proyecto"])
	err = store.UpdateModel(&model, proyect["proyecto"])
	IfErr(err, w, r)
	fmt.Println("redirect", err, r)
	http.Redirect(w, r, "/Crud/"+proyect["proyecto"], http.StatusFound)
}
func deleteModelHandler(w http.ResponseWriter, r *http.Request) {
	proyect := mux.Vars(r)
	err := store.DeleteModel(&Model{Title: proyect["id"]}, proyect["proyecto"])
	IfErr(err, w, r)
	http.Redirect(w, r, "/Crud/"+proyect["proyecto"], http.StatusFound)
}

func parseFuncion(fn *Function, r *http.Request) {
	fn.Call = r.Form.Get("llamada")
	fn.Description = r.Form.Get("funcion")
	fn.Return = r.Form.Get("return")
	fn.Codigo = r.Form.Get("codigo")
}
func createFunctionHandler(w http.ResponseWriter, r *http.Request) {
	function := Function{}
	proyect := mux.Vars(r)
	err := r.ParseForm()
	IfErr(err, w, r)
	parseFuncion(&function, r)
	fmt.Println("Funcion: ", function, "Proyecto: ", proyect["proyecto"])
	err = store.CreateFunction(&function, proyect["proyecto"])
	IfErr(err, w, r)
	http.Redirect(w, r, "/Crud/"+proyect["proyecto"], http.StatusFound)
}
func updateFunctionHandler(w http.ResponseWriter, r *http.Request) {
	funcion := Function{}
	proyect := mux.Vars(r)
	err := r.ParseForm()
	IfErr(err, w, r)
	parseFuncion(&funcion, r)
	fmt.Println("Funcion: ", funcion, "Proyecto: ", proyect["proyecto"])
	err = store.UpdateFunction(&funcion, proyect["proyecto"])
	IfErr(err, w, r)
	http.Redirect(w, r, "/Crud/"+proyect["prtoyecto"], http.StatusFound)
}
func deleteFunctionHandler(w http.ResponseWriter, r *http.Request) {
	proyect := mux.Vars(r)
	id, err := strconv.Atoi(proyect["id"])
	err = store.DeleteFunction(&Function{ID: id}, proyect["proyecto"])
	IfErr(err, w, r)
	http.Redirect(w, r, "/Crud", http.StatusFound)
}

func parseNota(nt *Note, r *http.Request) {
	nt.Title = r.Form.Get("nota")
	nt.Text = r.Form.Get("cuerpo")
}
func createNotasHandler(w http.ResponseWriter, r *http.Request) {
	nota := Note{}
	vars := mux.Vars(r)
	err := r.ParseForm()
	IfErr(err, w, r)
	parseNota(&nota, r)
	err = store.CreateNotas(&nota, vars["proyecto"])
	IfErr(err, w, r)
	http.Redirect(w, r, "/Crud/"+vars["proyecto"], http.StatusFound)
}
func updateNotasHandler(w http.ResponseWriter, r *http.Request) {
	nota := Note{}
	vars := mux.Vars(r)
	err := r.ParseForm()
	IfErr(err, w, r)
	parseNota(&nota, r)
	err = store.UpdateNotas(&nota, vars["proyecto"])
	IfErr(err, w, r)
	http.Redirect(w, r, "/Crud/"+vars["proyecto"], http.StatusFound)
}
func deleteNotasHandler(w http.ResponseWriter, r *http.Request) {
	proyect := mux.Vars(r)
	id, err := strconv.Atoi(proyect["id"])
	err = store.DeleteNotas(&Note{ID: id}, proyect["proyecto"])
	IfErr(err, w, r)
	http.Redirect(w, r, "/Crud"+proyect["proyecto"], http.StatusFound)
}

func parseTask(ts *Task, r *http.Request) {
	ts.Text = r.Form.Get("tarea")
	Done := r.Form.Get("completo")
	if Done == "true" {
		ts.Done = true
	} else {
		ts.Done = false
	}
}
func createTareasHandler(w http.ResponseWriter, r *http.Request) {
	tarea := Task{}
	vars := mux.Vars(r)
	err := r.ParseForm()
	IfErr(err, w, r)
	parseTask(&tarea, r)
	fmt.Println("Tarea", tarea, "Proyecto", vars["proyecto"])
	err = store.CreateTareas(&tarea, vars["proyecto"])
	IfErr(err, w, r)
	http.Redirect(w, r, "/Crud"+vars["proyecto"], http.StatusFound)
}
func updateTareasHandler(w http.ResponseWriter, r *http.Request) {
	tarea := Task{}
	vars := mux.Vars(r)
	err := r.ParseForm()
	IfErr(err, w, r)
	parseTask(&tarea, r)
	fmt.Println("Tarea", tarea, "Proyecto", vars["proyecto"])
	err = store.UpdateTareas(&tarea, vars["proyecto"])
	IfErr(err, w, r)
	http.Redirect(w, r, "/Crud"+vars["proyecto"], http.StatusFound)
}
func deleteTareasHandler(w http.ResponseWriter, r *http.Request) {
	proyect := mux.Vars(r)
	id, err := strconv.Atoi(proyect["id"])
	err = store.DeleteTareas(&Task{ID: id}, proyect["proyecto"])
	IfErr(err, w, r)
	http.Redirect(w, r, "/Crud", http.StatusFound)
}

func deleteFolder(dir string) {}
