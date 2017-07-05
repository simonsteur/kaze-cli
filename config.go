package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//Config struct
type Config struct {
	Sensu string `json:"sensu_server"`
	Port  string `json:"sensu_server_port"`
}

// Cfg reads the config.json file in the configuration dir
func Cfg() Config {
	if _, err := os.Stat(configfile); os.IsNotExist(err) {
		kazeCreateConfigFile("127.0.0.0", "")
	}
	file, err := ioutil.ReadFile(configfile)
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
