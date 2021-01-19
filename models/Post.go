package models

import (
	"database/sql"
	"fmt"
	"time"
)

//PostModel is the struct of a post
type PostModel struct {
	ID           int
	Titulo       string
	ProjectID    int
	ProjectTitle string // Obtenerlo en el query
	Fecha        time.Time
	Lenguajes    []*LenguageModel
	Replit       string
	Detalle      string
	Cuerpo       string
}

var (
	//Post is a variable to acces the methods
	Post PostModel
)

// CreatePost creates a new record in the database associating it wit the given projects_id
func (pt *PostModel) CreatePost() (err error) {
	err = Queries.CreatePost(pt)
	if err != nil {
		panic(err)
	}
	return nil
}

// ReadPost get all the post
func (pt *PostModel) ReadPost(id ...int) (one *PostModel, all []*PostModel, err error) {
	if len(id) != 0 && id[0] != 0 {
		// fmt.Println("ProjectID:", id[0])
		one, _, err = Queries.ReadPost(id[0])
	} else {
		_, all, err = Queries.ReadPost()
	}
	if err != nil {
		return nil, nil, err
	}
	return
}

// ReadProjectPost get all the post
func (pt *PostModel) ReadProjectPost(id int) (all []*PostModel) {
	all, err := Queries.ReadAllProjectPosts(id)
	if err != nil {
		panic(err)
	}
	return
}

// UpdatePost updates the given post
func (pt *PostModel) UpdatePost() {
	err := Queries.UpdatePost(pt)
	if err != nil {
		panic(err)
	}
	err = Queries.DeletePostRelationship(pt.ID)
	if err != nil {
		panic(err)
	}
	for _, leng := range pt.Lenguajes {
		err = Queries.CreatePostRelationship(leng, pt.ID)
		if err != nil {
			panic(err)
		}
	}
}

// DeletePost deletes the given post
func (pt *PostModel) DeletePost() {
	err := Queries.DeletePost(pt)
	if err != nil {
		panic(err)
	}
}

//HasProjectID is to identify if the note is about to be updated if the return is true or is to be created if the return is false
func (pt *PostModel) HasProjectID() bool {
	if pt.ProjectID == 0 {
		return false
	}
	return true
}

