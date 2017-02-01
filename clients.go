package main

import "fmt"
import "encoding/json"

type proxyclient struct {
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	Subscriptions []string `json:"subscriptions"`
	Environment   string   `json:"environment"`
}

func createNewClient(name, address, environment string, subscriptions []string) {
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
	req := new(request)
	req.Method = "POST"
	req.URL = clientapi
	req.Payload = payload
	res := doSensuAPIRequest(req)
	result := prettyJSON(string(res))
	fmt.Printf(result)
}

func deleteClient(client string) {
	req := new(request)
	req.Method = "DELETE"
	req.URL = clientapi + "/" + client
	res := doSensuAPIRequest(req)
	result := prettyJSON(string(res))
	fmt.Printf(result)
}

func listClients(client string) {
	req := new(request)
	req.Method = "GET"
	if client != "" {
		req.URL = clientapi + "/" + client
	} else {
		req.URL = clientapi
	}
	res := doSensuAPIRequest(req)
	result := prettyJSON(string(res))
	fmt.Printf(result)
}
