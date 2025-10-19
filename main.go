package main

import (
	"fmt"
	"log"
	"os"
	"tasks/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(config.HelpMessage)
	}

	switch os.Args[1] {
	case "list":
		fmt.Println("Listing Tasks")
	case "add":
		if len(os.Args) < 3 {
			log.Fatal(config.HelpMessage)
		}
		task := os.Args[2]

		fmt.Printf("Adding task: %s\n", task)
	case "complete":
		if len(os.Args) < 3 {
			log.Fatal(config.HelpMessage)
		}
		taskid := os.Args[2]

		fmt.Printf("completing task: %s\n", taskid)
	case "delete":
		if len(os.Args) < 3 {
			log.Fatal(config.HelpMessage)
		}
		taskid := os.Args[2]

		fmt.Printf("Deleting task: %s\n", taskid)
	default:
		log.Fatal(config.HelpMessage)
	}
}
