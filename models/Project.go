package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/lib/pq"
)

//ProjectModel is the main structure of the page, stores the general data of every project
type ProjectModel struct {
	ID          int
	Titulo      string
	Lenguajes   []*LenguageModel
	Detalle     string
	Descripcion string
	Progreso    int
	Github      string
	Link        string
	// Icon        string
	// Banner      string
	Tiempo    time.Time //in minutes
	WakaLinks []string  // lins of wakatime to retrive the time
	Notas     []*NoteModel
	Galeria   []*ImageModel
	Posts     []*PostModel
}

var (
	//Project variable to access the methods
	Project ProjectModel
	// APIKey is the api of wakatime
	APIKey = os.Getenv("WAKATIME_API")
)

// NewProject is a new instance of the model
func NewProject(tit string, leng []*LenguageModel, det string, des string, pro int, git string, link string, time time.Time, waka []string, notes []*NoteModel, images []*ImageModel) *ProjectModel {
	return &ProjectModel{
		Titulo:      tit,
		Lenguajes:   leng,
		Detalle:     det,
		Descripcion: des,
		Progreso:    pro,
		Github:      git,
		Link:        link,
		Tiempo:      time,
		WakaLinks:   waka,
		Notas:       notes,
		Galeria:     images,
	}
}

//CreateProject create an instance of the model
func (pr *ProjectModel) CreateProject() {
	pr.Tiempo = getWakaTime(pr.WakaLinks)
	id, err := Queries.CreateProject(pr)
	if err != nil {
		panic(err)
	}
	pr.ID = id
	// fmt.Println("Project ID:", id)
	for _, nt := range pr.Notas {
		nt.ProjectID = id
		err = nt.CreateNote()
		if err != nil {
			panic(err)
		}
	}
	for _, im := range pr.Galeria {
		im.ProjectID = id
		err = im.CreateImage()
		if err != nil {
			panic(err)
		}
	}
	createRelationship(pr.Lenguajes, pr.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(pr)
}

//ReadProject read all projects or just the one with the given ID
func (pr *ProjectModel) ReadProject(id ...int) (*ProjectModel, []*ProjectModel, error) {
	var err error
	var one *ProjectModel
	var all []*ProjectModel
	if len(id) != 0 && id[0] != 0 {
		// fmt.Println("ProjectID:", id[0])
		one, err = Queries.ReadProject(id[0])
		if one != nil {
			// fmt.Println("Notes from ProjectID", one.ID)
			one.Notas = Note.ReadNote(one.ID)
			one.Galeria = Image.ReadImage(one.ID)
			one.Lenguajes = readRelationships(one.ID)
			one.Posts = Post.ReadProjectPost(one.ID)
		}
	} else {
		all, err = Queries.ReadAllProject()
		for _, one := range all {
			one.Lenguajes = readRelationships(one.ID)
		}
	}
	if err != nil {
		return nil, nil, err
	}
	return one, all, nil
}

// UpdateProject updates the project
func (pr *ProjectModel) UpdateProject() {
	err := Queries.UpdateProject(pr)
	updateRelationships(pr.Lenguajes, pr.ID)
	if err != nil {
		panic(err)
	}
	for _, nt := range pr.Notas {
		if nt.ID == 0 {
			nt.ProjectID = pr.ID
			err := nt.CreateNote()
			if err != nil {
				panic(err)
			}
		} else {
			nt.UpdateNote()
		}
	}
	for _, im := range pr.Galeria {
		if im.ID == 0 {
			im.ProjectID = pr.ID
			err := im.CreateImage()
			if err != nil {
				panic(err)
			}
		} else {
			im.UpdateImage()
		}
	}
}

// DeleteProject deletes the project
func (pr *ProjectModel) DeleteProject() {
	err := Queries.DeleteProject(pr)
	if err != nil {
		panic(err)
	}
}

func getWakaTime(links []string) (total time.Time) {
	var data interface{} //= make(map[string]string)
	for _, project := range links {
		url := fmt.Sprintf("https://wakatime.com/api/v1/users/current/stats/last_year?api_key=%v&project=%v", APIKey, project)
		res, err := http.Get(url)
		if err != nil {
			continue
			// panic(err)
		}
		defer res.Body.Close()
		decoder := json.NewDecoder(res.Body)
		decoder.Decode(&data)
		// fmt.Println(url)
		x := data.(map[string]interface{})
		y := x["data"].(map[string]interface{})
		if len(y["categories"].([]interface{})) > 0 {
			z := y["categories"].([]interface{})[0].(map[string]interface{})
			// total.Add(time.Second * time.Duration(z["seconds"].(float64)))
			total = total.Add(time.Second * time.Duration(z["total_seconds"].(float64)))
			// total.Add(time.Minute * time.Duration(z["minutes"].(float64)))
			// total.Add(time.Hour * time.Duration(z["hours"].(float64)))
			fmt.Println("Time: ", total, " Added: ", z["digital"])
		}
	}
	return
}

//===============================================================================================
//=========================================== QUERIES ===========================================
//===============================================================================================

func (query *dbStore) CreateProject(pr *ProjectModel) (int, error) {
	var ID int
	today := time.Now().Truncate(time.Hour * 24)
	data, err := query.db.Query(`INSERT INTO projects(	titulo, detalle, descripcion, progreso, github, link, tiempo, wakalinks, updated) VALUES ($1,$2,$3,$4,$5,$6,$7,$8, $9);`, pr.Titulo, pr.Detalle, pr.Descripcion, pr.Progreso, pr.Github, pr.Link, pr.Tiempo, pq.Array(pr.WakaLinks), today)
	defer data.Close()
	if err != nil {
		return 0, err
	}
	rows, err := query.db.Query("SELECT project_id FROM projects WHERE titulo=$1 AND detalle = $2", pr.Titulo, pr.Detalle)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&ID); err != nil {
			return 0, err
		}
	}
	return ID, err
}

