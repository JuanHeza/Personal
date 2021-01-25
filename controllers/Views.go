package controllers

import (
	"html/template"
	"time"

	"github.com/JuanHeza/Personal/models"
)

var (
	//Templates variable with all the templates data
	Templates *template.Template
	files     = []string{
		"views/edit.html",
		"views/error.html",
		"views/footer.html",
		"views/header.html",
		"views/home.html",
		"views/logIn.html",
		"views/post.html",
		"views/postEditor.html",
		"views/postList.html",
		"views/project.html",
		"views/projectForm.html",
		"views/projectList.html",
		"views/staticInfo.html",
	}
)

//InitTemplates initialize the variable with all the templates
func InitTemplates() {
	TemplateFunctions := template.FuncMap{
		"Links":                Links,
		"StaticData":           GetStatics,
		"mod":                  func(i int) bool { return i%2 == 0 },
		"fecha":                Fecha,
		"IsAdminAutenticathed": IsAdminAutenticathed,
		"Lenguajes":            models.Lenguage.ReadAll,
		"today":                func() string { return Fecha(time.Now(), true) },
		"HTML":                 func(in string) template.HTML { return template.HTML(in) },
		"Horas":                Horas,
	}
	Templates = template.New("")
	Templates = template.Must(Templates.Funcs(TemplateFunctions).ParseFiles(files...))
}
