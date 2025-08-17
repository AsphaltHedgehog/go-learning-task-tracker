package task

import (
	"fmt"
)

func AddTool(draft NewTaskDraft) {
	newTask := AddTask(draft.Description, draft.Status)

	fmt.Printf("[%d] %s\n    (%s)\n    Created at: %s\n", newTask.ID, newTask.Description, newTask.Status, newTask.CreatedAt)
}

func UpdateTool(id int, description string, status Status) {
	updatedTask := UpdateTask(id, description, status)
	if updatedTask == nil {
		fmt.Printf("Task not found")
		return
	}

	fmt.Printf("Updated task:\n[%d] %s\n    (%s)\n    Created at: %s Last update: %s\n", updatedTask.ID, updatedTask.Description, updatedTask.Status, updatedTask.CreatedAt, updatedTask.UpdatedAt)
}

func DeleteTool(id int) {
	removedTask := RemoveTask(id)

	if removedTask == nil {
		fmt.Printf("Task not found")
		return
	}

	fmt.Printf("Deleted task:\n[%d] %s\n    (%s)\n    Created at: %s Last update: %s\n", removedTask.ID, removedTask.Description, removedTask.Status, removedTask.CreatedAt, removedTask.UpdatedAt)

}

func UpdateStatusTool(id int, status string) {
	updatedTask := UpdateTaskStatus(id, Status(status))
	if updatedTask == nil {
		fmt.Println("Task not found or unsupported status change")
		return
	}

	fmt.Printf("[%d] %s\n    (%s)\n    Created at: %s Last update: %s\n", updatedTask.ID, updatedTask.Description, updatedTask.Status, updatedTask.CreatedAt, updatedTask.UpdatedAt)
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
