package controllers

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/JuanHeza/Personal/models"
	"github.com/gorilla/mux"
)

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/Create/Project", createProjectHandler).Methods("POST")
	router.HandleFunc("/Update/Project", updateProjectHandler).Methods("PUT")
	router.HandleFunc("/Delete/Project/{id}", deleteProjectHandler).Methods(http.MethodDelete)
	router.HandleFunc("/Edit/", func(w http.ResponseWriter, r *http.Request) { log.Println("Done") }) //WORKING
	return router
}

func NewModelForm() *url.Values {
	form := url.Values{
		"titulo":      []string{"Titulo_1"},
		"detalle":     []string{"Detalle_1"},
		"github":      []string{"Github_1"},
		"link":        []string{"Link_1"},
		"descripcion": []string{"Descrpcion_1"},
		"progreso":    []string{"99"},
		"lenguajes[]": []string{""},
		"title[]":     []string{"test"},
		"body[]":      []string{"test"},
		"date[]":      []string{"test"},
		"note_id[]":   []string{"test"},
		"caption[]":   []string{"test"},
		"gallery[]":   []string{"test"},
	}
	form.Set("titulo", "Test")
	form.Add("dato[]", "Dato 1")
	form.Add("dato[]", "Dato 2")
	form.Add("campo[]", "Campo 1")
	form.Add("campo[]", "Campo 2")
	return &form
}

func Test_ProjectHandlers(t *testing.T) {
	models.SetupDatabase(true)
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		id     string
		form   *url.Values
		url    string
		method string
		want   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valido",
			args: args{
				id:     "1",
				url:    "/Delete/Project/1",
				method: http.MethodDelete,
				want:   http.StatusFound,
			},
		}, {
			name: "valido",
			args: args{
				url:    "/Delete/Project/",
				method: http.MethodDelete,
				want:   http.StatusNotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.method, tt.args.url, nil) //bytes.NewBufferString(tt.args.prueba.form.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			if err != nil {
				t.Fatal(err)
			}
			recorder := httptest.NewRecorder()
			router := router()
			router.ServeHTTP(recorder, req)
			if status := recorder.Code; status != tt.args.want {
				t.Errorf("handler returned wrong status code, got %v want %v", status, tt.args.want)
			}
		})
	}
}

func Test_parseForm(t *testing.T) {
	// tiempo, _ := time.Parse("15:04:05", "07:41:32") //Now().Round(time.Second),
	fecha, _ := time.Parse("02/01/2006", "23/10/2020")
	tests := []struct {
		name    string
		form    map[string][]string //*url.Values
		files   []string
		wantPr  *models.ProjectModel
		wantErr bool
	}{
		{
			"Create",
			map[string][]string{ //&url.Values{
				"project_id":       {"0"},
				"detalle":          {"uno muy bonito"},
				"nombre":           {"Personal"},
				"link":             {"link"},
				"github":           {"github"},
				"progreso":         {"44"},
				"descripcion":      {"se supone que este va mas largo"},
				"lenguajes[]":      {"Go", "CSS", "HTML"},
				"wakalinks":        {"link_1", "link_2"},
				"note_title[]":     {"PROJECTO 1 NOTA 1", "PROJECTO 1 NOTA 2"},
				"note_body[]":      {"DETALLE 1.1", "DETALLE 1.2"},
				"note_date[]":      {"23-10-2020", "23-10-2020"},
				"note_id[]":        {"0", "0"},
				"update_title[]":   {"Update_1", "Update_2"},
				"update_caption[]": {"Caption_1", "Caption_2"},
				"update_index[]":   {"1", "2"},
				"caption[]":        {"DETALLE 1.1", "DETALLE 1.2"},
				//"icon"
				//"banner"
				// "gallery_titles[]": {"PROJECTO 1 IMAGEN 1", "PROJECTO 1 IMAGEN 2"},
			},
			[]string{"571202.jpg", "avatar.jpg"},
			&models.ProjectModel{
				// ID:          1,
				Titulo:      "Personal",
				Detalle:     "uno muy bonito",
				Descripcion: "se supone que este va mas largo",
				Progreso:    44,
				Github:      "github",
				Link:        "link",
				// Tiempo:      tiempo,
				WakaLinks: []string{"link_1", "link_2"},
				Notas: []*models.NoteModel{
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
				Galeria: []*models.ImageModel{
					{
						Titulo:  "Update_1",
						Detalle: "Caption_1",
						ID:      1,
					},
					{
						Titulo:  "Update_2",
						Detalle: "Caption_2",
						ID:      2,
					}, {
						Detalle: "DETALLE 1.1",
						Titulo:  "571202.jpg",
					}, {
						Detalle: "DETALLE 1.2",
						Titulo:  "avatar.jpg",
					},
				},
				Lenguajes: []*models.LenguageModel{
					{Titulo: "Go"},
					{Titulo: "CSS"},
					{Titulo: "HTML"},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requestBody bytes.Buffer
			multipartWriter := multipart.NewWriter(&requestBody)
			for _, val := range tt.files {
				file, err := os.Open(val)
				if err != nil {
					t.Errorf("INVALID FILE")
				}
				defer file.Close()
				fileWriter, err := multipartWriter.CreateFormFile("gallery[]", val)
				if err != nil {
					t.Errorf(err.Error())
				}
				_, err = io.Copy(fileWriter, file)
				if err != nil {
					t.Errorf(err.Error())
				}
			}
			for key, value := range tt.form {
				//fileWriter, err := multipartWriter.CreateFormField(key)
				//_, err = fileWriter.Write([]byte(value))
				for _, val := range value {
					_ = multipartWriter.WriteField(key, val)

				}
			}
			multipartWriter.Close()

			req, err := http.NewRequest(http.MethodPost, "tt.args.url", &requestBody) //bytes.NewBufferString(tt.args.prueba.form.Encode()))
			req.Header.Add("Content-Type", multipartWriter.FormDataContentType())

			gotPr, err := parseProjectForm(req)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseForm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := range gotPr.Lenguajes {
				if !reflect.DeepEqual(gotPr.Lenguajes[i], tt.wantPr.Lenguajes[i]) {
					t.Errorf("parseForm() = %v, want %v", gotPr.Lenguajes[i], tt.wantPr.Lenguajes[i])
				}
			}
			for j := range gotPr.Notas {
				if !reflect.DeepEqual(gotPr.Notas[j], tt.wantPr.Notas[j]) {
					t.Errorf("parseForm() = %v, want %v", gotPr.Notas[j], tt.wantPr.Notas[j])
				}
			}
			for k := range gotPr.Galeria {
				if !reflect.DeepEqual(gotPr.Galeria[k], tt.wantPr.Galeria[k]) {
					t.Errorf("parseForm() = %v, want %v", gotPr.Galeria[k], tt.wantPr.Galeria[k])
				}
			}
			if !reflect.DeepEqual(gotPr, tt.wantPr) {
				t.Errorf("parseForm() = %v, want %v", gotPr, tt.wantPr)
			}
		})
	}
}
