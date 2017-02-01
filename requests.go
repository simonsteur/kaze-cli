package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func doSensuAPIRequest(method, url string, payload []byte) []byte {
	// form request
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(payload))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "application/json")
	req.Header.Add("cache-control", "no-cache")
	// do request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	// read and close response
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	result := body
	// return result
	return result
}
