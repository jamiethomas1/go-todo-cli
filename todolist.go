package main

import "fmt"

type TodoItem struct {
	Task     string `json:"task"`
	Complete bool   `json:"complete"`
}

func (ti *TodoItem) toggleComplete() {
	ti.Complete = !ti.Complete
}

type TodoList struct {
	Items []TodoItem `json:"items"`
}

func (tl *TodoList) push(ti TodoItem) {
	tl.Items = append(tl.Items, ti)
}

func (tl *TodoList) getTask(index int) (string, error) {
	if index < 0 || index >= len(tl.Items) {
		return "", fmt.Errorf("index out of range")
	}

	return tl.Items[index].Task, nil
}

func (tl *TodoList) getComplete(index int) (bool, error) {
	if index < 0 || index >= len(tl.Items) {
		return false, fmt.Errorf("index out of range")
	}

	return tl.Items[index].Complete, nil
}

func (tl *TodoList) toggleItem(index int) error {
	if index < 0 || index >= len(tl.Items) {
		return fmt.Errorf("index out of range")
	}

	tl.Items[index].Complete = !tl.Items[index].Complete

	return nil
}

func (tl *TodoList) drop(index int) error {
	if index < 0 || index >= len(tl.Items) {
		return fmt.Errorf("index out of range")
	}

	tl.Items = slice_remove(tl.Items, index)

	return nil
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

func slice_remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
