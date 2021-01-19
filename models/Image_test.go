package models

import (
	"reflect"
	"testing"
)

// TestImageModel_CreateImage 		PASS
func TestImageModel_CreateImage(t *testing.T) {
	SetupDatabase(true)
	tests := []struct {
		name string
		im   *ImageModel
		ex   *ImageModel
	}{
		{
			"Create",
			&ImageModel{
				ProjectID: 2,
				Titulo:    "PROJECTO 2 IMAGEN 2",
				Detalle:   "DETALLE 2.2",
			},
			&ImageModel{
				ID:        1, //4,
				ProjectID: 2,
				Titulo:    "PROJECTO 2 IMAGEN 2",
				Detalle:   "DETALLE 2.2",
			},
		}, {
			"Error",
			&ImageModel{
				ID:      99,
				Titulo:  "ERROR",
				Detalle: "ERROR",
			},
			nil,
		}, {
			"Inexistent",
			&ImageModel{
				Titulo:  "ERROR",
				Detalle: "ERROR",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ex *ImageModel
			err := tt.im.CreateImage()
			if err == nil {
				ex = &ImageModel{}
				rows, err := getDatabase().db.Query("SELECT image_id, titulo, detalle, project_id FROM images WHERE image_id=$1", tt.im.ID)
				if err != nil {
					t.Errorf(err.Error())
				}
				defer rows.Close()
				for rows.Next() {
					if err := rows.Scan(&ex.ID, &ex.Titulo, &ex.Detalle, &ex.ProjectID); err != nil {
						t.Errorf(err.Error())
					}
				}
				if !reflect.DeepEqual(&ex, &tt.ex) {
					t.Errorf("Expected: %v, got %v", tt.ex, ex)
				}
			}
		})
	}
}

// TestImageModel_ReadImage			PASS
func TestImageModel_ReadImage(t *testing.T) {
	SetupDatabase(true)
	tests := []struct {
		name string
		ID   int
		want []*ImageModel
	}{
		{
			"ID_1",
			1,
			[]*ImageModel{images[0], images[1]},
		}, {
			"ID_2",
			99,
			[]*ImageModel{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Image.ReadImage(tt.ID)
			for index := 0; index < len(got); index++ {
				if !reflect.DeepEqual(got[index], tt.want[index]) {
					t.Errorf("ImageModel.ReadImage() got1 = %v, want %v", got[index], tt.want[index])
				}
			}
		})
	}
}

//TestImageModel_UpdateImage		PASS
func TestImageModel_UpdateImage(t *testing.T) {
	SetupDatabase(true)
	tests := []struct {
		name string
		pr   int
		nt   int
		up   *ImageModel
	}{
		{
			"Update 1",
			1,
			1,
			images[3],
		}, {
			"Update 2",
			3,
			1,
			images[3],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetupDatabase(true)
			images := Image.ReadImage(tt.pr)
			for _, image := range images {
				if image.ID == tt.nt {
					image.Titulo = tt.up.Titulo
					image.Detalle = tt.up.Detalle
					tt.up.ID = image.ID
					tt.up.ProjectID = image.ProjectID
					image.UpdateImage()
				}
			}
			images = Image.ReadImage(tt.pr)
			for _, image := range images {
				if image.ID == tt.nt {
					if !reflect.DeepEqual(image, tt.up) {
						t.Errorf("ImageModel.UpdateImage() got1 = %v, want %v", Image, tt.up)
					}
				}
			}

		})
	}
}

// TestImageModel_DeleteImage		PASS
func TestImageModel_DeleteImage(t *testing.T) {
	SetupDatabase(true)
	tests := []struct {
		name string
		im   *ImageModel
	}{
		{
			"Delete",
			images[0],
		},
		{
			"Error",
			&ImageModel{
				ID: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.im.DeleteImage()
			got := Image.ReadImage(tt.im.ProjectID)
			for _, image := range got {
				if image.ID == tt.im.ID {
					t.Errorf("ImageModel.DeleteImage() got1 = %v, want %v", image, nil)
				}
			}
		})
	}
}
