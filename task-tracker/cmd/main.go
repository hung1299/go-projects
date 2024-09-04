package main

import (
	"fmt"
	"hung1299/go-projects/task-tracker/internal/file"
	"hung1299/go-projects/task-tracker/internal/task"
)

func main() {
	// cmd := task.InitRootCommand()

	// err := cmd.Execute()

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	file.Save("test.json", []byte(`{"tasks": []}`))

	fmt.Println("tasks", task.GetAllTasks())
}