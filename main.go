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
		item, err := tl.get(config.Remove - 1)
		if err != nil {
			log.Fatal("error getting todo item at given index")
		}

		log.Printf("Removing item \"%s\"\n", item)

		err = tl.drop(config.Remove - 1)
		if err != nil {
			log.Fatal("error removing todo item at given index")
		}
	}

	err = writeTodoList(&tl)
	if err != nil {
		log.Fatal("error writing todo list to file")
	}
}
