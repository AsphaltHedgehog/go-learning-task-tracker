package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"test-task/logger"
)

const path = "\\AppData\\Local\\Temp\\go-task-tracker\\"
const file = "tasks.json"

func pathToFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	pathToFile := homeDir + path + file

	fmt.Print(pathToFile)
	return pathToFile, nil
}

func isTaskFileExist(filePath string) (bool, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return false, err
	}

	logger.LogVerbose(false, "FileInfo: name %s, size %d, last modification %s", fileInfo.Name(), fileInfo.Size(), fileInfo.ModTime())

	return true, nil
}

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

func isValidTaskJson(path string) (bool, error) {
	fileB, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	var task Task
	err = json.Unmarshal(fileB, &task)
	if err != nil {
		return false, err
	}

	return true, nil
}

func createDirectory() error {
	err := os.MkdirAll(path, os.FileMode(0644))
	if err != nil {
		return err
	}
	return nil
}

func generateDefaultData() ([]byte, error) {
	tasks := []Task{
		Task{
			ID:          1,
			Description: "Hi, welcome to this little project, feel free to change my status, delete or use any other CLI tool :3",
			Status:      StatusInProgress,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	jsonBTasks, err := json.Marshal(tasks)
	if err != nil {
		return nil, err
	}

	return jsonBTasks, nil
}

func createFile() error {
	defaultTasksData, err := generateDefaultData()
	if err != nil {
		return err
	}

	err = os.WriteFile(path+file, defaultTasksData, os.FileMode(0644))
	if err != nil {
		return err
	}

	return nil
}

func FsInit() bool {
	filePath, err := pathToFile()
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return false
	}

	isFileExist, err := isTaskFileExist(filePath)
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return false
	}

	if !isFileExist {
		err := createDirectory()
		if err != nil {
			log.Panic(err)
			return false
		}
		err = createFile()
		if err != nil {
			log.Panic(err)
			return false
		}

		return false
	}

	isFileFormatValid, err := isValidTaskJson(filePath)
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return false
	}

	if !isFileFormatValid {
		err := createFile()
		if err != nil {
			log.Panic(err)
			return false
		}
		return false
	}

	return true
}
