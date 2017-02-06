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

func trowError(err string) {
	color.Red("error: %s", err)
	os.Exit(1)
}
