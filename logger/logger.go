package logger

import (
	"fmt"
	"log"
)

const grn = "\033[32m"
const def = "\033[0m"

func LogVerbose(verbose bool, isError bool, format string, a ...interface{}) {
	if !verbose {
		return
	}

	msg := fmt.Sprintf(format, a...)
	log.Println(grn + msg + def)
}
