package main

import (
	"flag"
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
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		showAll := listCmd.Bool("a", false, "Shows all tasks including completed ones")

		if err := listCmd.Parse(os.Args[2:]); err != nil {
			log.Fatal(err)
		}

		if err := commands.ListTasks(*showAll); err != nil {
			log.Fatal(err)
		}
	case "add":
		if len(os.Args) < 3 {
			log.Fatal(config.HelpMessage)
		}
		description := os.Args[2]

		if err := commands.AddTask(description); err != nil {
			log.Fatal(err)
		}
	case "complete":
		if len(os.Args) < 3 {
			log.Fatal(config.HelpMessage)
		}
		taskid, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		if err := commands.CompleteTask(taskid); err != nil {
			log.Fatal(err)
		}
	case "delete":
		if len(os.Args) < 3 {
			log.Fatal(config.HelpMessage)
		}
		taskid, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		if err := commands.DeleteTask(taskid); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal(config.HelpMessage)
	}
}
