package task

import (
	"github.com/spf13/cobra"
)


func InitRootCommand() *cobra.Command {
	rc := &cobra.Command{
		Use:   "task-tracker",
		Short: "[Short] Task Tracker is a CLI to tracking your task",
		Long: "[Long] Task Tracker is a CLI to tracking your task",
	}

	rc.AddCommand(AddCommand)
	rc.AddCommand(ListCommand)

	return rc
}