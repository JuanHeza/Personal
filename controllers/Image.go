package controllers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/JuanHeza/Personal/models"
	"github.com/gorilla/mux"
)

func deleteImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if IfErr(err, w, r) {
		log.Printf("/Delete/Image/%v @ Project.deleteImageHandler", id)
		im := &models.ImageModel{ID: id}
		//DeleteFolder()
		im.DeleteImage()
		path := fmt.Sprintf("/static/images/%v/%v", im.ProjectID, im.Titulo)
		if IfErr(DeletePicture(path), w, r) {
			http.Redirect(w, r, "/Edit/", http.StatusFound)
		}
	}
}

func uploadGalleryImage(folder string /* file multipart.File, */, handler []*multipart.FileHeader, captions []string) (imageList []*models.ImageModel) { //, err error) {
	// https://tutorialedge.net/golang/go-file-upload-tutorial/
	for index := range handler {
		// fmt.Println(captions)
		// fmt.Println(titles)
		img := &models.ImageModel{}
		file, err := handler[index].Open()
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		out, err := os.Create(folder + handler[index].Filename)

		defer out.Close()
		if err != nil {
			fmt.Println("Unable to create the file for writing. Check your write access privilege")
			return
		}

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			fmt.Println(err)
			return
		}
		img.Titulo = handler[index].Filename
		img.Detalle = captions[index]
		// img.Titulo = titles[index]
		// fmt.Println("Files uploaded successfully : ")
		// fmt.Println(handler[index].Filename + "\n")
		imageList = append(imageList, img)
	}
	return
}

func uploadImage(folder string, file multipart.File, handler *multipart.FileHeader, err error, name ...string) (string, error) {
	var f *os.File
	if err != nil {
		return "", err
	}
	// fmt.Printf("Uploaded File: %+v \t %+v \n", handler.Filename, handler.Size)
	defer file.Close() //close the file when we finish
	//this is path which  we want to store the file
	// fmt.Println(folder)
	// fmt.Println(name)
	if len(name) > 0 {
		fmt.Println("File:", folder+name[0])
		f, err = os.OpenFile(folder+name[0], os.O_WRONLY|os.O_CREATE, 0666)
	} else {
		f, err = os.OpenFile(folder+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	}
	// fmt.Printf("%s\n", f.Name())
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)
	if len(name) > 0 {
		return name[0], nil
	}
	return handler.Filename, nil
}

//DeletePicture is for delete the pic in the given path
func DeletePicture(path string) (err error) {
	fmt.Println("Deleting pictures")
	err = os.Remove("." + path)

	if err != nil {
		return
	}
	fmt.Println("File " + path + " successfully deleted")
	return
}
