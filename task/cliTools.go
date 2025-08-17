package task

import (
	"fmt"
)

func AddTool(draft NewTaskDraft) {
	newTask := AddTask(draft.Description, draft.Status)

	fmt.Printf("[%d] %s\n    (%s)\n    Created at: %s\n", newTask.ID, newTask.Description, newTask.Status, newTask.CreatedAt)
}

func UpdateTool(id int, description string) {

}

func DeleteTool(id int) {

}

func UpdateStatusTool(id int, status string) {

}

func ListTool(status Status) {
	tasks := ListTasks()
	anyFound := false

	for _, t := range *tasks {
		if status != "" && t.Status != status {
			continue
		}
		anyFound = true
		fmt.Printf("[%d] %s\n    (%s)\n    Created at: %s Last update: %s\n", t.ID, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
	}

	if !anyFound {
		fmt.Println("Nothing found")
	}
}
