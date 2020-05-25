package main

import (
	"database/sql"
)

//Store get and create data
type Store interface {
	CreateProyect(pr *Proyect) error
	GetProyect() ([]*Proyect, error)
}

//dbStore implements Store interface & use the connection object
type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateProyect(pr *Proyect) error {
	_, err := store.db.Query("INSERT INTO proyects(proyect,description) VALUES ($1,$2)", pr.Data, pr.Description)
	return err
}

func (store *dbStore) GetProyect() ([]*Proyect, error) {
	rows, err := store.db.Query("SELECT proyect, description from proyects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	proyects := []*Proyect{}
	for rows.Next() {
		proyect := &Proyect{}

		if err := rows.Scan(&proyect.Data, &proyect.Description); err != nil {
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
