package main

import (
	"os"

	"github.com/fatih/color"
)

func handleError(err error) {

	color.Red("error: %s", err)
	os.Exit(1)
}

func handleWarning(warning string) {

	color.Yellow("Warning: %s", warning)
}

func checkValueWarnings(values map[string]string) {

	for k, v := range values {
		if v == "" {
			handleWarning(k + " not specified, using default value")
		}
	}
}
