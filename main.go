package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/web_logger/handlers"
	"github.com/web_logger/models"
	"github.com/web_logger/services"
)

var (
	pathToConfig = "./config/config.yml"
)

func main() {
	var conf models.Config
	fmt.Printf("Load settings from %s", pathToConfig)

	conf.LoadFromYaml(pathToConfig)
	fmt.Println("Success!")

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

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Server.Port), r))
}
