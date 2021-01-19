package models

import (
	"reflect"
	"testing"
)

// TestLenguageModel_CreateLenguage		PASS
func TestLenguageModel_CreateLenguage(t *testing.T) {
	SetupDatabase(false)
	tests := []struct {
		name string
		ln   *LenguageModel
		want error
	}{
		{
			"JAVASCRIPT",
			&LenguageModel{
				Titulo: "Javascript",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ln.CreateLenguage(); err != tt.want {
				t.Errorf("LenguageModel.CreateLenguage() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

// TestLenguageModel_ReadByName			PASS
func TestLenguageModel_ReadByName(t *testing.T) {
	SetupDatabase(true)
	tests := []struct {
		name    string
		title   string
		want    *LenguageModel
		wantErr error
	}{
		{
			"GO",
			"Go",
			&LenguageModel{
				ID:     1,
				Titulo: "Go",
			},
			nil,
		},
		{
			"ERROR",
			"PHP",
			nil,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lenguage.ReadByName(tt.title)
			if err != tt.wantErr {
				t.Errorf("LenguageModel.ReadByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LenguageModel.ReadByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestLenguageModel_ReadByID		PASS
func TestLenguageModel_ReadByID(t *testing.T) {
	SetupDatabase(true)
	tests := []struct {
		name    string
		ID      int
		want    *LenguageModel
		wantErr error
	}{
		{
			"GO",
			1,
			&LenguageModel{
				ID:     1,
				Titulo: "Go",
			},
			nil,
		},
		{
			"ERROR",
			99,
			nil,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lenguage.ReadByID(tt.ID)
			if err != tt.wantErr {
				t.Errorf("LenguageModel.ReadByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LenguageModel.ReadByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

//TestLenguageModel_ReadAll			PASS
func TestLenguageModel_ReadAll(t *testing.T) {
	SetupDatabase(true)
	tests := []struct {
		name string
		want []*LenguageModel
	}{
		{
			"Todos",
			lenguajes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Lenguage.ReadAll()
			for index, one := range tt.want {
				if !reflect.DeepEqual(got[index], one) {
					t.Errorf("LenguageModel.ReadAll() = %v, want %v", got[index], one)
				}
			}
		})
	}
}

//Test_createRelationship			PASS
func Test_createRelationship(t *testing.T) {
	SetupDatabase(true)
	type args struct {
		lengs []*LenguageModel
		proj  int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			"Prueba #1",
			args{
				[]*LenguageModel{{Titulo: "Go"}, {Titulo: "CSS"}, {Titulo: "HTML"}},
				1,
			},
			[][2]int{{1, 1}, {1, 2}, {1, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var all [][2]int
			createRelationship(tt.args.lengs, tt.args.proj)
			rows, err := getDatabase().db.Query("SELECT project_id, lenguaje_id FROM proj_leng WHERE project_id=$1", tt.args.proj)
			if err != nil {
				t.Errorf(err.Error())
			}
			defer rows.Close()
			for rows.Next() {
				var one [2]int
				if err := rows.Scan(&one[0], &one[1]); err != nil {
					t.Errorf(err.Error())
				}
				all = append(all, one)
			}
			if !reflect.DeepEqual(all, tt.want) {
				t.Errorf("createRelationships() = %v, want %v", all, tt.want)
			}
		})
	}
}

//Test_readRelationships			PASS
func Test_readRelationships(t *testing.T) {
	SetupDatabase(true)
	type args struct {
		proj      int
		relations []*LenguageModel
	}
	tests := []struct {
		name string
		args args
		want []*LenguageModel
	}{
		{
			"Prueba 1",
			args{1, []*LenguageModel{{Titulo: "Go"}, {Titulo: "CSS"}, {Titulo: "HTML"}}},
			lenguajes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createRelationship(tt.args.relations, tt.args.proj)
			got := readRelationships(tt.args.proj)
			if len(got) != len(tt.want) {
				t.Errorf("Mismatching lenghts")
			}
			for index := range got {
				if !reflect.DeepEqual(got[index], tt.want[index]) {
					t.Errorf("readRelationships() = %v, want %v", got[index], tt.want[index])
				}
			}
		})
	}
}

// Test_updateRelationships			PASS
func Test_updateRelationships(t *testing.T) {
	SetupDatabase(true)
	type args struct {
		lengs []*LenguageModel
		proj  int
	}
	tests := []struct {
		name string
		args args
		want []*LenguageModel
	}{
		{
			"Prueba",
			args{[]*LenguageModel{{Titulo: "Go"}, {Titulo: "CSS"}, {Titulo: "Javascript"}}, 1},
			[]*LenguageModel{
				{
					ID:     1,
					Titulo: "Go",
				}, {
					ID:     2,
					Titulo: "CSS",
				}, {
					ID:     4,
					Titulo: "Javascript",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateRelationships(tt.args.lengs, tt.args.proj)
			got := readRelationships(tt.args.proj)
			if len(got) != len(tt.want) {
				t.Errorf("Mismatching lenghts")
			}
			for index := range got {
				if !reflect.DeepEqual(got[index], tt.want[index]) {
					t.Errorf("readRelationships() = %v, want %v", got[index], tt.want[index])
				}
			}
		})
	}
}
