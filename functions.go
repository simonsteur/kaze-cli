package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Bulk struct used for bulk client creation of clients
type Bulk struct {
	Clients []Client `json:"clients"`
	Stashes []Stash  `json:"Stashes"`
	Results []Result `json:"Results"`
}

// Client struct used in
type Client struct {
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	Subscriptions []string `json:"subscriptions"`
	Environment   string   `json:"environment"`
}

// Stash struct
type Stash struct {
	Path    string      `json:"path"`
	Content interface{} `json:"content"`
	Expire  string      `json:"expire"`
}

// Result Struct
type Result struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	Output string `json:"output"`
	Status int    `json:"status"`
}

//readFileClients reads in the json file specified and places it in the clients struct.
func readFileBulk(f string) Bulk {
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

//readFile reads in the json file specified and places it in a struct.
func readFile(f string) []byte {
	file, err := ioutil.ReadFile(f)
	if err != nil {
		handleError(err)
	}
	return file
}

//resultHandler handles how the results from the functions should be returned.
func resultHandler(res []byte) {
	result := prettyJSON(string(res))
	fmt.Printf(result)
}

//kazeList lists all return values or a single value
func kazeList(api, value string) {
	req := new(request)
	req.Method = "GET"
	if value != "" {
		req.URL = api + "/" + value
	} else {
		req.URL = api
	}
	res := doSensuAPIRequest(req)
	if string(res) == "" && value != "" {
		trowError(value + "not found.")
	}
	if string(res) == "" {
		trowError("something went wrong, no results returned.")
	}
	resultHandler(res)
}

//kazeDelete deletes the specified object
func kazeDelete(api, value string) {
	if value == "" {
		trowError("no name specified.")
	}
	req := new(request)
	req.Method = "DELETE"
	req.URL = api + "/" + value
	res := doSensuAPIRequest(req)
	resultHandler(res)
}

// //kazeCreateClient creates a proxyclient
// func kazeCreateClient(api, name, address, environment string, subscriptions []string) {

// 	if name == "" {
// 		trowError("no name specified.")
// 	}
// 	if address == "" {
// 		trowError("no address specified.")
// 	}
// 	if environment == "" {
// 		trowError("no environment specified.")
// 	}

// 	pl := &Proxyclient{
// 		Name:          name,
// 		Address:       address,
// 		Subscriptions: subscriptions,
// 		Environment:   environment,
// 	}

// 	payload, err := json.Marshal(pl)
// 	if err != nil {
// 		handleError(err)
// 	}
// 	req := new(request)
// 	req.Method = "POST"
// 	req.URL = api
// 	req.Payload = payload
// 	res := doSensuAPIRequest(req)
// 	resultHandler(res)
// }

//kazeCreate
// func kazeCreate(api, file string) {
// 	values := readFileClients(file)
// 	for _, value := range values.Proxyclient {
// 		payload, err := json.Marshal(Client)
// 		if err != nil {
// 			handleError(err)
// 		}
// 		req := new(request)
// 		req.Method = "POST"
// 		req.URL = api
// 		req.Payload = payload
// 		res := doSensuAPIRequest(req)
// 		resultHandler(res)
// 	}
// }

//kazeCreate
func kazeCreate(api, file string) {
	payload := readFile(file)
	req := new(request)
	req.Method = "POST"
	req.URL = api
	req.Payload = payload
	fmt.Print(string(payload))
	res := doSensuAPIRequest(req)
	fmt.Print(res)
	resultHandler(res)
}
