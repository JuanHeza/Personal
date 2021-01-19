package models

import (
	_ "fmt"
	"reflect"
	"testing"
)

// TestNoteModel_CreateNote		PASS
func TestNoteModel_CreateNote(t *testing.T) {
	SetupDatabase(true)
	type args struct {
		projectsID int
	}
	tests := []struct {
		name string
		nt   *NoteModel
		ex   *NoteModel
	}{
		{
			"Create",
			&NoteModel{
				ProjectID: 2,
				Titulo:    "PROJECTO 1 NOTA 1",
				Fecha:     fecha,
				Detalle:   "DETALLE 1.1",
			},
			&NoteModel{
				ID:        4,
				ProjectID: 2,
				Titulo:    "PROJECTO 1 NOTA 1",
				Fecha:     fecha,
				Detalle:   "DETALLE 1.1",
			},
		}, {
			"Error",
			&NoteModel{
				ID:      99,
				Titulo:  "ERROR",
				Fecha:   fecha,
				Detalle: "ERROR",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ex *NoteModel
			err := tt.nt.CreateNote()
			if err == nil {
				ex = &NoteModel{}
				rows, err := getDatabase().db.Query("SELECT note_id, titulo, detalle, fecha, project_id FROM notes WHERE note_id=$1", tt.nt.ID)
				if err != nil {
					t.Errorf(err.Error())
				}
				defer rows.Close()
				for rows.Next() {
					if err := rows.Scan(&ex.ID, &ex.Titulo, &ex.Detalle, &ex.Fecha, &ex.ProjectID); err != nil {
						t.Errorf(err.Error())
					}
				}
				ex.Fecha = ex.Fecha.UTC()
			}
			if !reflect.DeepEqual(&ex, &tt.ex) {
				t.Errorf("Expected: %v, got %v", tt.ex, ex)
			}
		})
	}
}

// TestNoteModel_ReadNote		PASS
func TestNoteModel_ReadNote(t *testing.T) {
	SetupDatabase(true)
	tests := []struct {
		name string
		ID   int
		ex   []*NoteModel
	}{
		{
			"ID_1",
			1,
			[]*NoteModel{notes[0], notes[1]},
		}, {
			"ID_2",
			99,
			[]*NoteModel{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			all := Note.ReadNote(tt.ID)
			for index := 0; index < len(all); index++ {
				if !reflect.DeepEqual(all[index], tt.ex[index]) {
					t.Errorf("NoteModel.ReadNote() got1 = %v, want %v", all[index], tt.ex[index])
				}
			}
		})
	}
}

//TestNoteModel_UpdateNote		PASS
func TestNoteModel_UpdateNote(t *testing.T) {
	tests := []struct {
		name string
		pr   int
		nt   int
		up   *NoteModel
	}{
		{
			"Update 1",
			1,
			1,
			notes[3],
		}, {
			"Update 2",
			2,
			1,
			notes[3],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetupDatabase(true)
			notes := Note.ReadNote(tt.pr)
			for _, note := range notes {
				if note.ID == tt.nt {
					note.Titulo = tt.up.Titulo
					note.Detalle = tt.up.Detalle
					note.Fecha = tt.up.Fecha
					tt.up.ID = note.ID
					tt.up.ProjectID = note.ProjectID
					note.UpdateNote()
				}
			}
			notes = Note.ReadNote(tt.pr)
			for _, note := range notes {
				if note.ID == tt.nt {
					if !reflect.DeepEqual(note, tt.up) {
						t.Errorf("NoteModel.UpdateNote() got1 = %v, want %v", note, tt.up)
					}
				}
			}

		})
	}
}

// TestNoteModel_DeleteNote 	PASS
func TestNoteModel_DeleteNote(t *testing.T) {
	SetupDatabase(true)
	tests := []struct {
		name string
		nt   *NoteModel
	}{
		{
			"Delete",
			notes[0],
		},
		{
			"Error",
			&NoteModel{
				ID: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.nt.DeleteNote()
			got := Note.ReadNote(tt.nt.ProjectID)
			for _, note := range got {
				if note.ID == tt.nt.ID {
					t.Errorf("NoteModel.DeleteNote() got1 = %v, want %v", note, nil)
				}
			}
		})
	}
}
