package app

import (
	"fmt"
	"hung1299/go-projects/task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var AddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to list",
	Long:  "Add a new task to list",
	Args:  cobra.ExactArgs(1),
	Run:   addTask,
}

func addTask(cmd *cobra.Command, args []string) {
	description := args[0]
	t := task.New(description)

	fmt.Printf("Task added successfully (ID: %v)", t.ID)
}
