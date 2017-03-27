package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Bulk struct used for bulk client creation of clients
type Bulk struct {
	Client []Client `json:"clients"`
	Stash  []Stash  `json:"Stashes"`
	Result []Result `json:"Results"`
}

// Client struct
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
	Expire  int         `json:"expire"`
}

// Result Struct
type Result struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	Output string `json:"output"`
	Status int    `json:"status"`
}

// Clear struct
type Clear struct {
	Subscription string
	Check        string
	Client       string
	ID           string
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
func kazeDelete(api string, values []string) {
	for _, v := range values {
		req := new(request)
		req.Method = "DELETE"
		req.URL = api + "/" + v
		res := doSensuAPIRequest(req)
		resultHandler(res)
	}
}

func kazeCreateClient() {
	if file != "" {
		bulk := readFileBulk(file)
		for _, v := range bulk.Client {
			payload, err := json.Marshal(v)
			if err != nil {
				handleError(err)
			}
			postPayload(clientsapi, payload)
		}
	} else {
		str := &Client{
			Name:          name[1],
			Address:       clientAddress,
			Subscriptions: clientSubscriptions,
			Environment:   clientEnvironment,
		}
		payload, err := json.Marshal(str)
		if err != nil {
			handleError(err)
		}
		postPayload(clientsapi, payload)
	}
}

func kazeCreateResult() {
	if file != "" {
		bulk := readFileBulk(file)
		for _, v := range bulk.Result {
			payload, err := json.Marshal(v)
			if err != nil {
				handleError(err)
			}
			postPayload(resultsapi, payload)
		}
	} else {
		str := &Result{
			Name:   name[1],
			Source: resultSource,
			Output: resultOutput,
			Status: resultStatus,
		}
		payload, err := json.Marshal(str)
		if err != nil {
			handleError(err)
		}
		postPayload(resultsapi, payload)
	}
}

func kazeCreateStash() {
	if file != "" {
		bulk := readFileBulk(file)
		for _, v := range bulk.Stash {
			payload, err := json.Marshal(v)
			if err != nil {
				handleError(err)
			}
			postPayload(stashesapi, payload)
		}
	} else {
		str := &Stash{
			Path:    stashPath,
			Content: stashContent,
			Expire:  stashExpire,
		}
		payload, err := json.Marshal(str)
		if err != nil {
			handleError(err)
		}
		postPayload(stashesapi, payload)
	}
}

func kazeClear(values []string) {
	if client {
		for _, v := range values {
			str := &Clear{
				Client: v,
				Check:  silenceCheckName,
			}
			payload, err := json.Marshal(str)
			if err != nil {
				handleError(err)
			}
			postPayload(silencedapiclear, payload)
		}
	}
	if silenceSubscription {
		for _, v := range values {
			str := &Clear{
				Subscription: v,
				Check:        silenceCheckName,
			}
			payload, err := json.Marshal(str)
			if err != nil {
				handleError(err)
			}
			postPayload(silencedapiclear, payload)
		}
	}
	if all {
		type ID []struct {
			ID string `json:"id"`
		}
		var id ID
		req := new(request)
		req.Method = "GET"
		req.URL = silencedapi
		res := doSensuAPIRequest(req)
		if string(res) == "" {
			fmt.Print("no silenced entries to be cleared.")
			os.Exit(0)
		}
		json.Unmarshal(res, &id)

		for _, v := range id {
			str := &Clear{
				ID: v.ID,
			}
			payload, err := json.Marshal(str)
			if err != nil {
				handleError(err)
			}
			postPayload(silencedapiclear, payload)
		}
	}
	fmt.Print("cleared all silenced entries.")
}
