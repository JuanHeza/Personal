package main

import (
	"bytes"
	_ "encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

type prueba struct {
	id      string
	form    *url.Values
	spected interface{}
}

type pruebaF struct {
	id      string
	form    *url.Values
	spected *Function
}

type pruebaN struct {
	id      string
	form    *url.Values
	spected *Note
}

type pruebaT struct {
	id      string
	form    *url.Values
	spected *Task
}

type variable struct {
	routeVariable string
	shouldPass    bool
}

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/Crud/Funcion", createFunctionHandler).Methods("POST")
	router.HandleFunc("/Crud/Funcion", updateFunctionHandler).Methods("PUT")
	router.HandleFunc("/Crud/Funcion/{id}", deleteFunctionHandler).Methods("DELETE")

	router.HandleFunc("/Crud/Modelo", createModelHandler).Methods("POST")
	router.HandleFunc("/Crud/Modelo", updateModelHandler).Methods("PUT")
	router.HandleFunc("/Crud/Modelo/{id}", deleteModelHandler).Methods("DELETE")

	router.HandleFunc("/Crud/Notas", createNotasHandler).Methods("POST")
	router.HandleFunc("/Crud/Notas", updateNotasHandler).Methods("PUT")
	router.HandleFunc("/Crud/Notas/{id}", deleteNotasHandler).Methods("DELETE")

	router.HandleFunc("/Crud/Tarea", createTareasHandler).Methods("POST")
	router.HandleFunc("/Crud/Tarea", updateTareasHandler).Methods("PUT")
	router.HandleFunc("/Crud/Tarea/{id}", deleteTareasHandler).Methods("DELETE")
	return router
}

////////////////////////////////////////////////////////////////////////////////////////////

func NewModelForm() *url.Values {
	form := url.Values{}
	form.Set("titulo", "Test")
	form.Add("dato[]", "Dato 1")
	form.Add("dato[]", "Dato 2")
	form.Add("campo[]", "Campo 1")
	form.Add("campo[]", "Campo 2")
	return &form
}

