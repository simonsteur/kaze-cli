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

func help() {
	fmt.Printf("kaze-cli is command line interface tool for sensu operations\n\n")
	fmt.Printf("Usage:\n")
	fmt.Print("  kaze [command] [options]\n\n")
	fmt.Printf("Commands:\n")
	fmt.Print("  list             list objects\n")
	fmt.Print("  create-client    creates a proxy client\n")
	fmt.Print("  create-result    creates a check result\n")
	fmt.Print("  create-stash     creates a stash\n")
	fmt.Print("  delete           delete clients, results, stashes\n")
	fmt.Print("  clear-silence    clear a silence entry\n")
	fmt.Print("  check            request to schedule a check\n")
	fmt.Print("  resolve          resolve a check result\n")
	fmt.Print("  help             print help text\n\n\n")
	fmt.Print("for help use: kaze [command] -help")
	fmt.Print("\n\n")
	os.Exit(1)
}

//pretty JSON turns json input into a more readable and pretty json string
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

//readFile reads in the json file specified
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
	res, statusCode := doSensuAPIRequest(req)
	if statusCode == 204 || statusCode == 201 {
		fmt.Print("success")
	} else {
		resultHandler(res)
	}
	return res
}
