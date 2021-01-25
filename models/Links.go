package models

import (
	"fmt"
)

//LinkModel is an instance with the data of link section
type LinkModel struct {
	ID   int
	Icon string
	Link string
}

var (
	//Link variable to accesss to the methods
	Link *LinkModel
)

//CreateLink create a new link record
func (ln *LinkModel) CreateLink() error {
	err := Queries.CreateLink(ln)
	if err != nil {
		panic(err)
	}
	// fmt.Println(im)
	return nil
}

// ReadLink fills the staticdata map with the contact links
func (ln *LinkModel) ReadLink() {
	linkList, err := Queries.ReadLink()
	if err != nil {
		panic(err)
	}
	StaticDataCollection["link"] = linkList
}

// UpdateLink updates the data of the record with the ID
func (ln *LinkModel) UpdateLink() {
	_, err := Queries.UpdateLink(ln)
	if err != nil {
		panic(err)
	}
}

//DeleteLink deletes the record with the given ID
func (ln *LinkModel) DeleteLink() string {
	icon, err := Queries.DeleteLink(ln)
	if err != nil {
		panic(err)
	}
	return icon
}

//===============================================================================================
//============================================ Links ============================================
//===============================================================================================

func (query *dbStore) CreateLink(ln *LinkModel) error {
	data, err := query.db.Query(`INSERT INTO links(icon , link ) VALUES ($1,$2);`, ln.Icon, ln.Link)
	defer data.Close()
	if err != nil {
		return err
	}
	return nil
}

func (query *dbStore) ReadLink() ([]*LinkModel, error) {
	links, err := query.db.Query("SELECT link, icon, link_id FROM links")
	if err != nil {
		fmt.Println("ReadLink:", err.Error())
		return nil, err
	}
	var ln []*LinkModel
	defer links.Close()
	for links.Next() {
		var link = &LinkModel{}
		if err = links.Scan(&link.Link, &link.Icon, &link.ID); err != nil {
			fmt.Println("ReadStatics in SEcond Query:", err)
			return nil, err
		}
		ln = append(ln, link)
	}
	return ln, nil
}

//UpdateLink updates the link instance
func (query *dbStore) UpdateLink(ln *LinkModel) (string, error) {
	if ln.Icon != "" {
		path, err := query.db.Query("SELECT icon FROM links WHERE link_id = $1", ln.Icon)
		var icon string
		defer path.Close()
		for path.Next() {
			path.Scan(&icon)
		}
		_, err = query.db.Query("UPDATE links SET link = $1, icon = $3 WHERE link_id = $2", ln.Link, ln.ID, ln.Icon)
		return icon, err
	}
	data, err := query.db.Query("UPDATE links SET link = $1 WHERE link_id = $2", ln.Link, ln.ID)
	defer data.Close()
	Link.ReadLink()
	return "", err
}

//DeleteLink deletes the link with the given id
func (query *dbStore) DeleteLink(ln *LinkModel) (string, error) {
	path, err := query.db.Query("SELECT icon FROM links WHERE link_id = $1", ln.ID)
	fmt.Println("path ", err, ln.ID)
	var icon string
	defer path.Close()
	for path.Next() {
		path.Scan(&icon)
		fmt.Println(icon)
	}
	_, err = query.db.Query("DELETE FROM links WHERE link_id = $1", ln.ID)
	return icon, err
}
