package app

import (
	"hung1299/go-projects/github-user-activity/internal/user"

	"github.com/spf13/cobra"
)

func InitializeCommand() *cobra.Command {
	return &cobra.Command{
		Use: "github-activity",
		Short: "GitHub Activity is a simple tool to fetch the recent activity",
		Long : "GitHub Activity is a simple tool to fetch the recent activity of a Github user and display it in the terminal",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			user.GetEvents(args[0])
		},
	}
}