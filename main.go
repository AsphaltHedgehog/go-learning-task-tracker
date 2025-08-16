package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
		log.SetFlags(log.Ltime)
		logger.Verbose = true
		logger.LogVerbose(false, "Verbose mod")
	}

	isFileSystemInit := task.FsInit()
	logger.LogVerbose(false, "Is program file system ready: %t", isFileSystemInit)

	if len(os.Args) > 1 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "add":
		break

	case "update":
		break

	case "delete":
		break

	case "mark":
		break

	case "list":
		break

	default:
		printHelp()
	}

}

func printHelp() {
	fmt.Println("Use this commands to access task-cli tool:")
	fmt.Println("  add [description]          				Add new task with description")
	fmt.Println("  update [id] [description, status]  Update task, can use without status argument")
	fmt.Println("  delete [id]                     		Delete task")
	fmt.Println("  mark in-progress [id]           		Update status to in-progress")
	fmt.Println("  mark done [id]                  		Update status to done")
	fmt.Println("  list empty/done/todo/in-progress   List all task or specified by status (status: todo, in-progress, done)")
}
