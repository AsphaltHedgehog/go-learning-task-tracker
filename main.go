package main

import (
	"flag"
	"log"
)

var Verbose bool

func init() {
	flag.BoolVar(&Verbose, "verbose", false, "enable verbose logging")
}

func main() {
	flag.Parse()
	if Verbose {
		log.SetFlags(log.Ltime | log.Lshortfile)
		LogVerbose("Verbose mod")
	}
}
