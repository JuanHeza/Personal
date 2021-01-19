package main

import (
	"errors"

)

//Store get and create data
type Store interface {
	CreateLink(ln string, in string) error
	UpdateLink(id int, ln string, in ...string) error
	DeleteLink(id int) error
}

var store Store

//InitStore method
func InitStore(s Store) {
	store = s
}

// CreateLink is something
func (store *dbStore) CreateLink(ln string, in string) error {
	_, err := store.db.Query("INSERT INTO links(link, icon) VALUES ($1,$2)", ln, in)
	return err
}

// UpdateLink is something
// func (store *dbStore) UpdateLink(id int, ln string, in ...string) error {
// 	if len(in) > 1 {
// 		path, err := store.db.Query("SELECT icon FROM links WHERE id = $1", id)
// 		var icon string
// 		defer path.Close()
// 		for path.Next() {
// 			path.Scan(&icon)
// 			DeletePicture("."+icon)
// 		}
// 		_, err = store.db.Query("UPDATE links SET link = $1, icon = $3 WHERE id = $2", ln, id)
// 		return err
// 	}
// 	_, err := store.db.Query("UPDATE links SET link = $1 WHERE id = $2", ln, id)
// 	return err
// }

// DeleteLink is something
// func (store *dbStore) DeleteLink(id int) error {
// 	path, err := store.db.Query("SELECT icon FROM links WHERE id = $1", id)
// 	fmt.Println("path ", err, id)
// 	var icon string
// 	defer path.Close()
// 	for path.Next() {
// 		path.Scan(&icon)
// 		fmt.Println(icon)
// 		DeletePicture("."+icon)
// 	}
// 	_, err = store.db.Query("DELETE FROM links WHERE id = $1", id)
// 	return err
// }