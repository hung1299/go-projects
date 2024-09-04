package task

import "github.com/spf13/cobra"

var ListCommand = &cobra.Command{
	Use: "list",
	Short: "List all tasks",
	Long: "List all tasks that can be filter by status",
	Run: func(cmd *cobra.Command, args []string) {

	},
}