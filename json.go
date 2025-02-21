package main

import (
	"encoding/json"
	"log"
	"os"
)

func readTodoList() TodoList {
	_, err := os.Stat(os.Getenv("GO_TODOLIST_PATH"))
	if os.IsNotExist(err) {
		log.Println("tasks file does not exist. creating...")
		_, err := os.Create(os.Getenv("GO_TODOLIST_PATH"))
		if err != nil {
			log.Fatal("error creating tasks file")
		}

		return TodoList{}
	}

	data, err := os.ReadFile(os.Getenv("GO_TODOLIST_PATH"))
	if err != nil {
		log.Fatal("error reading tasks file")
	}

	if !json.Valid(data) {
		log.Fatal("invalid json in tasks file")
	}

	var tl TodoList
	err = json.Unmarshal(data, &tl)
	if err != nil {
		log.Fatal("error unmarshaling tasks file json")
	}

	return tl
}

func writeTodoList(tl *TodoList) error {
	file, err := os.OpenFile(os.Getenv("GO_TODOLIST_PATH"), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("error opening tasks file")
	}

	js, err := json.Marshal(tl)
	if err != nil {
		log.Fatal("error marshaling json")
	}

	_, err = file.Write(js)
	if err != nil {
		log.Fatal("error writing to file")
	}

	return nil
}
