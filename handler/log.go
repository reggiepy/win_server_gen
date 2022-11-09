package handler

import (
	"fmt"
)

var Verbose bool

func VerboseLog(message string) {
	if Verbose {
		fmt.Println(message)
	}
}
