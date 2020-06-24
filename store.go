package main

import (
	"database/sql"

	"github.com/lib/pq"
)

//Store get and create data
type Store interface {
	CreateProyect(pr *Proyect) error
	CreateProject(pr *Projects) error
	DeleteProject(pr string) error
	GetProyect(id ...string) ([]*Projects, error)
}

//dbStore implements Store interface & use the connection object
type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateProyect(pr *Proyect) error {
	_, err := store.db.Query("INSERT INTO proyects(proyect,description) VALUES ($1,$2)", pr.Data, pr.Description)
	return err
}

func (store *dbStore) DeleteProject(pr string) error {
	_, err := store.db.Query("DELETE FROM informacion_general	WHERE nombre = $1", pr)
	return err
}

func (store *dbStore) CreateProject(pr *Projects) error {
	_, err := store.db.Query("INSERT INTO Informacion_General(	nombre, lenguajes, descripcion, introduccion, progreso, icon, banner) VALUES ($1,$2,$3,$4,$5,$6,$7);", pr.Name, pq.Array(pr.Language), pr.Description, pr.Introduccion, pr.Progress, pr.Icon, pr.Banner)
	return err
}

func (store *dbStore) GetProyect(id ...string) ([]*Projects, error) {
	var rows *sql.Rows
	var err error
	if len(id) > 0 {
		rows, err = store.db.Query("SELECT nombre, descripcion, icon, banner, progreso, introduccion, lenguajes FROM informacion_general WHERE nombre = $1", id[0])
	} else {
		rows, err = store.db.Query("SELECT nombre, descripcion, icon, banner, progreso, introduccion, lenguajes FROM informacion_general")
	}
	//rows, err := store.db.Query("SELECT proyect, description from proyects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	proyects := []*Projects{}
	for rows.Next() {
		proyect := &Projects{}
		if err := rows.Scan(&proyect.Name, &proyect.Description, &proyect.Icon, &proyect.Banner, &proyect.Progress, &proyect.Introduccion, pq.Array(&proyect.Language)); err != nil {
			return nil, err
		}
		proyects = append(proyects, proyect)
	}
	return proyects, nil
}

var store Store

//InitStore method
func InitStore(s Store) {
	store = s
}
