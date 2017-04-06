package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func usagePrint() {
	fmt.Printf("kaze-cli is command line interface tool for sensu operations\n\n")
	fmt.Printf("Usage:\n")
	fmt.Print("  kaze ", os.Args[1], " [options]\n\n")
	fmt.Printf("Parameters:\n")
}

//pretty JSON turns json input into a more readably and pretty json string
func prettyJSON(input string) string {
	var output bytes.Buffer
	err := json.Indent(&output, []byte(input), "", "\t")
	if err != nil {
		handleWarning("failed to make json pretty, sorry.")
		return input
	}
	return output.String()
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

func postPayload(api string, payload []byte) []byte {
	req := new(request)
	req.Method = "POST"
	req.URL = api
	req.Payload = payload
	res := doSensuAPIRequest(req)
	resultHandler(res)
	return res
}
