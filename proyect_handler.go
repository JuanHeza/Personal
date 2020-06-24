package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	//APIKey is the WakaTime key
	APIKey = "502f3c9e-67d4-48ce-a6b9-77dbe3887e7c"
)

func proyectHandler(w http.ResponseWriter, r *http.Request) {
	var actual Projects
	errorPage := Templates.Lookup("proyect")
	vars := mux.Vars(r)
	actual = ProyectData[vars["name"]]
	actual.Time = getWakaTime(vars["name"])
	if err := errorPage.ExecuteTemplate(w, "proyect", actual); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getWakaTime(Project string) []string {
	var data interface{} //= make(map[string]string)
	url := fmt.Sprintf("https://wakatime.com/api/v1/users/current/stats/last_year?api_key=%s&project=%s", APIKey, Project)
	res, err := http.Get(url)
	if err != nil {
		return nil
		// panic(err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&data)
	fmt.Println(url)
	x := data.(map[string]interface{})
	y := x["data"].(map[string]interface{})
	if len(y["categories"].([]interface{})) > 0 {
		z := y["categories"].([]interface{})[0].(map[string]interface{})
		return []string{fmt.Sprint(z["hours"]), fmt.Sprint(z["minutes"]), fmt.Sprint(z["seconds"])}
	}
	return nil //"Time Not Initialized"
}
