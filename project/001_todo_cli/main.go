package main

import (
	"fmt"
	"os"
)

const (
	OPERATION_ADD  = "add"
	OPERATION_LIST = "list"
	OPERATION_DONE = "done"
)

func main() {
	repo := NewTodoJson("todo.json")
	todouc := NewtodoUsecase(&repo)
	if len(os.Args) < 2 {
		todouc.todoHelp()
		return
	}
	if os.Args[1] != OPERATION_LIST && len(os.Args) != 3 {
		fmt.Println("Error: Missing required argument")
		fmt.Println("Usage:")
		todouc.todoHelp()
		return
	}
	switch os.Args[1] {
	case OPERATION_ADD:
		{
			todouc.todoAdd(os.Args[2])
		}
	case OPERATION_LIST:
		{
			todouc.todoList()
		}

	case OPERATION_DONE:
		{
			todouc.todoDone(os.Args[2])
		}
	}

}
