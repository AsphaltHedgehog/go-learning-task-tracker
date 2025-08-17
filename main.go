package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"test-task/logger"
	"test-task/task"
)

var Verbose bool

func init() {
	if len(os.Args) == 1 {
		return
	}

	for _, arg := range os.Args[1:] {
		if arg == "-v" || arg == "--verbose" {
			Verbose = true
			break
		}
	}
}

func main() {
	if Verbose {
		log.SetFlags(log.Ltime)
		logger.Verbose = true
		logger.LogVerbose(false, "Verbose mod")
	}

	isFileSystemInit := task.FsInit()
	logger.LogVerbose(false, "Is program file system ready: %t", isFileSystemInit)

	if len(os.Args) == 1 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 4 {
			printHelp()
			break
		}

		taskDraft := task.NewTaskDraft{
			Description: os.Args[2],
			Status:      task.Status(os.Args[3]),
		}

		task.AddTool(taskDraft)

	case "update":
		break

	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Input correct id", os.Args[2])
			return
		}
		task.DeleteTool(id)

	case "mark":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Input correct id", os.Args[2])
			return
		}
		task.UpdateStatusTool(id, os.Args[3])

	case "list":
		if len(os.Args) < 3 {
			task.ListTool("")
			break
		}
		task.ListTool(task.Status(os.Args[2]))

	case "help":
		printHelp()

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

	fmt.Println("Use -v/--verbose for log info")
}
