package task

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	"test-task/logger"
)

const tempUserDir = "AppData\\Local\\Temp\\go-task-tracker"
const file = "tasks.json"

var savedFilePath string

func pathToFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, tempUserDir), nil
}

func isTaskFileExist(filePath string) (bool, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logger.LogVerbose(true, "Task file does not exist")
			return false, nil
		}
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

	var tasks []Task
	err = json.Unmarshal(fileB, &tasks)
	if err != nil {
		var syntaxError *json.SyntaxError
		var typeError *json.UnmarshalTypeError

		if errors.As(err, &syntaxError) || errors.As(err, &typeError) {
			logger.LogVerbose(true, "Task file corrupted, resetting it to default state")
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func createDirectory(path string) error {
	err := os.MkdirAll(path, os.FileMode(0640))
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

func createFile(filePath string) error {
	defaultTasksData, err := generateDefaultData()
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, defaultTasksData, os.FileMode(0640))
	if err != nil {
		return err
	}

	return nil
}

func FsInit() bool {
	folderPath, err := pathToFile()
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return false
	}

	filePath := filepath.Join(folderPath, file)
	savedFilePath = filePath
	isFileExist, err := isTaskFileExist(filePath)
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return false
	}

	if !isFileExist {
		err := createDirectory(folderPath)
		if err != nil {
			log.Panic(err)
		}
		err = createFile(filePath)
		if err != nil {
			log.Panic(err)
		}
	}

	isFileFormatValid, err := isValidTaskJson(filePath)
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return false
	}

	if !isFileFormatValid {
		err := createFile(filePath)
		if err != nil {
			log.Panic(err)
		}
	}

	return true
}

type NewTaskDraft struct {
	Description string `json:"description"`
	Status      Status `json:"status"`
}

func readFile() ([]Task, error) {
	fileBArray, err := os.ReadFile(savedFilePath)
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return nil, err
	}

	var tasks []Task

	if err = json.Unmarshal(fileBArray, &tasks); err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return nil, err
	}

	return tasks, nil
}

func appendToArray(newElement Task, array []Task) error {
	newArray := append(array, newElement)
	updatedArr, err := json.Marshal(&newArray)
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return err
	}

	err = os.WriteFile(savedFilePath, updatedArr, 0644)
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return err
	}

	return nil
}

func AddTask(task NewTaskDraft) (*Task, error) {
	tasks, err := readFile()
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return nil, err
	}

	lastId := tasks[len(tasks)-1].ID
	currentTime := time.Now()

	newTask := Task{
		ID:          lastId + 1,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	err = appendToArray(newTask, tasks)
	if err != nil {
		return nil, err
	}

	return &newTask, nil
}
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return nil, err
	}

	err = os.WriteFile(savedFilePath, updatedTasksArr, 0644)
	if err != nil {
		logger.LogVerbose(true, "Error: %v", err)
		return nil, err
	}

	return &newTask, nil
}
