package models

import (
	"fmt"
)

//StaticData structure used to store the data of header and footer
type StaticData struct {
	Introduccion string
	About        string
	Tutorial     string
	Contacto     string
	// Link         [][]string /* array de [link, link image]*/
	// Avatar       string
	// Banner       string
	Leng string
}

//Statics is the interface to acces the functions
type Statics interface {
	ReadStatics()
	UpdateStatics(data *StaticData)
}

var (
	//Static variable to acces the methods
	Static StaticData
	//StaticDataCollection is the data used in the header and footer, Key: lenguage & Value: data
	StaticDataCollection = make(map[string]interface{})
)

//NewStaticData creates a new instance of the struct
func NewStaticData(intro string, about string, tuto string, contact string) *StaticData {
	return &StaticData{Introduccion: intro, About: about, Tutorial: tuto, Contacto: contact}
}

//CreateStatics seeds an instance for the object
func (s *StaticData) CreateStatics() {
	err := Queries.CreateStatics(s)
	if err != nil {
		panic(err)
	}
}

//ReadStatics fills the collection whit data
func (s *StaticData) ReadStatics() {
	err := Queries.ReadStatics(StaticDataCollection)
	ln := Link.ReadLink()
	StaticDataCollection["link"] = ln
	// fmt.Println(StaticDataCollection)
	if err != nil {
		panic(err)
	}
}

//UpdateStatics this updates the static data in the database
func (s *StaticData) UpdateStatics() {
	err := Queries.UpdateStatics(s)
	if err != nil {
		panic(err)
	}
	Static.ReadStatics()
}

// func (s *StaticData) deleteStatics() {}

//===============================================================================================
//=========================================== Statics ===========================================
//===============================================================================================

func (query *dbStore) CreateStatics(st *StaticData) error {
	data, err := query.db.Query(`INSERT INTO statics(about ,contacto ,introduccion, tutorial, leng ) VALUES ($1,$2,$3,$4,$5);`, st.About, st.Contacto, st.Introduccion, st.Tutorial, st.Leng)
	defer data.Close()
	if err != nil {
		return err
	}
	return nil
}

func (query *dbStore) ReadStatics(sdC map[string]interface{}) error {
	data, err := query.db.Query("SELECT * FROM statics;")
	defer data.Close()

	for data.Next() {
		st := StaticData{}
		if err = data.Scan(&st.Introduccion, &st.About, &st.Tutorial, &st.Contacto, &st.Leng); err != nil {
			fmt.Println("ReadStatics in First Query:", err)
			return err
		}
		sdC[st.Leng] = st
	}
	return nil
}

func (query *dbStore) UpdateStatics(st *StaticData) error {
	data, err := query.db.Query("UPDATE statics SET introduccion = $1, about = $2, tutorial = $3, contacto = $4 WHERE leng = $5", st.Introduccion, st.About, st.Tutorial, st.Contacto, st.Leng)
	defer data.Close()
	return err
}

//func (query *dbStore) DeleteStaics(){}
