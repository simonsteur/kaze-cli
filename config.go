package main

import (
	"encoding/json"
	"io/ioutil"
)

//Config struct
type Config struct {
	Sensu string `json:"sensu_server"`
	Port  string `json:"sensu_server_port"`
}

// Cfg reads the config.json file in the configuration dir
func Cfg() Config {
	// if configFile == "" {
	// 	handleWarning("no configuration file specified")
	// }
	// file, err := ioutil.ReadFile(configFile)
	file, err := ioutil.ReadFile("./configuration/config.json")
	if err != nil {
		handleError(err)

	}
	var cf Config
	err = json.Unmarshal(file, &cf)
	if err != nil {
		handleError(err)

	}
	return cf
}
