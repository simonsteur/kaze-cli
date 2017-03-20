package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Bulk struct used for bulk client creation of clients
type Bulk struct {
	Client []Client `json:"clients"`
	Stash  []Stash  `json:"Stashes"`
	Result []Result `json:"Results"`
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
func kazeList(api string, values []string) {
	req := new(request)
	req.Method = "GET"
	if len(values) != 0 {
		for _, value := range values {
			req.URL = api + "/" + value
			res := doSensuAPIRequest(req)
			if string(res) == "" && value != "" {
				trowError(value + "not found.")
			}
			if string(res) == "" {
				trowError("something went wrong, no results returned.")
			}
			resultHandler(res)
		}
	} else {
		req.URL = api
		res := doSensuAPIRequest(req)
		if string(res) == "" && len(values) != 0 {
			trowError("not found.")
		}
		if string(res) == "" {
			trowError("something went wrong, no results returned.")
		}
		resultHandler(res)
	}
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

func kazeCreateClient() {
	if file != "" {
		bulk := readFileBulk(file)
		for _, v := range bulk.Client {
			payload, err := json.Marshal(v)
			if err != nil {
				handleError(err)
			}
			req := new(request)
			req.Method = "POST"
			req.URL = clientsapi
			req.Payload = payload
			res := doSensuAPIRequest(req)
			result := prettyJSON(string(res))
			fmt.Printf(result)
		}
	} else {
		cl := &Client{
			Name:          name[1],
			Address:       clientAddress,
			Subscriptions: clientSubscriptions,
			Environment:   clientEnvironment,
		}
		payload, err := json.Marshal(cl)
		if err != nil {
			handleError(err)
		}
		req := new(request)
		req.Method = "POST"
		req.URL = clientsapi
		req.Payload = payload
		res := doSensuAPIRequest(req)
		result := prettyJSON(string(res))
		fmt.Printf(result)
	}
}

func kazeCreateBulkResults() {
	bulk := readFileBulk(file)
	for _, v := range bulk.Result {
		payload, err := json.Marshal(v)
		if err != nil {
			handleError(err)
		}
		req := new(request)
		req.Method = "POST"
		req.URL = resultsapi
		req.Payload = payload
		res := doSensuAPIRequest(req)
		result := prettyJSON(string(res))
		fmt.Printf(result)
	}
}

func kazeCreateBulkStashes() {
	bulk := readFileBulk(file)
	for _, v := range bulk.Result {
		payload, err := json.Marshal(v)
		if err != nil {
			handleError(err)
		}
		req := new(request)
		req.Method = "POST"
		req.URL = stashesapi
		req.Payload = payload
		res := doSensuAPIRequest(req)
		result := prettyJSON(string(res))
		fmt.Printf(result)
	}
}
