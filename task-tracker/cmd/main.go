package main

import (
	"fmt"
	"hung1299/go-projects/task-tracker/cmd/app"
	"os"
)

func main() {
	cmd := app.InitRootCommand()

	err := cmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
