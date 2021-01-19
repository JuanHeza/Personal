package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JuanHeza/Personal/models"
	"github.com/gorilla/mux"
)

func deleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if IfErr(err, w, r) {
		log.Printf("/Delete/Link/%v @ Project.deleteLinkHandler", id)
		ln := &models.LinkModel{ID: id}
		//DeleteFolder()
		if IfErr(DeletePicture("/static/stylesheets/icons/"+ln.DeleteLink()), w, r) {
			http.Redirect(w, r, "/Edit/", http.StatusFound)
		}
	}
}

func updateLinkHandler(w http.ResponseWriter, r *http.Request) {
	var link string
	//	var link []string
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	log.Printf("/Update/Link/%v @ Project.updateLinkHandler", id)
	if err != nil {
		//fmt.Println("Error at Models.UpdateLinkHanlder: ", err)
	}
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		//fmt.Println("Parseform")
		err = r.ParseForm()
		IfErr(err, w, r)
	}
	//	link = append(link, r.Form["links"]...)
	link = r.Form.Get("links")
	// log.Println(link, id)
	// log.Println("Multipart: ", r.MultipartForm)
	// log.Println("Form: ", r.Form)
	if ok := r.MultipartForm.File["linkImage"]; ok != nil {
		icon, handler, err := r.FormFile("linkImage")
		file, err := uploadImage("./statics/stylesheets/icons", icon, handler, err)
		var ret string
		if IfErr(err, w, r) {
			ret, err = models.Queries.UpdateLink(&models.LinkModel{ID: id, Link: link, Icon: file})
		}
		// iconName, err := models.Queries.UpdateLink(id, link[0], handler.Filename)
		log.Println(ret, " == ", r.Form.Get("icon"))
		DeletePicture("." + r.Form.Get("icon"))
	} else {
		// log.Println("NO")
		//		_, err = models.Queries.UpdateLink(&models.LinkModel{ID: id, Link: link[0]})
		_, err = models.Queries.UpdateLink(&models.LinkModel{ID: id, Link: link})
	}
	if err != nil {
		fmt.Println(err)
	}
}

/*


// DeleteLinkHandler IS #ERROR
func DeleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		icon, err := models.Queries.DeleteLink(&models.LinkModel{ID: id})
		fmt.Println(err)
		DeletePicture("." + icon)
	}
	fmt.Println(id)
}

*/
