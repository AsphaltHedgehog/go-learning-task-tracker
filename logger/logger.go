package logger

import (
	"fmt"
	"log"
)

const red = "\033[31m"
const grn = "\033[32m"
const def = "\033[0m"

var Verbose bool = false

func LogVerbose(isError bool, format string, a ...any) {
	if !Verbose {
		return
	}

	msg := fmt.Sprintf(format, a...)

	if isError {
		log.Println(red + msg + def)
		return
	}

	log.Println(grn + msg + def)
}
