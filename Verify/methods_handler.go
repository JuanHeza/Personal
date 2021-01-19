package main

//https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/#serving-static-files
import (
	"fmt"
	"io"
	_ "io/ioutil"
	"log"
	"mime/multipart"
	"os"
)


func uploadGalleryImage(folder string /* file multipart.File, */, handler []*multipart.FileHeader, titles []string, proyect ...string) { //, err error) {
	// https://tutorialedge.net/golang/go-file-upload-tutorial/
	for index := range handler {
		img := Image{}
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
		img.Src = folder + handler[index].Filename
		img.Title = titles[index]
		fmt.Println("Files uploaded successfully : ")
		fmt.Println(handler[index].Filename + "\n")
	}
	// if err != nil {
	// 	fmt.Println("Error Retrieving the File")
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()
	// // fmt.Printf("Uploaded File: %+v \t %+v \n", handler.Filename, handler.Size)
	// //name = proyectto-*.png
	// tempFile, err := ioutil.TempFile(folder, proyect+"-*.png")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer tempFile.Close()
	// // read all of the contents of our uploaded file into a
	// // byte array
	// fileBytes, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // write this byte array to our temporary file
	// tempFile.Write(fileBytes)
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
func deleteFolder(dir string) {}
