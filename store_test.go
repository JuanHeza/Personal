package main

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

//StoreSuite struct
type StoreSuite struct {
	suite.Suite
	store *dbStore
	db    *sql.DB
}

/*
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "JHZ697heza"
	dbname   = "proyect_encyclopedia"
) */

//SetupSuite create a db connection
func (s *StoreSuite) SetupSuite() {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

//SetupTest deletes all entries in the table for the test
func (s *StoreSuite) SetupTest() {
	_, err := s.db.Query("DELETE FROM proyects")
	if err != nil {
		s.T().Fatal(err)
	}
}

//TearDownSuite close the db connection
func (s *StoreSuite) TearDownSuite() {
	s.db.Close()
}

//TestStoreSuite run the test bellow
func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

//TestCreateProyect test the CreateProyect method
func (s *StoreSuite) TestCreateProyect() {
	s.store.CreateProyect(&Proyect{
		Description: "test Description",
		Data:        "test proyect",
	})

	res, err := s.db.Query(`SELECT COUNT(*) FROM proyects WHERE description='test Description' AND proyect='test proyect'`)
	if err != nil {
		s.T().Fatal(err)
	}

	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}

//TestGetProyect test the GetProyect method
// func (s *StoreSuite) TestGetProyect() {
// 	_, err := s.db.Query(`INSERT INTO proyects (proyect, description) VALUES ('#Proyect','#Description')`)
// 	if err != nil {
// 		s.T().Fatal(err)
// 	}

// 	proyects, err := s.store.GetProyect()
// 	if err != nil {
// 		s.T().Fatal(err)
// 	}

// 	nProyects := len(proyects)
// 	if nProyects != 1 {
// 		s.T().Errorf("incorrect count, wanted 1, got %d", nProyects)
// 	}

// 	expectedProyect := Proyect{"#Proyect", "#Description"}
// 	if *proyects[0] != expectedProyect {
// 		s.T().Errorf("incorrect details, expected %v, got %v", expectedProyect, *proyects[0])
// 	}
// }
