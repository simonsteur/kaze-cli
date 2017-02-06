package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"os"

	"github.com/urfave/cli"
)

//Bulk struct used for bulk creation / deleting of clients
type Bulk struct {
	Proxyclient []Proxyclient `json:"clients"`
}

//Proxyclient struct
type Proxyclient struct {
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	Subscriptions []string `json:"subscriptions"`
	Environment   string   `json:"environment"`
}

func readBulkfile(f string) Bulk {
	var bulk Bulk
	file, err := ioutil.ReadFile(f)
	if err != nil {
		handleError(err)
	}
	err = json.Unmarshal(file, &bulk)
	if err != nil {
		handleError(err)
	}
	return bulk
}

func manageClient(c *cli.Context) {

	message := "you selected more than one operation, please only specify one operation (list, create, delete)."

	//check if more than one operation has been specified
	if clientList && clientDelete && clientCreate {
		trowError(message)
		os.Exit(1)
	}
	if clientList && clientDelete {
		trowError(message)
		os.Exit(1)
	}
	if clientList && clientCreate {
		trowError(message)
		os.Exit(1)
	}
	if clientDelete && clientCreate {

		trowError(message)
		os.Exit(1)
	}
	if clientDelete && clientList {
		trowError(message)
		os.Exit(1)
	}

	// execute operations

	if clientList {
		listClients(clientName)
	}
	if clientCreate {
		if clientBulk == true {
			createNewClientBulk(clientBulkFile)

		} else {
			createNewClient(clientName, clientAddress, clientEnvironment, clientSubscriptions)
		}
	}
	if clientDelete {
		deleteClient(clientName)
	}
}

func createNewClientBulk(file string) {
	clients := readBulkfile(file)
	for _, client := range clients.Proxyclient {
		payload, err := json.Marshal(client)
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
}

func createNewClient(name, address, environment string, subscriptions []string) {

	if name == "" {
		trowError("no name specified.")
	}
	if address == "" {
		trowError("no address specified.")
	}
	if environment == "" {
		trowError("no environment specified.")
	}

	pl := &Proxyclient{
		Name:          name,
		Address:       address,
		Subscriptions: subscriptions,
		Environment:   environment,
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
	if clientName == "" {
		trowError("no client name specified.")
	}
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
	if string(res) == "" && client != "" {
		trowError(client + "not found.")
	}
	if string(res) == "" {
		trowError("something went wrong, no results returned.")
	}
	result := prettyJSON(string(res))
	fmt.Printf(result)
}
