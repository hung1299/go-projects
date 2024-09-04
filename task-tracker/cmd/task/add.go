package task

import "github.com/spf13/cobra"

var AddCommand = &cobra.Command{
	Use: "add",
	Short: "Add a new task to list",
	Long: "Add a new task to list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
