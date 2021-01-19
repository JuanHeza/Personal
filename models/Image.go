package models

import (
	"errors"
)

//ImageModel is the representation of an image with an alt text, a description and the link to the image
type ImageModel struct {
	ID        int
	ProjectID int
	Titulo    string
	Detalle   string
	Link      string
}

var (
	//Image variable to access the methods
	Image ImageModel
)

// CreateImage creates a new image record in the database and save the image in the project folder
func (im *ImageModel) CreateImage() error {
	if im.ProjectID == 0 {
		return errors.New("Invalid ProjectID")
	}
	err := Queries.CreateImage(im)
	if err != nil {
		panic(err)
	}
	return nil
}

// ReadImage retrive the images of the given project
func (im *ImageModel) ReadImage(ID int) []*ImageModel {
	images, err := Queries.ReadImage(ID)
	if err != nil {
		panic(err)
	}
	return images
}

// UpdateImage modify the caption or title of the given image
func (im *ImageModel) UpdateImage() {
	err := Queries.UpdateImage(im)
	if err != nil {
		panic(err)
	}
}

// DeleteImage deletes the image from the database and from the project folder
func (im *ImageModel) DeleteImage() {
	err := Queries.DeleteImage(im)
	if err != nil {
		panic(err)
	}
}

//===============================================================================================
//============================================ Image ============================================
//===============================================================================================

func (query *dbStore) CreateImage(im *ImageModel) error {
	var ID int
	_, err := query.db.Query(`INSERT INTO images(titulo, detalle, project_id) VALUES ($1,$2,$3);`,
		im.Titulo, im.Detalle, im.ProjectID)
	rows, err := query.db.Query("SELECT image_id FROM images WHERE titulo=$1 AND project_id = $2", im.Titulo, im.ProjectID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&ID); err != nil {
			return err
		}
	}
	im.ID = ID
	return nil
}

func (query *dbStore) ReadImage(id int) ([]*ImageModel, error) {
	var all []*ImageModel
	var ex *ImageModel
	rows, err := getDatabase().db.Query("SELECT image_id, titulo, detalle, project_id FROM images WHERE project_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		ex = &ImageModel{}
		if err := rows.Scan(&ex.ID, &ex.Titulo, &ex.Detalle, &ex.ProjectID); err != nil {
			return nil, err
		}
		all = append(all, ex)
	}
	return all, nil
}

func (query *dbStore) UpdateImage(im *ImageModel) error {
	_, err := query.db.Query(`UPDATE images SET titulo=$1, detalle=$2 WHERE image_id=$3 AND project_id=$4;`, im.Titulo, im.Detalle, im.ID, im.ProjectID)
	return err
}

func (query *dbStore) DeleteImage(im *ImageModel) error {
	rows, err:= query.db.Query(`SELECT titulo, project_id FROM images WHERE image_id = $1`,im.ID)
	defer rows.Close()
	for rows.Next(){
		if err = rows.Scan(&im.Titulo, &im.ProjectID); err != nil {
			return err
		}
	}
	_, err = query.db.Query(`DELETE FROM images WHERE image_id=$1;`, im.ID)
	return err
}
