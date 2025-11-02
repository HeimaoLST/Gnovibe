package main

import (
	"fmt"
	"strconv"
)

type TodoUsecase struct {
	repo TodoRepo
}

func NewtodoUsecase(repo TodoRepo) *TodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}

func (uc *TodoUsecase) todoHelp() {
	fmt.Println("Usage:")

	fmt.Println("./todo add string \n eg. ./todo add helloworld\n add an entry to your todo list")
	fmt.Println("./todo list\n list all todo entries you have input")
	fmt.Println("./todo done index\n eg. ./todo done 1\n to mark the entry has been done")
}

func (uc *TodoUsecase) todoAdd(something string) {
	if something == "" {
		fmt.Println("you should input the valid string")
		return
	}
	if err := uc.repo.Save(something); err != nil {
		fmt.Println("Save Error: ", err.Error())
		return
	}

	fmt.Println("Your todo has been saved")

}

func (uc *TodoUsecase) todoDone(index string) {
	num, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("Atoi Error: ", err.Error())
		return
	}
	if err = uc.repo.Done(num - 1); err != nil {
		fmt.Println("Done Error: ", err.Error())
		return
	}

	fmt.Println("The todo entry you chose has done")

}

func (uc *TodoUsecase) todoList() {
	items, err := uc.repo.List()

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	fmt.Println("index\t", "name\t", "status")

	for i, item := range items {
		fmt.Printf("%d\t%s\t", i+1, item.Name)
		if item.Status == STATUS_DONE {
			fmt.Printf("Done\n")
		} else {
			fmt.Printf("Doing\n")
		}
	}

}
