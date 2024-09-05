package app

import (
	"fmt"
	"hung1299/go-projects/task-tracker/internal/task"
	"strconv"

	"github.com/spf13/cobra"
)

var DeleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "delete a task out of the list",
	Long:  "delete a task out of the list",
	Args:  cobra.ExactArgs(1),
	Run:   deleteTask,
}

func deleteTask(_ *cobra.Command, args []string) {
	idArgs := args[0]

	ID, err := strconv.Atoi(idArgs)
	if err != nil {
		fmt.Println("ID must be number!!")
		return
	}

	result := task.Remove(ID)
	if result {
		fmt.Printf("Task[%v] deleted!!", ID)
	} else {
		fmt.Printf("Task with ID[%v] not found!!", ID)
	}
}
