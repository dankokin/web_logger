package main

import (
	"fmt"
	"github.com/web_logger/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/web_logger/handlers"
	"github.com/web_logger/services"
)

var (
	pathToConfig = "appsettings.json"
)

func main() {
	conf := models.InitConfigFile(pathToConfig)
	db, err := services.NewDB(conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database is ready!")

	env := handlers.DataStoreEnvironment{
		Db: db,
	}

	r := mux.NewRouter()

	r.HandleFunc("/new_log", env.SaveLog).Methods(http.MethodPost)
	r.HandleFunc("/logs", env.GetLogs).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.ServerPort), r))
}
