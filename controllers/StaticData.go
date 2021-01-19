package controllers

import (
	"log"
	"net/http"

	"fmt"

	"github.com/JuanHeza/Personal/models"
)

//StaticDataHandler is a representation of the staticData
type StaticDataHandler struct {
	data models.StaticData
}

//Static is an instance to access the methods
var Static StaticDataHandler

//GetStatics return a certain value in the map
func GetStatics(op int) string {
	var leng = "es"
	if UserSession != nil {
		leng = fmt.Sprint(UserSession.Values["leng"])
	}
	var st models.StaticData
	st = models.StaticDataCollection[leng].(models.StaticData)
	switch op {
	case 1:
		return st.Introduccion
	case 2:
		return st.About
	case 3:
		return st.Tutorial
	case 4:
		return st.Contacto
	case 5:
		return leng
	}
	return ""
}

//UpdateStatics IS #ERROR
func UpdateStatics(w http.ResponseWriter, r *http.Request) {
	log.Printf("/Update/Data/ @ StaticData.UpdateStatics")
	var multipart = true
	folder := "/static/stylesheet/icons"
	err := r.ParseMultipartForm(10 << 20) //
	if err != nil {
		err = r.ParseForm()
		multipart = false
	}
	if IfErr(err, w, r) {
		data := &models.StaticData{Introduccion: r.Form.Get("introduccion"), About: r.Form.Get("about"), Tutorial: r.Form.Get("tutorial"), Contacto: r.Form.Get("contact"), Leng: r.Form.Get("leng")}
		fmt.Println(data)

		if multipart {
			icon, handler, err := r.FormFile("avatar")
			uploadImage(folder, icon, handler, err, "avatar.png")

			banner, handler, err := r.FormFile("banner")
			uploadImage(folder, banner, handler, err, "background.png")
			links := r.Form["links[]"]
			fmt.Println(r.MultipartForm)
			icons := r.MultipartForm.File["linkImage[]"]
			uploadGalleryImage(folder, icons, links)
			for i := range links {
				newLink := &models.LinkModel{Link: links[i], Icon: icons[i].Filename}
				err := newLink.CreateLink()
				if err != nil {
					panic(err)
				}
			}
		}
		data.UpdateStatics()
		http.Redirect(w, r, "/Edit/", http.StatusFound)
	}
}
