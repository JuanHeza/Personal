package models

// import "fmt"

var ()

//LenguageModel is a representation of the data in the database
type LenguageModel struct {
	ID     int
	Titulo string
}

var (
	//Lenguage is an instance of the model to acces to the methods
	Lenguage LenguageModel
)

//CreateLenguage a new record in the database
func (ln *LenguageModel) CreateLenguage() error {
	err := Queries.CreateLenguage(ln)
	if err != nil {
		panic(err)
	}
	// fmt.Println(im)
	return nil
}

//ReadByName the title and id of the given name
func (ln *LenguageModel) ReadByName(name string) (*LenguageModel, error) {
	var one *LenguageModel
	one, err := Queries.ReadLenguageTitle(name)
	if err != nil {
		panic(err)
	}
	return one, nil
}

//ReadByID the title and id of the given Id
func (ln *LenguageModel) ReadByID(id int) (*LenguageModel, error) {
	var one *LenguageModel
	one, err := Queries.ReadLenguageID(id)
	if err != nil {
		panic(err)
	}
	return one, nil
}

// ReadAll the lenguages of the database or just the relations if a project is given
func (ln *LenguageModel) ReadAll() []*LenguageModel {
	var all []*LenguageModel
	var err error
	all, err = Queries.ReadAllLenguages()
	if err != nil {
		panic(err)
	}
	return all
}

// ReadAllByLenguage the lenguages of the database or just the relations if a project is given
func (ln *LenguageModel) ReadAllByLenguage(leng int) map[int]*ProjectModel {
	var all map[int]*ProjectModel
	var err error
	all, err = Queries.ReadLenguageRelationship(leng)
	if err != nil {
		panic(err)
	}
	return all
}

func createRelationship(lengs []*LenguageModel, proj int) {
	for _, leng := range lengs {
		ln, err := Lenguage.ReadByName(leng.Titulo)
		if err != nil {
			panic(err)
		}
		if ln == nil {
			ln = &LenguageModel{Titulo: leng.Titulo}
			err = ln.CreateLenguage()
		}
		Queries.CreateRelationship(proj, ln.ID)
	}
}

func readRelationships(proj int) []*LenguageModel {
	IDs, err := Queries.ReadRelationship(proj)
	if err != nil {
		panic(err)
	}
	var all []*LenguageModel
	for _, ID := range IDs {
		one, err := Lenguage.ReadByID(ID)
		if err != nil {
			panic(err)
		}
		all = append(all, one)
	}
	return all
}

func updateRelationships(lengs []*LenguageModel, proj int) {
	err := Queries.DeleteRelationship(proj)
	if err != nil {
		panic(err)
	}
	createRelationship(lengs, proj)
}

//===============================================================================================
//=========================================== QUERIES ===========================================
//===============================================================================================

func (query *dbStore) CreateLenguage(ln *LenguageModel) error {
	var ID int
	_, err := query.db.Query(`INSERT INTO lenguajes(titulo) VALUES ($1);`,
		ln.Titulo)
	rows, err := query.db.Query("SELECT lenguaje_id FROM lenguajes WHERE titulo=$1", ln.Titulo)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&ID); err != nil {
			return err
		}
	}
	ln.ID = ID
	return nil
}

func (query *dbStore) ReadLenguageTitle(titulo string) (*LenguageModel, error) {
	var ln *LenguageModel
	rows, err := query.db.Query("SELECT lenguaje_id FROM lenguajes WHERE titulo=$1", titulo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		ln = &LenguageModel{}
		if err := rows.Scan(&ln.ID); err != nil {
			return nil, err
		}
	}
	return ln, nil
}

func (query *dbStore) ReadLenguageID(id int) (*LenguageModel, error) {
	var ln *LenguageModel
	rows, err := query.db.Query("SELECT lenguaje_id,titulo FROM lenguajes WHERE lenguaje_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		ln = &LenguageModel{}
		if err := rows.Scan(&ln.ID, &ln.Titulo); err != nil {
			return nil, err
		}
	}
	return ln, nil
}

func (query *dbStore) ReadAllLenguages() ([]*LenguageModel, error) {
	var one *LenguageModel
	var all []*LenguageModel
	rows, err := query.db.Query("SELECT lenguaje_id, titulo FROM lenguajes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		one = &LenguageModel{}
		if err := rows.Scan(&one.ID, &one.Titulo); err != nil {
			return nil, err
		}
		all = append(all, one)
	}
	return all, nil
}

// func (query *dbStore) UpdateLenguage(ln *LenguageModel) error {
// 	_, err := query.db.Query(`UPDATE lenguajes SET titulo=$1 WHERE lenguaje_id=$2;`, ln.Titulo, ln.ID)
// 	return err
// }

// func (query *dbStore) DeleteLenguage(ln *LenguageModel) error {
// 	_, err := query.db.Query(`DELETE FROM lenguajes WHERE lenguaje_id=$1;`, ln.ID)
// 	return err
// }

//===============================================================================================
//========================================== Prog_Leng ==========================================
//===============================================================================================

func (query *dbStore) CreateRelationship(proj, leng int) error {
	_, err := query.db.Query(`INSERT INTO proj_leng(lenguaje_id, project_id) VALUES ($1,$2);`,
		leng, proj)
	return err
}

func (query *dbStore) ReadRelationship(proj int) ([]int, error) {
	var leng []int
	rows, err := query.db.Query("SELECT lenguaje_id FROM proj_leng WHERE project_id=$1", proj)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ID int
		if err := rows.Scan(&ID); err != nil {
			return nil, err
		}
		leng = append(leng, ID)
	}
	return leng, nil
}

func (query *dbStore) ReadLenguageRelationship(leng int) (map[int]*ProjectModel, error) {
	var list = make(map[int]*ProjectModel)
	rows, err := query.db.Query("SELECT pr.project_id, pr.titulo, pr.detalle, lengs.titulo, lengs.lenguaje_id  FROM proj_leng AS rl JOIN lenguajes AS ln ON rl.lenguaje_id = ln.lenguaje_id JOIN projects AS pr ON rl.project_id = pr.project_id JOIN proj_leng AS rels ON rels.project_id = pr.project_id JOIN lenguajes AS lengs ON lengs.lenguaje_id = rels.lenguaje_id WHERE ln.lenguaje_id = $1", leng)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var pr = &ProjectModel{}
		var ln = &LenguageModel{}
		if err := rows.Scan(&pr.ID, &pr.Titulo, &pr.Detalle, &ln.Titulo, &ln.ID); err != nil {
			return nil, err
		}
		if _, ok := list[pr.ID]; ok == false {
			list[pr.ID] = pr
		}
		list[pr.ID].Lenguajes = append(list[pr.ID].Lenguajes, ln)
	}
	return list, nil
}

func (query *dbStore) DeleteRelationship(proj int) error {
	_, err := query.db.Query(`DELETE FROM proj_leng WHERE project_id=$1;`, proj)
	return err
}
