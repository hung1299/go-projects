package app

import (
	"github.com/spf13/cobra"
)

func InitRootCommand() *cobra.Command {
	rc := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task Tracker is a CLI to tracking your task",
		Long:  "Task Tracker is a CLI to tracking your task",
	}

	rc.AddCommand(AddCommand)
	rc.AddCommand(DeleteCommand)
	rc.AddCommand(UpdateCommand)
	rc.AddCommand(ListCommand)
	rc.AddCommand(MarkTodoCommand)
	rc.AddCommand(MarkInProgressCommand)
	rc.AddCommand(MarkDoneCommand)

	return rc
}
