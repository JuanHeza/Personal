package models

import "time"

var (
	tiempo, _ = time.Parse("15:04:05", "07:41:32")     //Now().Round(time.Second),
	fecha, _  = time.Parse("02/01/2006", "23/10/2020") //Now().Round(time.Second),
	projects  = []*ProjectModel{
		{
			// ID:          1,
			Titulo:      "Personal",
			Detalle:     "uno muy bonito",
			Descripcion: "se supone que este va mas largo",
			Status:      "Proceso",
			Github:      "ttps://github.com/JuanHeza",
			Link:        "ttps://github.com/JuanHeza",
			// Tiempo:      tiempo,
			Label: "Personal",
			Notas: []*NoteModel{
				{
					Titulo:  "PROJECTO 1 NOTA 1",
					Fecha:   fecha,
					Detalle: "DETALLE 1.1",
				},
				{
					Titulo:  "PROJECTO 1 NOTA 2",
					Fecha:   fecha,
					Detalle: "DETALLE 1.2",
				},
			},
			Galeria: []*ImageModel{
				{
					Titulo:  "img (3).png",
					Detalle: "DETALLE 1.1",
				},
				{
					Titulo:  "img (4).png",
					Detalle: "DETALLE 1.2",
				},
			},
			Lenguajes: []*LenguageModel{
				{Titulo: "Go"},
				{Titulo: "CSS"},
				{Titulo: "HTML"},
			},
		},
		{
			// ID:          2,
			Titulo:      "Artemis",
			Detalle:     "mi favorito",
			Descripcion: "hi patito feo",
			Status:      "Proceso",
			Github:      "https://github.com/JuanHeza",
			Link:        "https://github.com/JuanHeza",
			// Tiempo:      tiempo,
			Label: "Personal",
			Notas: []*NoteModel{
				{
					Titulo:  "PROJECTO 2 NOTA 1",
					Fecha:   fecha,
					Detalle: "DETALLE 2.1",
				},
			},
			Galeria: []*ImageModel{
				{
					Titulo:  "PROJECTO 2 IMAGEN 1",
					Detalle: "DETALLE 2.1",
				},
			},
		},
	}
	notes = []*NoteModel{
		{
			ID:        1,
			ProjectID: 1,
			Titulo:    "PROJECTO 1 NOTA 1",
			Fecha:     fecha,
			Detalle:   "DETALLE 1.1",
		},
		{
			ID:        2,
			ProjectID: 1,
			Titulo:    "PROJECTO 1 NOTA 2",
			Fecha:     fecha,
			Detalle:   "DETALLE 1.2",
		}, {
			// ID:        1,
			// ProjectID: 1,
			Titulo:  "UPDATE PROJECTO 1 NOTA 1",
			Fecha:   fecha,
			Detalle: "UPDATE DETALLE 1.1",
		},
		{
			// ID:        2,
			// ProjectID: 1,
			Titulo:  "UPDATE PROJECTO 1 NOTA 2",
			Fecha:   fecha,
			Detalle: "UPDATE DETALLE 1.2",
		},
	}
	images = []*ImageModel{
		{
			ID:        1,
			ProjectID: 1,
			Titulo:    "PROJECTO 1 IMAGEN 1",
			Detalle:   "DETALLE 1.1",
		},
		{
			ID:        2,
			ProjectID: 1,
			Titulo:    "PROJECTO 1 IMAGEN 2",
			Detalle:   "DETALLE 1.2",
		},
		{
			// ID:        1,
			// ProjectID: 1,
			Titulo:  "UPDATE |-| PROJECTO 1 IMAGEN 1",
			Detalle: "UPDATE |-| DETALLE 1.1",
		},
		{
			// ID:        2,
			// ProjectID: 1,
			Titulo:  "UPDATE |-| PROJECTO 1 IMAGEN 2",
			Detalle: "UPDATE |-| DETALLE 1.2",
		},
	}
	lenguajes = []*LenguageModel{
		{
			ID:     1,
			Titulo: "Go",
		}, {
			ID:     2,
			Titulo: "CSS",
		}, {
			ID:     3,
			Titulo: "HTML",
		}, {
			Titulo: "C#",
		}, {
			Titulo: "Arduino",
		}, {
			Titulo: "Ruby",
		},
	}
	statics = []*StaticData{
		{
			Introduccion: "Este sitio funciona como portafolio donde se pueden encontrar los proyectos que tengo completados, en desarrollo y en planeacion, principalmente de software pero tambien se pueden encontrar algo de electronica, IoT y algunas cosas mas",
			About:        "Juan Enrique un desarrollador de software mexicano, siendo 'Heza' una abreviacion de sus apellidos, con gusto por los retos ya que es una forma muy interesante de aprender cosas nuevas y usarlas en el momento, esta especializado en el lenguaje Go y enfocado en aprender el desarrollo web en el ambito de Backend",
			Tutorial:     "Aqui puedes encontrar la lista completa de los proyectos, asi como tambien posts de como se hicieron algunas cosas en los diferentes proyectos.',",
			Contacto:     "Informacion de contacto",
			Leng:         "es",
		},
		{
			Introduccion: "This site works as a portafolio where you can found the proyects i have completed, in development or planning, mainly of software but also you can found some of electronic, IoT and some other things",
			About:        "Juan Enrique is a mexican software developer,'Heza' its just an abbreviation of his last names, he likes the challenges since it`s a very interesting way to learn new things and use them at the moment, he is specialized in the Go language and focused on learning web development in the field of Backend",
			Tutorial:     "You can found a complete list of proyects and posts of how to do some things i have use in the proyects",
			Contacto:     "Contact Information",
			Leng:         "en",
		},
	}
	links = []*LinkModel{
		{Link: "mailto:juanehza@hotmail.com", Icon: "envelope.png", ID: 1},
		{Link: "https://repl.it/@JuanHeza/", Icon: "Repl.it.png", ID: 2},
		{Link: "https://github.com/JuanHeza", Icon: "GitHub-Mark-32px.png", ID: 3},
		{Link: "/api/Home/", Icon: "GitHub-Mark-32px.png", ID: 4},
	}
	posts = []*PostModel{
		{
			Titulo:    "Project Post",
			ProjectID: 1,
			Fecha:     fecha,
			Replit:    "#",
			Detalle:   "DETALLE DEL POST",
			Cuerpo:    "AQUI DEBES PONER EL CUERPO EXPICANDO TODA LA CACERIA",
			Lenguajes: []*LenguageModel{
				{Titulo: "Go"},
			},
		},
		{
			Titulo:    "Nil Post",
			ProjectID: 0,
			Fecha:     fecha,
			Replit:    "#",
			Detalle:   "DETALLE DEL POST",
			Cuerpo:    "AQUI DEBES PONER EL CUERPO EXPICANDO TODA LA CACERIA",
			Lenguajes: []*LenguageModel{
				{Titulo: "Ruby"},
			},
		},
	}
)

//SeedProjects fill the database with the given projects
func SeedProjects() {
	for _, value := range projects {
		value.CreateProject()
	}
}

// SeedPost initialize some posts to the database
func SeedPost() {
	for _, post := range posts {
		post.CreatePost()
	}
}

//SeedLenguages fill the database with the given Lenguages
func SeedLenguages() {
	for _, value := range lenguajes {
		value.CreateLenguage()
	}
}

// SeedStatics fills the database with the static data
func SeedStatics() {
	for _, value := range statics {
		value.CreateStatics()
	}
	for _, link := range links {
		link.CreateLink()
	}
	SeedLenguages()
}
