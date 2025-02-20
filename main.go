package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type TodoItem struct {
	Task     string `json:"task"`
	Complete bool   `json:"complete"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading environment variables")
	}

	tl := getTodoList()

	if len(os.Args) == 1 {
		tl.show()
		return
	}
}
