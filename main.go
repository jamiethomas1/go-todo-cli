package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
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

	in := bufio.NewReader(os.Stdin)
	var ti TodoItem

	fmt.Println("Add new todo item:")

	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatal("error reading input")
	}

	ti.Task = strings.TrimSpace(line)
	ti.Complete = false

	tl.push(ti)

	err = writeTodoList(&tl)
	if err != nil {
		log.Fatal("error writing TodoItem to file")
	}

	fmt.Printf("Task \"%s\" added with completion status: %t\n", ti.Task, ti.Complete)
}
