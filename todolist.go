package main

import "fmt"

type TodoItem struct {
	Task     string `json:"task"`
	Complete bool   `json:"complete"`
}

type TodoList struct {
	Items []TodoItem `json:"items"`
}

func (tl *TodoList) push(ti TodoItem) {
	tl.Items = append(tl.Items, ti)
}

func (tl *TodoList) show() {
	fmt.Println("To-do List")

	for _, item := range tl.Items {
		var checkbox rune
		if item.Complete {
			checkbox = '☑'
		} else {
			checkbox = '☐'
		}
		fmt.Printf("%c %s\n", checkbox, item.Task)
	}
}
