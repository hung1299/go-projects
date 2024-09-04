package app

import (
	"fmt"
	"hung1299/go-projects/task-tracker/internal/task"
	"strconv"

	"github.com/spf13/cobra"
)

var MarkTodoCommand = &cobra.Command{
	Use:   "mark-todo",
	Short: "mark task is todo",
	Long:  "mark task is todo",
	Args:  cobra.ExactArgs(1),
	Run:   markTask,
}
var MarkInProgressCommand = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "mark task is in progress",
	Long:  "mark task is in progress",
	Args:  cobra.ExactArgs(1),
	Run:   markTask,
}
var MarkDoneCommand = &cobra.Command{
	Use:   "mark-done",
	Short: "mark task is done",
	Long:  "mark task is done",
	Args:  cobra.ExactArgs(1),
	Run:   markTask,
}

func markTask(cmd *cobra.Command, args []string) {
	idArgs := args[0]

	ID, err := strconv.Atoi(idArgs)
	if err != nil {
		fmt.Println("ID must be number!!")
		return
	}

	mark := cmd.Use
	status := task.MarkStatusMap[mark]

	result := task.UpdateStatus(status, ID)
	if result {
		fmt.Printf("Task[%v]'s status updated!!", ID)
	} else {
		fmt.Printf("Task with ID[%v] not found!!", ID)
	}
}
