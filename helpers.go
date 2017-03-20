package main

import (
	"bytes"
	"encoding/json"
)

// //parseCSString parses a comma seperated string and makes sure no emtpy strings are added to the array
// func parseCSString(input string) []string {
// 	var output []string
// 	parsed := strings.Split(input, ",")
// 	for _, v := range parsed {
// 		if v != "" {
// 			output = append(output, v)
// 		}
// 	}
// 	return output
// }

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