func Test_updateModelHandler(t *testing.T) { //PASSED
	mockStore := InitMockStore()
	testData := prueba{}

	mockStore.On("UpdateModel", &Model{Title: "Test", Data: []Data{Data{Name: "Campo 1", DataType: "Dato 1"}, Data{Name: "Campo 2", DataType: "Dato 2"}}}, "Personal").Return(nil).Once()
	testData.form = NewModelForm()
	testData.id = "Personal"
	req, err := http.NewRequest("PUT", "/Crud/"+testData.id+"/Modelo", bytes.NewBufferString(testData.form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router := router()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v want %v", status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}
func Test_createModelHandler(t *testing.T) { //PASSED
	mockStore := InitMockStore()
	testData := prueba{}

	mockStore.On("CreateModel", &Model{Title: "Test", Data: []Data{Data{Name: "Campo 1", DataType: "Dato 1"}, Data{Name: "Campo 2", DataType: "Dato 2"}}}, "Personal").Return(nil).Once()
	testData.form = NewModelForm()
	testData.id = "Personal"
	req, err := http.NewRequest("POST", "/Crud/"+testData.id+"/Modelo", bytes.NewBufferString(testData.form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router := router()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v want %v", status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}
func Test_deleteModelHandler(t *testing.T) { //PASSED
	mockStore := InitMockStore()
	// testData := prueba{id: "personal"}

	mockStore.On("DeleteModel", &Model{ID: 6}).Return(nil).Once()
	req, err := http.NewRequest("DELETE", "/Crud/Modelo/6", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	router := router()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v want %v", status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}

////////////////////////////////////////////////////////////////////////////////////////////

func newFunctionForm() *url.Values {
	form := url.Values{}
	form.Set("llamada", "Prueba")
	form.Set("return", "")
	form.Set("funcion", "Datos de Prueba")
	form.Set("codigo", "Nada")
	return &form
}

var function = pruebaF{
	id:      "personal",
	spected: &Function{ID: 0, Call: "Prueba", Return: "", Description: "Datos de Prueba", Codigo: "Nada"},
	form:    newFunctionForm(),
}

func Test_createFunctionHandler(t *testing.T) { //PASSED
	mockStore := InitMockStore()
	testData := function

	mockStore.On("CreateFunction", testData.spected, testData.id).Return(nil).Once()
	req, err := http.NewRequest("POST", "/Crud/"+testData.id+"/Funcion", bytes.NewBufferString(testData.form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	router := router()
	router.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v want %v", status, http.StatusFound)
	}
	mockStore.AssertExpectations(t)
}

func Test_deleteFunctionHandler(t *testing.T) { //PASSED
	mockStore := InitMockStore()
	testData := pruebaF{
		// id:      "personal",
		spected: &Function{ID: 6},
	}
	mockStore.On("DeleteFunction", testData.spected).Return(nil).Once()
	req, err := http.NewRequest("DELETE", "/Crud/Funcion/"+strconv.Itoa(testData.spected.ID), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	router := router()
	router.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v want %v", status, http.StatusFound)
	}
	mockStore.AssertExpectations(t)
}

func Test_updateFunctionHandler(t *testing.T) { //PASSED
	mockStore := InitMockStore()
	testData := function
	mockStore.On("UpdateFunction", testData.spected, testData.id).Return(nil).Once()
	req, err := http.NewRequest("PUT", "/Crud/"+testData.id+"/Funcion", bytes.NewBufferString(testData.form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	router := router()
	router.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v wanted %v", status, http.StatusFound)
	}
	mockStore.AssertExpectations(t)
}

////////////////////////////////////////////////////////////////////////////////////////////

func newNoteForm() *url.Values {
	form := url.Values{}
	form.Set("nota", "Prueba")
	form.Set("cuerpo", "Nada")
	return &form
}

var nota = pruebaN{
	id:      "Personal",
	spected: &Note{ID: 0, Title: "Prueba", Text: "Nada"},
	form:    newNoteForm(),
}

func Test_createNotasHandler(t *testing.T) { //PASSED
	mockStore := InitMockStore()
	testData := nota

	mockStore.On("CreateNotas", testData.spected, testData.id).Return(nil).Once()
	req, err := http.NewRequest("POST", "/Crud/"+testData.id+"/Notas", bytes.NewBufferString(testData.form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	router := router()
	router.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v want %v", status, http.StatusFound)
	}
	mockStore.AssertExpectations(t)
}

func Test_deleteNotasHandler(t *testing.T) { //PASSED
	mockStore := InitMockStore()
	testData := pruebaN{
		id:      "Personal",
		spected: &Note{ID: 0, Title: "", Text: ""},
	}
	mockStore.On("DeleteNotas", testData.spected, testData.id).Return(nil).Once()
	req, err := http.NewRequest("DELETE", "/Crud/"+(testData.id)+"/Notas/"+strconv.Itoa(testData.spected.ID), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	router := router()
	router.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v want %v", status, http.StatusFound)
	}
	mockStore.AssertExpectations(t)
}

func Test_updateNotasHandler(t *testing.T) { //PASSED
	mockStore := InitMockStore()
	testData := nota
	mockStore.On("UpdateNotas", testData.spected, testData.id).Return(nil).Once()
	req, err := http.NewRequest("PUT", "/Crud/"+testData.id+"/Notas", bytes.NewBufferString(testData.form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	router := router()
	router.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v wanted %v", status, http.StatusFound)
	}
	mockStore.AssertExpectations(t)
}

////////////////////////////////////////////////////////////////////////////////////////////
//PASSED

func newTaskForm() *url.Values {
	form := url.Values{}
	form.Set("tarea", "Test")
	form.Set("completo", "true")
	return &form
}

var tarea = pruebaT{
	id:      "Personal",
	spected: &Task{ID: 0, Done: true, Text: "Test"},
	form:    newTaskForm(),
}

func Test_createTareasHandler(t *testing.T) { //PASSED
	mockStore := InitMockStore()
	testData := tarea
	log.Println(testData)
	mockStore.On("CreateTareas", testData.spected, testData.id).Return(nil).Once()
	req, err := http.NewRequest("POST", "/Crud/"+testData.id+"/Tarea", bytes.NewBufferString(testData.form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	router := router()
	router.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v want %v", status, http.StatusFound)
	}
	mockStore.AssertExpectations(t)
}

func Test_deleteTareasHandler(t *testing.T) { //PASSED
	mockStore := InitMockStore()
	testData := pruebaT{
		id:      "Personal",
		spected: &Task{ID: 6},
	}
	mockStore.On("DeleteTareas", testData.spected, testData.id).Return(nil).Once()
	req, err := http.NewRequest("DELETE", "/Crud/"+(testData.id)+"/Tarea/"+strconv.Itoa(testData.spected.ID), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	router := router()
	router.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v want %v", status, http.StatusFound)
	}
	mockStore.AssertExpectations(t)
}

func Test_updateTareasHandler(t *testing.T) {
	mockStore := InitMockStore()
	testData := tarea
	mockStore.On("UpdateTareas", testData.spected, testData.id).Return(nil).Once()
	req, err := http.NewRequest("PUT", "/Crud/"+testData.id+"/Tarea", bytes.NewBufferString(testData.form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	router := router()
	router.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code, got %v wanted %v", status, http.StatusFound)
	}
	mockStore.AssertExpectations(t)
}
