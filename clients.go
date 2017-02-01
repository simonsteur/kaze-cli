package main

import "fmt"
import "encoding/json"

type proxyclient struct {
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	Subscriptions []string `json:"subscriptions"`
	Environment   string   `json:"environment"`
}

func createNewProxyClient() {
	method := "POST"
	url := "http://" + config.Sensu + ":" + config.Port + "/clients"
	pl := &proxyclient{
		Name:          "main-fw",
		Address:       "192.168.201.253",
		Subscriptions: []string{"network"},
		Environment:   "production",
	}
	payload, err := json.Marshal(pl)
	if err != nil {
		handleError(err)
	}
	res := doSensuAPIRequest(method, url, payload)
	fmt.Print(string(res))
}
