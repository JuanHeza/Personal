package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// CreateHandler handler for all the create routes
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	// if cr := IsAdminAutenticathed(); cr.Admin && cr.Auth {
	vars := mux.Vars(r)
	switch vars["model"] {
	case "Project":
		createProjectHandler(w, r)
		break
	case "Post":
		createPostHandler(w, r)
		break
	// case "Image":
	// 	break
	// case "Note":
	// 	break
	case "Link":
		break
	default:
		http.Redirect(w, r, "/Error/", http.StatusFound)
		break
	}
	// } else {
	// 	http.Redirect(w, r, "/Error/", http.StatusFound)
	// }
}

// UpdateHandler handler for all the update routes
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	// if cr := IsAdminAutenticathed(); cr.Admin && cr.Auth {
	vars := mux.Vars(r)
	switch vars["model"] {
	case "Project":
		updateProjectHandler(w, r)
		break
	case "Post":
		updatePostHandler(w, r)
		break
	// case "Image":
	// 	break
	// case "Note":
	// 	break
	case "Link":
		updateLinkHandler(w, r)
		break
	case "Data":
		UpdateStatics(w, r)
		break
	default:
		http.Redirect(w, r, "/Error/", http.StatusFound)
		break
	}
	// } else {
	// 	http.Redirect(w, r, "/Error/", http.StatusFound)
	// }
}

// DeleteHandler handler for all the delete routes
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// if cr := IsAdminAutenticathed(); cr.Admin && cr.Auth {
	vars := mux.Vars(r)
	switch vars["model"] {
	case "Project":
		deleteProjectHandler(w, r)
		break
	case "Post":
		deletePostHandler(w, r)
		break
	case "Image":
		deleteImageHandler(w, r)
		break
	case "Note":
		deleteNoteHandler(w, r)
		break
	case "Link":
		deleteLinkHandler(w, r)
		break
	default:
		http.Redirect(w, r, "/Error/", http.StatusFound)
		break
	}
	// } else {
	// 	http.Redirect(w, r, "/Error/", http.StatusFound)
	// }
}
