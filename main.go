package main

import (
	"flag"
	"log"

	"test-task/logger"
	"test-task/task"
)

var Verbose bool

func init() {
	flag.BoolVar(&Verbose, "verbose", false, "enable verbose logging")
}

func main() {
	flag.Parse()
	if Verbose {
		log.SetFlags(log.Ltime | log.Lshortfile)
		logger.Verbose = true
		logger.LogVerbose(false, "Verbose mod")
	}

	isFileSystemInit := task.FsInit()
	logger.LogVerbose(false, "Is program file system ready: %t", isFileSystemInit)
}
