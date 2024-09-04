package app

import (
	"fmt"
	"hung1299/go-projects/task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all tasks that can be filter by status",
	Run:   listTask,
}

func printTasks(ts task.Tasks) {
	for _, t := range ts {
		fmt.Println(t.ID, t)
	}
}

func listTask(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		ts := task.GetByStatus(task.All)
		printTasks(ts)
	} else {
		sa := args[0]
		s, ok := task.StatusMap[sa]
		if !ok {
			fmt.Println("Status not found")
			return
		}

		ts := task.GetByStatus(s)
		printTasks(ts)
	}
}
