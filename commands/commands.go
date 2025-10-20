package commands

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"tasks/config"
	"tasks/models"
	"text/tabwriter"
	"time"
)

func ListTasks(showAll bool) error {
	file, err := os.Open(config.TasksFile)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 4
	data, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var tasks []models.Task

	for _, row := range data {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			return err
		}

		description := row[1]

		created, err := time.Parse(time.RFC3339, row[2])
		if err != nil {
			return err
		}

		done, err := strconv.ParseBool(row[3])
		if err != nil {
			return err
		}

		tasks = append(tasks, models.Task{ID: id, Description: description, Created: created, Done: done})
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, "ID\tDescription\tCreated At\tCompleted")
	for _, task := range tasks {
		timeStr := task.Created.Format("02.01.2006 - 15:04:05")

		var doneStr string
		if task.Done {
			doneStr = "yes"
		} else {
			doneStr = "no"
		}

		if showAll || task.Done {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Description, timeStr, doneStr)
		}
	}

	w.Flush()

	return nil
}

func AddTask(description string) {
	log.Fatal("Not implemented!")
}

func CompleteTask(id int) {
	log.Fatal("Not implemented!")
}

func DeleteTask(id int) {
	log.Fatal("Not implemented!")
}
