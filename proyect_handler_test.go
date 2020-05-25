package main

import(
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetProyectsHandler(t *Testing.T){
	mockStore := InitMockStore()

	mockStore.On("GetProyects").Return([]*Proyect{
		{
	},nil).Once()

	req, err := http.NewRequest("GET","",nil)

	if err != nil{
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandleFunc(getProyectsHandler)
	hf.ServerHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK{
		t.Errorf("handler Returned wrong status code, got %v want %v", status, http.StatusOK)
	}

	expected := Proyect{"PokemonTCG","A small tcg game"}
	b := []Proyect{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil{
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler Returned unexpected Body, got %v, want %v", actual,expected)
	}

	mockStore.AssertExpectations(t)
}

func TestCreateProyectHandler(t *testing.T){
	mockStore := InitMockStore()

	mockStore.On("CreateProyect", &Proyect{"PokemonTCG","A small tcg game"}).Return(nil)

	form := newCreateProyectForm()
	req, err := http.NewRequest("POST","",bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "appplication/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Iota(len(form.Encode())))

	if err != nil{
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandleFunc(createProyectHandler)

	hf.ServeHTTP(recorder,req)

	if status := recorder.Code; status != http.StatusFound{
		t.Errorf("Handler Returned wrong status code,, got %v want %v", staus, http.StatusOK)
	}

	mockStore.AssertExpectations(t)
}

