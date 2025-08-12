package main

import (
	"fmt"
	"log"
)

const grn = "\033[32m"
const def = "\033[0m"

func LogVerbose(format string, a ...interface{}) {
	if !Verbose {
		return
	}

	msg := fmt.Sprintf(format, a...)
	log.Println(grn + msg + def)
}
