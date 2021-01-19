package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/JuanHeza/Personal/models"
	"github.com/gorilla/mux"
)

func deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if IfErr(err, w, r) {
		log.Printf("/Delete/Note/%v @ Project.deleteNoteHandler", id)
		nt := &models.NoteModel{ID: id}
		nt.DeleteNote()
		http.Redirect(w, r, "/Edit/", http.StatusFound)
	}
}
