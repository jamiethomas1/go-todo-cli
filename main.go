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
	}

	err = writeTodoList(&tl)
	if err != nil {
		log.Fatal("error writing todo list to file")
	}
}
