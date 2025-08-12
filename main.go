package main

import (
	"flag"
	"log"

	"test-task/logger"
)

var Verbose bool

func init() {
	flag.BoolVar(&Verbose, "verbose", false, "enable verbose logging")
}

func main() {
	flag.Parse()
	if Verbose {
		log.SetFlags(log.Ltime | log.Lshortfile)
		logger.LogVerbose(Verbose, false, "Verbose mod")
	}
}
