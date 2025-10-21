package utils

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"tasks/config"
	"tasks/models"
	"time"
)

func ReadCSV() ([]models.Task, error) {
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

	var tasks []models.Task

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

		tasks = append(tasks, models.Task{ID: id, Description: description, Created: created, Done: done})
	}

	return tasks, nil
}

func WriteCSV(tasks []models.Task) error {
	return errors.New("not implemented")
}
