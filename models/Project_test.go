package models

import (
	"reflect"
	"testing"
)

//TestProjectModel_CreateProject	PASS // NOTES => PASS // IMAGES => PASS //  LENGUAJES =>	PASS
func TestProjectModel_CreateProject(t *testing.T) {
	tests := []struct {
		name string
		pr   *ProjectModel
	}{
		{
			"Prueba_1",
			projects[0],
		}, {
			"Prueba_2",
			projects[1],
		},
	}
	for _, tt := range tests {
		SetupDatabase(false)
		t.Run(tt.name, func(t *testing.T) {
			var ex = &ProjectModel{}
			tt.pr.Lenguajes = lenguajes
			tt.pr.CreateProject()
			ex, _, err := Project.ReadProject(tt.pr.ID)
			if err != nil {
				t.Errorf(err.Error())
			}
			//ex.Lenguajes = readRelationships(ex.ID)
			compareNotes(tt.pr.Notas, ex.Notas, t)
			compareImages(tt.pr.Galeria, ex.Galeria, t)
			compareLenguajes(tt.pr.Lenguajes, ex.Lenguajes, t)
			// ex.Notas, tt.pr.Notas = nil, nil
			// ex.Galeria, tt.pr.Galeria = nil, nil
			if !reflect.DeepEqual(&ex, &tt.pr) {
				t.Errorf("Expected: %v, got %v", tt.pr, ex)
			}
		})
	}
}

//TestProjectModel_ReadProject		PASS // NOTES => PASS // IMAGES => PASS //  LENGUAJES =>	PASS
func TestProjectModel_ReadProject(t *testing.T) {
	SetupDatabase(true)
	tests := []struct {
		name    string
		args    []int
		wantOne *ProjectModel
		wantAll []*ProjectModel
	}{
		{
			"Uno",
			[]int{1},
			projects[0],
			nil,
		}, {
			"Todos",
			[]int{},
			nil,
			[]*ProjectModel{
				{
					ID:       1,
					Titulo:   "Personal",
					Detalle:  "uno muy bonito",
					Progreso: 44,
				},
				{
					ID:       2,
					Titulo:   "Artemis",
					Detalle:  "mi favorito",
					Progreso: 44,
				},
			},
		}, {
			"Error",
			[]int{3},
			nil,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			one, all, err := Project.ReadProject(tt.args...)
			if err != nil {
				t.Errorf(err.Error())
			}
			if tt.wantOne != nil {
				tt.wantOne.Lenguajes = lenguajes
			}
			if len(tt.args) > 0 {
				if one != nil {
					compareNotes(tt.wantOne.Notas, one.Notas, t)
					// one.Notas, tt.wantOne.Notas = nil, nil
					compareLenguajes(tt.wantOne.Lenguajes, one.Lenguajes, t)
					compareImages(tt.wantOne.Galeria, one.Galeria, t)
					one.Galeria, tt.wantOne.Galeria = nil, nil
				}
				if !reflect.DeepEqual(one, tt.wantOne) {
					t.Errorf("ProjectModel.ReadProject() got = %v, want %v", one, tt.wantOne)
				}
			} else {
				for index, val := range all {
					if val.ID == 1 {
						tt.wantAll[index].Lenguajes = lenguajes
					}
					// compareLenguajes(tt.wantAll[index].Lenguajes, val.Lenguajes, t)
					if !reflect.DeepEqual(val, tt.wantAll[index]) {
						t.Errorf("ProjectModel.ReadProject() got1 = %v, want %v", val, tt.wantAll[index])
					}
				}
			}

		})
	}
}

