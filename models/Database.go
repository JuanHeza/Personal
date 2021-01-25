package models

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq" // to get workiing postgresql
)

/*							NOTAS
hay que ver que pasa con el login que me pide a fuerzas que le envie el queries
*/

//Query interface of al the query list
type Query interface {
	LogIn(us *User) error
	CreateStatics(st *StaticData) error
	ReadStatics(sdC map[string]interface{}) error
	UpdateStatics(st *StaticData) error

	CreateLink(ln *LinkModel) error
	ReadLink() ([]*LinkModel, error)
	UpdateLink(ln *LinkModel) (string, error)
	DeleteLink(ln *LinkModel) (string, error)

	CreateProject(pr *ProjectModel) (int, error)
	ReadProject(id int) (*ProjectModel, error)
	ReadAllProject() ([]*ProjectModel, error)
	UpdateProject(pr *ProjectModel) error
	DeleteProject(pr *ProjectModel) error

	CreateNote(nt *NoteModel) error
	ReadNote(id int) ([]*NoteModel, error)
	UpdateNote(nt *NoteModel) error
	DeleteNote(nt *NoteModel) error

	CreateImage(nt *ImageModel) error
	ReadImage(id int) ([]*ImageModel, error)
	UpdateImage(nt *ImageModel) error
	DeleteImage(nt *ImageModel) error

	CreateLenguage(ln *LenguageModel) error
	ReadLenguageTitle(titulo string) (*LenguageModel, error)
	ReadLenguageID(id int) (*LenguageModel, error)
	ReadAllLenguages() ([]*LenguageModel, error)
	// UpdateLenguage(ln *LenguageModel) error
	// DeleteLenguage(ln *LenguageModel) error

	CreateRelationship(leng, proj int) error
	ReadRelationship(proj int) ([]int, error)
	ReadLenguageRelationship(leng int) (map[int]*ProjectModel, error)
	DeleteRelationship(proj int) error

	CreatePost(pt *PostModel) (err error)
	ReadPost(id ...int) (one *PostModel, many []*PostModel, err error)
	ReadAllProjectPosts(pr int) (all []*PostModel, err error)
	DeletePost(pt *PostModel) (err error)
	UpdatePost(pt *PostModel) error

	CreatePostRelationship(ln *LenguageModel, pt int) (err error)
	ReadPostRelationship(post int) (lengs []*LenguageModel, err error)
	DeletePostRelationship(pt int) (err error)
}

//dbStore implements Store interface & use the connection object
type dbStore struct {
	db *sql.DB
}

var (
	//Queries conect the functions to the interface to start a query
	Queries  Query
	database *dbStore
)

//InitQueries creates the interface
func InitQueries(q Query) {
	Queries = q
}

//StartConnection conects to the database with the provided credentials
func StartConnection(host string, port string, user string, password string, dbname string) {
	var connString string
	if os.Getenv("development") == "true" {
		connString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	} else {
		connString = os.Getenv("DATABASE_URL")
	}
	fmt.Println(connString)
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	database = &dbStore{db: db}
	InitQueries(database)
}

//GetDatabase Provides the database
func getDatabase() *dbStore {
	return database
}

//LogIn returns the data of the user
func (query *dbStore) LogIn(us *User) error {
	found := false
	users, err := query.db.Query("SELECT * FROM usuarios WHERE id = $1 AND password = $2", us.ID, us.Password)
	// users, err := query.db.Query("SELECT admin FROM usuarios WHERE id = $1 AND password = $2", us.ID, us.Password)
	defer users.Close()

	for users.Next() {
		found = true
		if err = users.Scan(&us.ID, &us.Password, &us.Admin); err != nil {
			// if err = users.Scan(&us.Admin); err != nil {
			return err
		}
	}
	if found {
		return nil
	}
	return errors.New("NOT FOUND")
}
