package main

import (
	"fmt"
	"hung1299/go-projects/github-user-activity/cmd/app"
	"os"
)

func main() {
	cmd := app.InitializeCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}