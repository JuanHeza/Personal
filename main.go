package main

import (
	"os"

	"log"
	"net/http"

	"github.com/JuanHeza/Personal/models"
	// "github.com/JuanHeza/Personal/controllers"
)

var (
	host     = os.Getenv("PQ_HOST")
	port     = os.Getenv("PQ_PORT")
	user     = os.Getenv("PQ_USER")
	password = os.Getenv("PQ_PASSWORD")
	dbname   = os.Getenv("PQ_DBNAME")
	ssl 	 = os.Getenv("PQ_SSL")
	//APIKey is the WakaTime key
)

func main() {
	models.StartConnection(host, port, user, password, dbname, ssl)
	models.SetupDatabaseDevelopment()
	r := newRouter()
	models.Static.ReadStatics()
	log.Println("listening http://127.0.0.1:8080/Home/")
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
