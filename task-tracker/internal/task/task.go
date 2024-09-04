package task

import (
	"encoding/json"
	"fmt"
	"hung1299/go-projects/task-tracker/internal/file"
	"os"
	"slices"
	"time"
)

const FileName = "test.json"

type Tasks []Task

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const (
	All = iota
	StatusTodo
	StatusInProgress
	StatusDone
)

var StatusMap = map[string]int{
	"todo":        StatusTodo,
	"in-progress": StatusInProgress,
	"done":        StatusDone,
}

var MarkStatusMap = map[string]int{
	"mark-todo":        StatusTodo,
	"mark-in-progress": StatusInProgress,
	"mark-done":        StatusDone,
}

func (ts Tasks) save() {
	bs, err := json.MarshalIndent(ts, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file.Save(FileName, bs)
}

func getTasksFromFile() Tasks {
	bs := file.Read(FileName)
	ts := Tasks{}

	if err := json.Unmarshal(bs, &ts); err != nil {
		// fmt.Println("Task is empty")
	}

	return ts
}

func GetByStatus(s int) (fts Tasks) {
	ts := getTasksFromFile()

	if s == All {
		return ts
	}

	for _, t := range ts {
		if t.Status == s {
			fts = append(fts, t)
		}
	}

	return fts
}

func New(d string) Task {
	t := Task{
		ID:          0,
		Description: d,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	ts := getTasksFromFile()

	if len(ts) > 0 {
		lastTaskID := ts[len(ts)-1].ID
		t.ID = lastTaskID + 1
	}

	ts = append(ts, t)
	ts.save()

	return t
}

func UpdateDescription(nd string, id int) bool {
	ts := getTasksFromFile()

	idx := slices.IndexFunc(ts, func(t Task) bool {
		return t.ID == id
	})

	if idx == -1 {
		return false
	}

	ts[idx].Description = nd
	ts[idx].UpdatedAt = time.Now()
	ts.save()
	return true
}

func UpdateStatus(s int, id int) bool {
	ts := getTasksFromFile()

	idx := slices.IndexFunc(ts, func(t Task) bool {
		return t.ID == id
	})

	if idx == -1 {
		return false
	}

	ts[idx].Status = s
	ts[idx].UpdatedAt = time.Now()
	ts.save()
	return true
}

func Remove(id int) bool {
	ts := getTasksFromFile()

	idx := slices.IndexFunc(ts, func(t Task) bool {
		return t.ID == id
	})

	if idx == -1 {
		return false
	}

	ts = append(ts[:idx], ts[idx+1:]...)
	ts.save()
	return true
}
