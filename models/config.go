package models

import (
	"encoding/json"
	"io/ioutil"
)

func InitConfigFile(cnfFile string) Config {
	jsonFile, err := ioutil.ReadFile(cnfFile)

	var cnf Config
	jerr := json.Unmarshal(jsonFile, &cnf)

	if err != nil || jerr != nil {
		panic(err)
	}

	return cnf
}
