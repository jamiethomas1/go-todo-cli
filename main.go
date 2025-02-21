package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading environment variables")
	}

	tl := readTodoList()

	if len(os.Args) == 1 {
		tl.show()
		return
	}

	config := ParseFlags()

	if config.Add != "" {
		ti := TodoItem{
			Task:     config.Add,
			Complete: false,
		}

		log.Printf("Adding value \"%s\"\n", config.Add)

		tl.push(ti)
	} else if config.Remove != -1 {
		item, err := tl.getTask(config.Remove - 1)
		if err != nil {
			log.Fatal("error getting todo item at given index")
		}

		err = tl.drop(config.Remove - 1)
		if err != nil {
			log.Fatal("error removing todo item at given index")
		}

		log.Printf("Removed item \"%s\"\n", item)
	} else if config.Toggle != -1 {
		item, err := tl.getTask(config.Toggle - 1)
		if err != nil {
			log.Fatal("error getting todo item at given index")
		}

		err = tl.toggleItem(config.Toggle - 1)
		if err != nil {
			log.Fatal("error toggling item")
		}

		var status string
		s, err := tl.getComplete(config.Toggle - 1)
		if err != nil {
			log.Fatal("error getting completion status of task at given index")
		}
		if s {
			status = "Complete"
		} else {
			status = "Not Complete"
		}

		log.Printf("Toggled item \"%s\" to \"%s\"\n", item, status)
	}

	err = writeTodoList(&tl)
	if err != nil {
		log.Fatal("error writing todo list to file")
	}
}
