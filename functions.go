package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Name          string   `json:"name,omitempty"`
	Address       string   `json:"address,omitempty"`
	Subscriptions []string `json:"subscriptions,omitempty"`
	Environment   string   `json:"environment,omitempty"`
}

// Stash struct
type Stash struct {
	Path    string      `json:"path,omitempty"`
	Content interface{} `json:"content,omitempty"`
	Expire  int         `json:"expire,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

// Result Struct
type Result struct {
	Source string `json:"source,omitempty"`
	Name   string `json:"name,omitempty"`
	Output string `json:"output,omitempty"`
	Status int    `json:"status,omitempty"`
}

// Clear struct
type Clear struct {
	Subscription string `json:"subscription,omitempty"`
	Check        string `json:"check,omitempty"`
	ID           string `json:"id,omitempty"`
}

// ID struct
type ID []struct {
	ID string `json:"id,omitemtpy"`
}

// Silence struct
type Silence struct {
	Subscription    string `json:"subscription,omitempty"`
	Check           string `json:"check,omitempty"`
	ID              string `json:"id,omitempty"`
	Creator         string `json:"ceator,omitempty"`
	ExpireOnResolve bool   `json:"expire_on_resolve,omitempty"`
	Expire          int    `json:"expire,omitempty"`
	Reason          string `json:"reason,omitempty"`
}

// CheckRequest struct
type CheckRequest struct {
	Subscribers []string `json:"subscribers,omitempty"`
	Check       string   `json:"check,omitempty"`
}

// Checks struct
type Checks []struct {
	Subscribers []string `json:"subscribers,omitempty"`
	Check       string   `json:"name,omitempty"`
}

// Events struct
type Events []struct {
	Client struct {
		Name string `json:"name,omitempty"`
	}
	Check struct {
		Name string `json:"name,omitempty"`
	}
}

// Resolve struct
type Resolve struct {
	Client string `json:"client,omitempty"`
	Check  string `json:"check,omitempty"`
}

//kazeConfigure creates a configuration file for kaze-cli to use
func kazeConfigure() {
	path := "/etc/kaze-cli/config.json"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Print("creating configuration file...")
		kazeCreateConfigFile(address, port, path)
	} else {
		fmt.Print("config file present, do you wish to override?")
		confirimation := confirm()
		if confirimation {
			fmt.Print("overriding configuration...")
			kazeCreateConfigFile(address, port, path)
		} else {
			fmt.Print("no action taken.")
		}
	}
}

func kazeCreateConfigFile(address, port, path string) {
	if port == "" {
		port = "4567"
	}

	c := &Config{
		Sensu: address,
		Port:  port,
	}
	output, err := json.Marshal(c)
	if err != nil {
		handleError(err)
	}
	if _, err := os.Stat("/etc/kaze-cli"); os.IsNotExist(err) {
		os.MkdirAll("/etc/kaze-cli", 0664)
	}
	ioutil.WriteFile(path, output, 0644)
}

//kazeList lists all return values or a single value
func kazeList(api string, values []string) {
	req := new(request)
	req.Method = "GET"
	if len(values) != 0 {
		for _, value := range values {
			req.URL = api + "/" + value
			res, _ := doSensuAPIRequest(req)
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
		res, _ := doSensuAPIRequest(req)
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
func kazeDelete(api string, values []string, checkname string) {
	//fmt.Print(values)
	if result {
		for _, v := range values {
			req := new(request)
			req.Method = "DELETE"
			req.URL = api + "/" + v + "/" + checkname
			doSensuAPIRequest(req)
			fmt.Print("deleted result")
		}
	} else {
		for _, v := range values {
			req := new(request)
			req.Method = "DELETE"
			req.URL = api + "/" + v
			res, statusCode := doSensuAPIRequest(req)
			if statusCode == 204 {
				fmt.Print("success")
			} else {
				resultHandler(res)
			}
		}
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
		s := &Client{
			Name:          name[0],
			Address:       clientAddress,
			Subscriptions: clientSubscriptions,
			Environment:   clientEnvironment,
		}
		payload, err := json.Marshal(s)
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
		s := &Result{
			Name:   name[0],
			Source: resultSource,
			Output: resultOutput,
			Status: resultStatus,
		}
		payload, err := json.Marshal(s)
		if err != nil {
			handleError(err)
		}
		postPayload(resultsapi, payload)
	}
}

func kazeCreateStash() {
	payload := readFile(file)
	postPayload(stashesapi, payload)
}

func kazeSilence(values []string) {
	if all {
		type checkname []struct {
			Check string `json:"check"`
		}
		var c checkname
		req := new(request)
		req.Method = "GET"
		req.URL = checksapi
		res, _ := doSensuAPIRequest(req)
		if string(res) == "" {
			fmt.Print("no checks found.")
			os.Exit(0)
		}
		json.Unmarshal(res, &c)

		for _, v := range c {
			s := &Silence{
				Check:           v.Check,
				Reason:          silenceReason,
				Creator:         silenceCreator,
				Expire:          silenceExpire,
				ExpireOnResolve: silenceExpireOnResolve,
			}
			payload, err := json.Marshal(s)
			if err != nil {
				handleError(err)
			}
			postPayload(silencedapi, payload)
		}
	} else {
		for _, v := range values {
			s := &Silence{
				Check:           silenceCheckName,
				Reason:          silenceReason,
				Creator:         silenceCreator,
				Expire:          silenceExpire,
				ExpireOnResolve: silenceExpireOnResolve,
			}
			if client {
				s.Subscription = "client:" + v
			}
			if silenceSubscription {
				s.Subscription = v
			}

			payload, err := json.Marshal(s)
			if err != nil {
				handleError(err)
			}
			postPayload(silencedapi, payload)
		}
	}
}

func kazeClear(values []string) {

	if all {
		var id ID
		req := new(request)
		req.Method = "GET"
		req.URL = silencedapi
		res, _ := doSensuAPIRequest(req)
		if string(res) == "" {
			fmt.Print("no silenced entries to be cleared.")
			os.Exit(0)
		}
		json.Unmarshal(res, &id)

		for _, v := range id {
			s := &Clear{
				ID: v.ID,
			}
			payload, err := json.Marshal(s)
			if err != nil {
				handleError(err)
			}
			postPayload(silencedapiclear, payload)
		}
		fmt.Print("cleared all silenced entries.")
	} else {
		for _, v := range values {
			s := &Clear{
				Check: silenceCheckName,
			}
			if client {
				s.ID = "client:" + v + ":" + silenceCheckName
			}
			if silenceSubscription {
				s.Subscription = v
			}
			payload, err := json.Marshal(s)
			if err != nil {
				handleError(err)
			}
			postPayload(silencedapiclear, payload)
		}
	}
}

func kazeCheck(values []string) {
	if all {
		var c Checks
		req := new(request)
		req.Method = "GET"
		req.URL = checksapi
		res, _ := doSensuAPIRequest(req)
		if string(res) == "" {
			trowError("something went wrong, no results returned.")
		}
		json.Unmarshal(res, &c)

		for _, v := range c {
			s := &CheckRequest{
				Check:       v.Check,
				Subscribers: v.Subscribers,
			}
			payload, err := json.Marshal(s)
			if err != nil {
				handleError(err)
			}
			postPayload(requestapi, payload)
		}
	} else {
		for _, v := range values {
			s := &CheckRequest{
				Check:       v,
				Subscribers: checkTarget,
			}
			payload, err := json.Marshal(s)
			if err != nil {
				handleError(err)
			}
			postPayload(requestapi, payload)
		}
	}
}

func kazeResolve(values []string) {
	if all {
		var e Events
		req := new(request)
		req.Method = "GET"
		req.URL = eventsapi
		res, _ := doSensuAPIRequest(req)
		if string(res) == "" {
			trowError("something went wrong, no events to resolve returned.")
		}
		json.Unmarshal(res, &e)

		for _, v := range e {
			s := &Resolve{
				Client: v.Client.Name,
				Check:  v.Check.Name,
			}
			payload, err := json.Marshal(s)
			if err != nil {
				handleError(err)
			}
			postPayload(resolveapi, payload)
		}

	} else {
		for _, v := range values {
			s := &Resolve{
				Client: v,
				Check:  resolveCheckName,
			}
			payload, err := json.Marshal(s)
			if err != nil {
				handleError(err)
			}
			postPayload(resolveapi, payload)
		}
	}
}
