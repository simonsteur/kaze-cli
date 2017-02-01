package main

import (
	"bytes"
	"encoding/json"
)

func prettyJSON(input string) string {
	var output bytes.Buffer
	err := json.Indent(&output, []byte(input), "", "\t")
	if err != nil {
		handleWarning("failed to make json pretty, sorry.")
		return input
	}
	return output.String()
}
