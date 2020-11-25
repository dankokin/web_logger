package main

import (
	"fmt"
	"github.com/dankokin/web_logger/models"
	"github.com/dankokin/web_logger/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dankokin/web_logger/handlers"
)

var (
	pathToConfig = "appsettings.json"
	pathToScheme = "./database/init.sql"
)

func main() {
	conf := models.InitConfigFile(pathToConfig)
	fmt.Println(conf)
	db, err := services.NewDB(conf)
	if err != nil {
		panic(err)
	}

	services.Setup(pathToScheme, db)
	fmt.Println("Database is ready!")

	env := handlers.DataStoreEnvironment{
		Db: db,
	}

	r := mux.NewRouter()

	r.HandleFunc("/new_log", env.SaveLog).Methods(http.MethodPost)
	r.HandleFunc("/logs", env.GetLogs).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.ServerPort), r))
}
