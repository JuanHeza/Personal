package main

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

//Store get and create data
type Store interface {
	CreateProyect(pr *Proyect) error
	CreateProject(pr *Projects) error
	DeleteProject(pr string) error
	GetProyect(id ...string) ([]*Projects, error)

	CreateFunction(fn *Function, pr string) error
	UpdateFunction(fn *Function) error
	DeleteFunction(fn *Function) error

	CreateModel(md *Model, pr string) error
	UpdateModel(md *Model, pr string) error
	DeleteModel(md *Model, pr string) error

	CreateTareas(tr *Task, pr string) error
	UpdateTareas(tr *Task) error
	DeleteTareas(tr *Task) error

	CreateNotas(nt *Note, pr string) error
	UpdateNotas(nt *Note) error
	DeleteNotas(nt *Note) error
}

//dbStore implements Store interface & use the connection object
type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateProyect(pr *Proyect) error {
	_, err := store.db.Query("INSERT INTO proyects(proyect,description) VALUES ($1,$2)", pr.Data, pr.Description)
	return err
}

func (store *dbStore) DeleteProject(pr string) error {
	_, err := store.db.Query("DELETE FROM informacion_general	WHERE nombre = $1", pr)
	return err
}

func (store *dbStore) CreateProject(pr *Projects) error {
	_, err := store.db.Query("INSERT INTO Informacion_General(	nombre, lenguajes, descripcion, introduccion, progreso, icon, banner) VALUES ($1,$2,$3,$4,$5,$6,$7);", pr.Name, pq.Array(pr.Language), pr.Description, pr.Introduccion, pr.Progress, pr.Icon, pr.Banner)
	return err
}

func (store *dbStore) GetProyect(id ...string) ([]*Projects, error) {
	var rows *sql.Rows
	var err error
	fmt.Println(id)
	if len(id) > 0 {
		rows, err = store.db.Query(" select  * from informacion_general where nombre=$1", id[0])
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		proyects := []*Projects{}
		for rows.Next() {
			proyect := &Projects{}
			if err := rows.Scan(&proyect.Name, &proyect.Description, &proyect.Icon, &proyect.Banner, &proyect.Progress, &proyect.Introduccion, pq.Array(&proyect.Language)); err != nil {
				fmt.Println("[ROWS] LINEA 52", err)
				return nil, err
			}

			functions, err := store.db.Query(" select  id_funcion, llamada, regresa, detalle, codigo from funciones where proyecto=$1", id[0])
			if err != nil {
				return nil, err
			}
			defer functions.Close()
			funcion := Function{}
			for functions.Next() {
				if err := functions.Scan(&funcion.ID, &funcion.Call, &funcion.Return, &funcion.Description, &funcion.Codigo); err != nil {
					fmt.Println("[FUNCIONES] LINEA 64", err)
					return nil, err
				}
				proyect.Functions = append(proyect.Functions, funcion)
			}

			notes, err := store.db.Query(" select  id_nota, titulo ,cuerpo from notas where proyecto=$1", id[0])
			if err != nil {
				return nil, err
			}
			defer notes.Close()
			nota := Note{}
			for notes.Next() {
				if err := notes.Scan(&nota.ID, &nota.Title, &nota.Text); err != nil {
					fmt.Println("[NOTAS] LINEA 78", err)
					return nil, err
				}
				proyect.Notes = append(proyect.Notes, nota)
			}

			tasks, err := store.db.Query(" select id_tarea, situacion, tarea from tareas where proyecto=$1", id[0])
			if err != nil {
				return nil, err
			}
			defer tasks.Close()
			tarea := Task{}
			for tasks.Next() {
				if err := tasks.Scan(&tarea.ID, &tarea.Done, &tarea.Text); err != nil {
					fmt.Println("[TAREAS] LINEA 92", err)
					return nil, err
				}
				proyect.Tasks = append(proyect.Tasks, tarea)
			}

			models, err := store.db.Query("select modelo, campos, datos from modelos where proyecto=$1", id[0])
			if err != nil {
				return nil, err
			}
			defer models.Close()
			for models.Next() {
				model := &Model{}
				data := new([2][]string)
				if err := models.Scan(&model.Title, pq.Array(&data[0]), pq.Array(&data[1])); err != nil {
					fmt.Println("[MODELOS] LINEA 107", err)
					return nil, err
				}
				for x, y := range data[0] {
					val := &Data{Name: y, DataType: data[1][x]}
					model.Data = append(model.Data, *val)
				}
				proyect.Models = append(proyect.Models, *model)
			}
			// fmt.Println(*proyect)
			proyects = append(proyects, proyect)
		}
		return proyects, nil
	}
	rows, err = store.db.Query("SELECT nombre, descripcion, icon, banner, progreso, introduccion, lenguajes FROM informacion_general")
	//rows, err := store.db.Query("SELECT proyect, description from proyects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	proyects := []*Projects{}
	for rows.Next() {
		proyect := &Projects{}
		if err := rows.Scan(&proyect.Name, &proyect.Description, &proyect.Icon, &proyect.Banner, &proyect.Progress, &proyect.Introduccion, pq.Array(&proyect.Language)); err != nil {
			fmt.Println(err)
			return nil, err
		}
		proyects = append(proyects, proyect)
	}
	return proyects, nil
}