// TestProjectModel_UpdateProject	PASS // NOTES => PASS // IMAGES => PASS //  LENGUAJES => 	PASS
func TestProjectModel_UpdateProject(t *testing.T) {
	SetupDatabase(true)
	a := projects[1]
	a.Notas = projects[0].Notas
	a.Galeria = projects[0].Galeria
	a.ID = 1
	a.Notas[0].Titulo = notes[3].Titulo
	a.Notas[0].Detalle = notes[3].Detalle
	a.Galeria[0].Titulo = images[3].Titulo
	a.Galeria[0].Detalle = images[3].Detalle
	a.Notas[0].Fecha = notes[3].Fecha
	a.Lenguajes = []*LenguageModel{lenguajes[0], lenguajes[2], {Titulo: "Javascript", ID: 4}}
	tests := []struct {
		name string
		pr   *ProjectModel
		ex   *ProjectModel
	}{
		{
			"Prueba_1",
			projects[0],
			a,
		},
		{
			"Error",
			&ProjectModel{
				ID: 10,
				Notas: []*NoteModel{
					{},
				},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pr.Titulo = "Artemis"
			tt.pr.Detalle = "mi favorito"
			tt.pr.Descripcion = "hi patito feo"
			tt.pr.Notas[0].Titulo = notes[3].Titulo
			tt.pr.Notas[0].Detalle = notes[3].Detalle
			tt.pr.Notas[0].Fecha = notes[3].Fecha
			tt.pr.Lenguajes = []*LenguageModel{lenguajes[0], lenguajes[2], {Titulo: "Javascript"}}
			tt.pr.UpdateProject()
			got, _, err := Project.ReadProject(tt.pr.ID)
			if err != nil {
				t.Errorf(err.Error())
			}
			if got != nil {
				compareNotes(tt.ex.Notas, got.Notas, t)
				got.Notas, tt.ex.Notas = nil, nil
				compareImages(tt.ex.Galeria, got.Galeria, t)
				compareLenguajes(tt.ex.Lenguajes, got.Lenguajes, t)
				got.Galeria, tt.ex.Galeria = nil, nil
			}
			if !reflect.DeepEqual(got, tt.ex) {
				t.Errorf("ProjectModel.UpdateProject() got1 = %v, want %v", got, tt.ex)
			}

		})
	}
}

// TestProjectModel_DeleteProject	PASS // NOTES => PASS // IMAGES => PASS //  LENGUAJES =>	PASS
func TestProjectModel_DeleteProject(t *testing.T) {
	SetupDatabase(true)
	tests := []struct {
		name string
		pr   *ProjectModel
	}{
		{
			"Delete",
			projects[0],
		},
		{
			"Error",
			&ProjectModel{
				ID: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pr.DeleteProject()
			got, _, err := Project.ReadProject(tt.pr.ID)
			if err != nil {
				t.Errorf(err.Error())
			}
			if got != nil {
				t.Errorf("ProjectModel.DeleteProject() got1 = %v, want %v", got, nil)
			}
		})
	}
}

func compareNotes(exp, got []*NoteModel, t *testing.T) {
	if len(got) != len(exp) {
		t.Errorf("Mismatching lenghts")
	}
	if len(got) != 0 {
		for index := range exp {
			if !reflect.DeepEqual(&exp[index], &got[index]) {
				t.Errorf("Expected: %v, got %v", exp[index], got[index])
			}
		}
	}
}

func compareImages(exp, got []*ImageModel, t *testing.T) {
	if len(got) != len(exp) {
		t.Errorf("Mismatching lenghts")
	}
	if len(got) != 0 {
		for index := range exp {
			if !reflect.DeepEqual(&exp[index], &got[index]) {
				t.Errorf("Expected: %v, got %v", exp[index], got[index])
			}
		}
	}
}

func compareLenguajes(exp, got []*LenguageModel, t *testing.T) {
	if len(got) != len(exp) {
		t.Errorf("Mismatching lenghts")
	}
	if len(got) != 0 {
		for index := range exp {
			if !reflect.DeepEqual(&exp[index], &got[index]) {
				t.Errorf("Expected: %v, got %v", exp[index], got[index])
			}
		}
	}
}
