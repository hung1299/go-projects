package app

import (
	"fmt"
	"hung1299/go-projects/task-tracker/internal/task"
	"strconv"

	"github.com/spf13/cobra"
)

var UpdateCommand = &cobra.Command{
	Use:   "update",
	Short: "update a new task to list",
	Long:  "update a new task to list",
	Args:  cobra.ExactArgs(2),
	Run:   updateTask,
}

func updateTask(_ *cobra.Command, args []string) {
	idArgs := args[0]

	ID, err := strconv.Atoi(idArgs)
	if err != nil {
		fmt.Println("ID must be number!!")
		return
	}

	nd := args[1]
	result := task.UpdateDescription(nd, ID)
	if result {
		fmt.Printf("Task[%v] updated!!", ID)
	} else {
		fmt.Printf("Task with ID[%v] not found!!", ID)
	}
}
