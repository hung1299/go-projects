package task

import (
	"encoding/json"
	"fmt"
	"hung1299/roadmap-projects/task-tracker/internal/file"
	"time"
)

type Status int

const FileName = "test.json"

const (
	StatusTodo = iota
	StatusInProgress
	StatusDone
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func GetAllTasks() []Task {
	bs := file.Read(FileName)
	ts := Tasks{}

	if err := json.Unmarshal(bs, &ts); err != nil {
		fmt.Println(err)
	}
	
	return ts.Tasks
}

func NewTask(description string) Task {
	task := Task{
		Description: description,
		Status: StatusTodo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tasks := GetAllTasks()
	
	if len(tasks) > 0 {
		lastTaskID := tasks[len(tasks) - 1].ID
		task.ID = lastTaskID + 1;
	}

	return task
}
