package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dankokin/web_logger/models"
	"github.com/dankokin/web_logger/services"
	"github.com/dankokin/web_logger/utils"
)

type DataStoreEnvironment struct {
	Db services.DataStore
}

// Function gets all logs from database by target name which were received during the 1 day in default case
// or in the interval specified by the user and returns them in JSON-format
// params: interval, target;
// default: interval = 1 day, target = all
func (env *DataStoreEnvironment) GetLogs(w http.ResponseWriter, r *http.Request) {
	args := make(map[string]interface{}, 2)
	interval := r.URL.Query().Get("interval")
	if interval != "" {
		if err := utils.CheckInterval(interval); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		args["interval"], _ = strconv.Atoi(interval)
	} else {
		args["interval"] = 1
	}

	target := r.URL.Query().Get("target")
	if target != "" {
		if err := utils.CheckTarget(target); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		args["target"] = target
	}

	logs, err := env.Db.GetAllLogs(args)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(logs)
}

// Simple handler which contain message from waf to database
func (env *DataStoreEnvironment) SaveLog(w http.ResponseWriter, r *http.Request) {
	var log models.WAFMessage
	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil {
		fmt.Println(r.Body)
		fmt.Println(log)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = env.Db.SaveMessage(log)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
