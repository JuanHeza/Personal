package models

import "fmt"

//https://www.postgresqltutorial.com/

var (
	schema = []struct {
		table string
		value string
	}{
		{
			"users",
			`user_id INT GENERATED ALWAYS AS IDENTITY,
			username TEXT,
			password TEXT,
			admin BOOLEAN,
			PRIMARY KEY(user_id)`,
		}, // USERS
		{
			"projects",
			`project_id INT GENERATED ALWAYS AS IDENTITY,
			titulo TEXT,
			detalle TEXT,
			descripcion TEXT,
			status TEXT,
			github TEXT,
			link TEXT,
			tiempo timestamp, 
			label TEXT,
			updated timestamp, 
			PRIMARY KEY(project_id)`,
		}, // PROJECTS
		{
			"lenguajes",
			`lenguaje_id INT GENERATED ALWAYS AS IDENTITY,
			titulo TEXT,
			PRIMARY KEY(lenguaje_id)`,
		}, // LENGUAJE
		{
			"proj_leng",
			`lenguaje_id INT NOT NULL,
			project_id INT NOT NULL,
			PRIMARY KEY (project_id, lenguaje_id),
			FOREIGN KEY (lenguaje_id)
				REFERENCES lenguajes (lenguaje_id)
				ON DELETE CASCADE,
			FOREIGN KEY (project_id)
				REFERENCES projects (project_id)
				ON DELETE CASCADE`,
		}, // PROG_LENG
		{
			"notes",
			`note_id INT GENERATED ALWAYS AS IDENTITY,
			titulo TEXT,
			fecha timestamp without time zone,
			detalle TEXT,
			project_id INT NOT NULL,
			PRIMARY KEY(note_id),
			FOREIGN KEY (project_id)
			REFERENCES projects (project_id)
			ON DELETE CASCADE`,
		}, // NOTES
		{
			"images",
			`image_id INT GENERATED ALWAYS AS IDENTITY,
			titulo TEXT,
			detalle TEXT,
			project_id INT NOT NULL,
			PRIMARY KEY(image_id),
			FOREIGN KEY (project_id)
				REFERENCES projects (project_id)
				ON DELETE CASCADE`,
		}, // IMAGES
		{
			"statics",
			`introduccion TEXT,
			about TEXT,
			tutorial TEXT,
			contacto TEXT,
			leng TEXT`, /* array de [link, link image]*/
		}, // STATICS
		{
			"links",
			`link_id INT GENERATED ALWAYS AS IDENTITY,
			link TEXT,
			icon TEXT`,
		}, // LINKS
		{
			"posts",
			`post_id INT GENERATED ALWAYS AS IDENTITY,
			titulo TEXT,
			fecha timestamp without time zone,
			detalle TEXT,
			cuerpo TEXT,
			replit TEXT,
			project_id INT,
			PRIMARY KEY(post_id),
			FOREIGN KEY (project_id)
			REFERENCES projects (project_id)
			ON DELETE SET NULL`,
		}, // POST
		{
			"post_leng",
			`lenguaje_id INT NOT NULL,
			post_id INT NOT NULL,
			PRIMARY KEY (post_id, lenguaje_id),
			FOREIGN KEY (lenguaje_id)
				REFERENCES lenguajes (lenguaje_id)
				ON DELETE CASCADE,
			FOREIGN KEY (post_id)
				REFERENCES posts (post_id)
				ON DELETE CASCADE`,
		}, // POST_LENG
	}
)

//SetupDatabase drop tables if exist and creates a new one
func SetupDatabase(seed bool) {
	var host, port, user, password, dbname = "localhost", "5432", "postgres", "JHZ697heza", "ProycetsDB"
	StartConnection(host, port, user, password, dbname)
	db := getDatabase()
	err := db.DropTables()
	if err != nil {
		panic(err)
	}
	err = db.CreateTables()
	if err != nil {
		panic(err)
	}
	if seed {

		SeedLenguages()
		SeedProjects()
		SeedPost()
	}
}

//SetupDatabaseDevelopment is something
func SetupDatabaseDevelopment() {
	db := getDatabase()
	err := db.DropTables()
	if err != nil {
		panic(err)
	}
	err = db.CreateTables()
	if err != nil {
		panic(err)
	}
	SeedStatics()
	SeedProjects()
	SeedPost()
}

//DropTables drop all tables in the shecma
func (store *dbStore) DropTables() error {
	for _, data := range schema {
		query := fmt.Sprintf("DROP TABLE IF EXISTS %v CASCADE;", data.table)
		rows, err := store.db.Query(query)
		defer rows.Close()
		if err != nil {
			return err
		}
	}
	fmt.Println("TABLES DROPED ")
	return nil
}

//CreateTables take the schema and create the tables of the database
func (store *dbStore) CreateTables() error {
	for _, data := range schema {
		query := fmt.Sprintf("CREATE TABLE %v(%v);", data.table, data.value)
		rows, err := store.db.Query(query)
		defer rows.Close()
		if err != nil {
			return err
		}
	}
	fmt.Println("TABLES CREATED ")
	return nil
}
