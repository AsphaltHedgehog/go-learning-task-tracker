package task

import (
	"fmt"
)

func AddTool(description string) {

}

func UpdateTool(id int, description string) {

}

func DeleteTool(id int) {

}

func UpdateStatusTool(id int, status string) {

}

func printHelp() {
	fmt.Println("List func")
	fmt.Println("  task-cli list [status]")
	fmt.Println("Show all tasks or filter them by status")
	fmt.Println("Args:  status (todo | in-progress | done)")
	fmt.Println("Available calls:")
	fmt.Println("  task-cli list")
	fmt.Println("  task-cli list todo")
	fmt.Println("  task-cli list in-progress")
	fmt.Println("  task-cli list done")
}

func List(status Status) {
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
