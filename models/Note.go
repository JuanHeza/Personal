package models

import (
	"errors"
	"fmt"
	"time"
)

//NoteModel is the structure for the Development notes of each Project ~~ Many Notes Belongs to One Project
type NoteModel struct {
	ID        int
	ProjectID int
	Titulo    string
	Fecha     time.Time
	Detalle   string
}

var (
	//Note variable to access the methods
	Note NoteModel
)

// CreateNote creates a new record in the database associating it wit the given projects_id
func (nt *NoteModel) CreateNote() error {
	if nt.ProjectID == 0 {
		return errors.New(fmt.Sprint("Invalid ProjectID ", nt))
	}
	err := Queries.CreateNote(nt)
	if err != nil {
		panic(err)
	}
	return nil
}

// ReadNote get all the notes of the project
func (nt *NoteModel) ReadNote(id int) []*NoteModel {
	all, err := Queries.ReadNote(id)
	if err != nil {
		panic(err)
	}
	return all
}

// UpdateNote updates the gven note
func (nt *NoteModel) UpdateNote() {
	err := Queries.UpdateNote(nt)
	if err != nil {
		panic(err)
	}
}

// DeleteNote deletes the given note
func (nt *NoteModel) DeleteNote() {
	err := Queries.DeleteNote(nt)
	if err != nil {
		panic(err)
	}
}

//===============================================================================================
//=========================================== QUERIES ===========================================
//===============================================================================================

func (query *dbStore) CreateNote(nt *NoteModel) error {
	var ID int
	data, err := query.db.Query(`INSERT INTO notes(titulo, detalle, fecha, project_id) VALUES ($1,$2,$3,$4);`, nt.Titulo, nt.Detalle, nt.Fecha, nt.ProjectID)
	defer data.Close()
	rows, err := query.db.Query("SELECT note_id FROM notes WHERE titulo=$1 AND project_id = $2", nt.Titulo, nt.ProjectID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&ID); err != nil {
			return err
		}
	}
	nt.ID = ID
	return nil
}

func (query *dbStore) ReadNote(id int) ([]*NoteModel, error) {
	var all []*NoteModel
	var ex *NoteModel
	rows, err := getDatabase().db.Query("SELECT note_id, titulo, detalle, fecha, project_id FROM notes WHERE project_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		ex = &NoteModel{}
		if err := rows.Scan(&ex.ID, &ex.Titulo, &ex.Detalle, &ex.Fecha, &ex.ProjectID); err != nil {
			return nil, err
		}
		ex.Fecha = ex.Fecha.UTC()
		all = append(all, ex)
	}
	return all, nil
}

func (query *dbStore) UpdateNote(nt *NoteModel) error {
	data, err := query.db.Query(`UPDATE notes SET titulo=$1, fecha=$2, detalle=$3 WHERE note_id=$4 AND project_id=$5;`, nt.Titulo, nt.Fecha, nt.Detalle, nt.ID, nt.ProjectID)
	defer data.Close()
	return err
}

func (query *dbStore) DeleteNote(nt *NoteModel) error {
	data, err := query.db.Query(`DELETE FROM notes WHERE note_id=$1 AND project_id=$2;`, nt.ID, nt.ProjectID)
	defer data.Close()
	return err
}
