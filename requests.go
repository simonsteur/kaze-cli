package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type request struct {
	Method  string
	URL     string
	Payload []byte
}

func doSensuAPIRequest(request *request) []byte {
	// form request
	req, _ := http.NewRequest(request.Method, request.URL, bytes.NewBuffer(request.Payload))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "application/json")
	req.Header.Add("cache-control", "no-cache")
	// do request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		handleError(err)
	}
	// read and close response
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode == 400 {
		trowError("400 (Bad Request)")
	}
	if res.StatusCode == 500 {
		trowError("500 (Internal Server Error)")
	}
	defer res.Body.Close()
	result := body
	// return result
	return result
}
