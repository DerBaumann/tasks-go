package models

import (
	"encoding/csv"
	"os"
	"strconv"
	"tasks/config"
	"time"
)

type Task struct {
	ID          int
	Description string
	Created     time.Time
	Done        bool
}

func ReadTasksFromStore() ([]Task, error) {
	file, err := os.Open(config.TasksFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 4
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tasks []Task

	for _, row := range data {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, err
		}

		description := row[1]

		created, err := time.Parse(time.RFC3339, row[2])
		if err != nil {
			return nil, err
		}

		done, err := strconv.ParseBool(row[3])
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, Task{ID: id, Description: description, Created: created, Done: done})
	}

	return tasks, nil
}