var store Store

//InitStore method
func InitStore(s Store) {
	store = s
}

func (store *dbStore) CreateFunction(fn *Function, pr string) error {
	_, err := store.db.Query("INSERT INTO funciones(proyecto, llamada, regresa, descripcion, codigo)	VALUES ($1,$2,$3,$4,$5)", pr, fn.Call, fn.Return, fn.Description, fn.Codigo)
	return err
}
func (store *dbStore) UpdateFunction(fn *Function) error {
	_, err := store.db.Query("UPDATE funciones SET  llamada=$1, regresa=$2, descripcion=$3, codigo=$4 WHERE id=%5;", fn.Call, fn.Return, fn.Description, fn.Codigo, fn.ID)
	return err
}
func (store *dbStore) DeleteFunction(fn *Function) error {
	_, err := store.db.Query("DELETE FROM funciones WHERE id=$1;", fn.ID)
	return err
}

func modelToArray(md *Model) ([]string, []string) {
	var campos, datos []string
	for _, y := range md.Data {
		campos = append(campos, y.Name)
		datos = append(datos, y.DataType)
	}
	return campos, datos
}

func (store *dbStore) CreateModel(md *Model, pr string) error {
	campos, datos := modelToArray(md)
	_, err := store.db.Query("INSERT INTO modelos(proyecto, modelo, campos, datos)	VALUES ($1,$2,$3,$4)", pr, md.Title, pq.Array(campos), pq.Array(datos))
	return err
}
func (store *dbStore) UpdateModel(md *Model, pr string) error {
	campos, datos := modelToArray(md)
	_, err := store.db.Query("UPDATE modelos SET  modelo=$1, campos=$2, datos=$3 WHERE modelo=$4 and proyecto=$5;", md.Title, pq.Array(campos), pq.Array(datos), md.Title, pr)
	return err
}
func (store *dbStore) DeleteModel(md *Model, pr string) error {
	_, err := store.db.Query("DELETE FROM modelos WHERE modelo=$1 and proyecto=$2;", md.Title, pr)
	return err
}

func (store *dbStore) CreateTareas(tr *Task, pr string) error {
	_, err := store.db.Query("INSERT INTO tareas(proyecto, tarea, situacion)	VALUES ($1, $2, $3)", pr, tr.Text, tr.Done)
	return err
}
func (store *dbStore) UpdateTareas(tr *Task) error {
	_, err := store.db.Query("UPDATE tareas SET tarea=$1, situacion=$2 WHERE id=$3;", tr.Text, tr.Done, tr.ID)
	return err
}
func (store *dbStore) DeleteTareas(tr *Task) error {
	_, err := store.db.Query("DELETE FROM tareas WHERE id=$1", tr.ID)
	return err
}

func (store *dbStore) CreateNotas(nt *Note, pr string) error {
	_, err := store.db.Query("INSERT INTO notas(titulo, cuerpo, proyecto) VALUES ($1,$2,$3)", nt.Title, nt.Text, pr)
	return err
}
func (store *dbStore) UpdateNotas(nt *Note) error {
	_, err := store.db.Query("UPDATE notas SET cuerpo=$1, titulo=$2, WHERE id=$3;", nt.Text, nt.Title, nt.ID)
	return err
}
func (store *dbStore) DeleteNotas(nt *Note) error {
	_, err := store.db.Query("DELETE FROM notas WHERE id=$1;", nt.ID)
	return err
}

/*
 select  i.*, f.llamada, f.regresa, f.detalle, f.codigo,n.titulo, n.cuerpo, t.situacion, t.tarea
  from informacion_general as i
  join funciones as f
  on nombre = f.proyecto
  join notas as n
  on nombre = n.proyecto
	join tareas as t
	on nombre = t.proyecto;

select modelo, campos, datos from modelos where proyecto='Personal';

insert into modelos(proyecto, modelo, campos, datos) values
('Personal','Project','{"Name","Language","Desription","Icon","Time","Progress","Model","Side"}','{"string","string","string","string","string","float32","Models Array","int"}'),
('Personal','Models','{"Title","Data"}','{"string","Data Array"}'),
('Personal','Data','{"Name","DataType","Description"}','{"string","string", "string"}')
;
*/