func (query *dbStore) ReadProject(id int) (*ProjectModel, error) {
	var one *ProjectModel
	rows, err := query.db.Query("SELECT project_id, titulo, detalle, descripcion, progreso, github, link, tiempo, wakalinks FROM projects WHERE project_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		one = &ProjectModel{}
		if err := rows.Scan(&one.ID, &one.Titulo, &one.Detalle, &one.Descripcion, &one.Progreso, &one.Github, &one.Link, &one.Tiempo, pq.Array(&one.WakaLinks)); err != nil {
			fmt.Println("ERROR ID;", id, "NO ENCONTRADO", err)
			return nil, err
		}
		one.Tiempo = one.Tiempo.UTC()
	}
	return one, nil
}

func (query *dbStore) ReadAllProject() ([]*ProjectModel, error) {
	var all []*ProjectModel
	rows, err := query.db.Query("SELECT project_id, titulo, detalle, progreso FROM projects") //lenguajes
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		pr := &ProjectModel{}
		if err := rows.Scan(&pr.ID, &pr.Titulo, &pr.Detalle, &pr.Progreso); err != nil {
			fmt.Println(err)
			return nil, err
		}
		all = append(all, pr)
	}
	return all, nil
}

func (query *dbStore) UpdateProject(pr *ProjectModel) error {
	data, err := query.db.Query(`UPDATE projects SET titulo=$1, detalle=$2, descripcion=$3, progreso=$4, github=$5, link=$6, tiempo=$7, wakalinks=$8 WHERE project_id=$9;`,	pr.Titulo, pr.Detalle, pr.Descripcion, pr.Progreso, pr.Github, pr.Link, pr.Tiempo, pq.Array(pr.WakaLinks), pr.ID)
	defer data.Close()
	return err
}

func (query *dbStore) DeleteProject(pr *ProjectModel) error {
	data, err := query.db.Query(`DELETE FROM projects WHERE project_id=$1;`, pr.ID)
	defer data.Close()
	return err
}
