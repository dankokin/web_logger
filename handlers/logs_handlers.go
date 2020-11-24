package handlers

import "net/http"

// Function gets all logs from database by target name which were received during the 1 day in default case
// or in the interval specified by the user and returns them in JSON-format
// params: interval, target;
// default: interval = 1 day, target = all
func GetLogs(w http.ResponseWriter, r *http.Request) {

}
