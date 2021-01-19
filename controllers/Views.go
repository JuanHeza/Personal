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
		"Views/edit.html",
		"Views/Error.html",
		"Views/footer.html",
		"Views/header.html",
		"Views/home.html",
		"Views/login.html",
		"Views/post.html",
		"Views/postEditor.html",
		"Views/postList.html",
		"Views/project.html",
		"Views/projectForm.html",
		"Views/projectList.html",
		"Views/staticInfo.html",
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