//===============================================================================================
//=========================================== QUERIES ===========================================
//===============================================================================================
//INSERT INTO posts(titulo, project_id,replit, detalle, cuerpo) VALUES ('Y',0,'$4','$5','$6')
func (query *dbStore) CreatePost(pt *PostModel) (err error) {
	if pt.ProjectID != 0 {
		_, err = query.db.Query(`INSERT INTO posts(titulo, project_id, fecha, replit, detalle, cuerpo) VALUES ($1,$2,$3,$4,$5,$6)`, pt.Titulo, pt.ProjectID, pt.Fecha, pt.Replit, pt.Detalle, pt.Cuerpo)
	} else {
		_, err = query.db.Query(`INSERT INTO posts(titulo, fecha, replit, detalle, cuerpo) VALUES ($1,$2,$3,$4,$5)`, pt.Titulo, pt.Fecha, pt.Replit, pt.Detalle, pt.Cuerpo)
	}
	if err != nil {
		return
	}
	rows, err := query.db.Query(`SELECT post_id FROM posts WHERE titulo = $1 AND replit = $2`, pt.Titulo, pt.Replit)
	if err != nil {
		return
	}
	defer rows.Close()
	var ID int
	for rows.Next() {
		if err = rows.Scan(&ID); err != nil {
			return
		}
	}
	for _, leng := range pt.Lenguajes {
		ln, err := Lenguage.ReadByName(leng.Titulo)
		if err != nil {
			panic(err)
		}
		if ln == nil {
			ln = &LenguageModel{Titulo: leng.Titulo}
			err = ln.CreateLenguage()
		}
		_, err = query.db.Query(`INSERT INTO post_leng(lenguaje_id, post_id) VALUES ($1,$2);`,
			ln.ID, ID)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (query *dbStore) ReadPost(id ...int) (one *PostModel, many []*PostModel, err error) {
	var idNull sql.NullInt64
	if len(id) > 0 {
		rows, err := query.db.Query("SELECT post_id, titulo, detalle, fecha, cuerpo, replit, project_id FROM posts WHERE post_id = $1 ", id[0]) //lenguajes
		if err != nil {
			return nil, nil, fmt.Errorf("Error: %v \nAt 129", err)
		}
		//SELECT pt.post_id, pt.titulo, pt.detalle, pt.fecha, pt.cuerpo, pt.replit FROM posts AS pt WHERE post_id = 2
		defer rows.Close()
		for rows.Next() {
			one = &PostModel{}
			if err := rows.Scan(&one.ID, &one.Titulo, &one.Detalle, &one.Fecha, &one.Cuerpo, &one.Replit, &idNull); err != nil {
				return nil, nil, fmt.Errorf("Error: %v \nAt 136", err)
			}
			if idNull.Valid {
				one.ProjectID = int(idNull.Int64)
				project, err := query.db.Query(`SELECT titulo FROM projects WHERE project_id = $1`, one.ProjectID)
				if err != nil {
					panic(err)
				}
				defer project.Close()
				for project.Next() {
					if err := project.Scan(&one.ProjectTitle); err != nil {
						return nil, nil, fmt.Errorf("Error: %v \nAt 146", err)
					}
				}
			}
			one.Lenguajes, err = Queries.ReadPostRelationship(one.ID)
			if err != nil {
				panic(err)
			}
		}
		if one == nil {
			return nil, nil, fmt.Errorf("Error, no post")
		}
	} else {
		rows, err := query.db.Query("SELECT pt.post_id, pt.titulo, pt.detalle, pt.fecha, pt.project_id FROM posts AS pt") //lenguajes
		if err != nil {
			return nil, nil, fmt.Errorf("Error: %v \nAt 158", err)
		}
		defer rows.Close()
		for rows.Next() {
			var post = &PostModel{}
			if err := rows.Scan(&post.ID, &post.Titulo, &post.Detalle, &post.Fecha, &idNull); err != nil {
				return nil, nil, fmt.Errorf("Error: %v \nAt 164", err)
			}
			if idNull.Valid {
				post.ProjectID = int(idNull.Int64)
				project, err := query.db.Query(`SELECT titulo FROM projects WHERE project_id = $1`, post.ProjectID)
				if err != nil {
					panic(err)
				}
				defer project.Close()
				for project.Next() {
					if err := project.Scan(&post.ProjectTitle); err != nil {
						return nil, nil, fmt.Errorf("Error: %v \nAt 146", err)
					}
				}
			}
			post.Lenguajes, err = Queries.ReadPostRelationship(post.ID)
			if err != nil {
				panic(err)
			}
			many = append(many, post)
		}
		if len(many) < 1 {
			return nil, nil, fmt.Errorf("Error, empty list")
		}
	}
	return
}

func (query *dbStore) ReadAllProjectPosts(pr int) (all []*PostModel, err error) {
	rows, err := query.db.Query(`SELECT * FROM posts WHERE project_id = $1`, pr)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		one := &PostModel{}
		if err = rows.Scan(&one.ID, &one.Titulo, &one.Fecha, &one.Detalle, &one.Cuerpo, &one.Replit, &one.ProjectID); err != nil {
			return nil, err
		}
		all = append(all, one)
	}
	return
}

func (query *dbStore) UpdatePost(pt *PostModel) error {
	var err error
	if pt.ProjectID != 0 {
		_, err = query.db.Query(`UPDATE posts SET titulo=$1, detalle=$2, cuerpo=$3, fecha=$4, replit=$5, project_id=$6 WHERE post_id=$7;`,
			pt.Titulo, pt.Detalle, pt.Cuerpo, pt.Fecha, pt.Replit, pt.ProjectID, pt.ID)
	} else {
		_, err = query.db.Query(`UPDATE posts SET titulo=$1, detalle=$2, cuerpo=$3, fecha=$4, replit=$5 WHERE post_id=$6;`,
			pt.Titulo, pt.Detalle, pt.Cuerpo, pt.Fecha, pt.Replit, pt.ID)
	}
	return err
}

//UPDATE posts SET titulo='$1', detalle='$2', cuerpo='$3', replit='$5' WHERE post_id=2
func (query *dbStore) DeletePost(pt *PostModel) (err error) {
	_, err = query.db.Query(`DELETE FROM posts WHERE post_id = $1`, pt.ID)
	return
}

//===============================================================================================
//========================================== Post_Leng ==========================================
//===============================================================================================

func (query *dbStore) CreatePostRelationship(ln *LenguageModel, pt int) (err error) {
	_, err = query.db.Query(`INSERT INTO post_leng(lenguaje_id, post_id) VALUES ($1,$2);`, ln.ID, pt)
	return
}

func (query *dbStore) ReadPostRelationship(post int) (lengs []*LenguageModel, err error) {
	rows, err := query.db.Query(` SELECT ln.lenguaje_id, ln.titulo FROM post_leng AS rl JOIN lenguajes AS ln ON rl.lenguaje_id = ln.lenguaje_id WHERE post_id = $1`, post)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		one := &LenguageModel{}
		if err = rows.Scan(&one.ID, &one.Titulo); err != nil {
			return nil, fmt.Errorf("Error: %v \nAt 221", err)
		}
		lengs = append(lengs, one)
	}
	return
}

func (query *dbStore) DeletePostRelationship(pt int) (err error) {
	_, err = query.db.Query(`DELETE FROM post_leng WHERE post_id=$1;`, pt)
	return
}
