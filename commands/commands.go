package commands

import (
	"cmp"
	"errors"
	"fmt"
	"os"
	"slices"
	"tasks/models"
	"tasks/utils"
	"text/tabwriter"
	"time"
)

func ListTasks(showAll bool) error {
	tasks, err := utils.ReadCSV()
	if err != nil {
		return err
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

func AddTask(description string) error {
	tasks, err := utils.ReadCSV()
	if err != nil {
		return err
	}

	lastTask := slices.MaxFunc(tasks, func(a, b models.Task) int {
		return cmp.Compare(a.ID, b.ID)
	})

	lastId := lastTask.ID

	task := models.Task{ID: lastId + 1, Description: description, Created: time.Now(), Done: false}

	tasks = append(tasks, task)

	if err := utils.WriteCSV(tasks); err != nil {
		return err
	}

	fmt.Println("Successfully added Task!")

	return nil
}

func CompleteTask(id int) error {
	return errors.New("not implemented")
}

func DeleteTask(id int) error {
	return errors.New("not implemented")
}
