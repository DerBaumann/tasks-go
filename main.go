package main

import (
	"log"
	"os"
	"strconv"
	"tasks/commands"
	"tasks/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(config.HelpMessage)
	}

	switch os.Args[1] {
	case "list":
		commands.ListTasks()
	case "add":
		if len(os.Args) < 3 {
			log.Fatal(config.HelpMessage)
		}
		description := os.Args[2]

		commands.AddTask(description)
	case "complete":
		if len(os.Args) < 3 {
			log.Fatal(config.HelpMessage)
		}
		taskid, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		commands.CompleteTask(taskid)
	case "delete":
		if len(os.Args) < 3 {
			log.Fatal(config.HelpMessage)
		}
		taskid, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		commands.DeleteTask(taskid)
	default:
		log.Fatal(config.HelpMessage)
	}
}
