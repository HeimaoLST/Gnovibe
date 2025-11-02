package main

import (
	"fmt"
	"strconv"
)

type todoUsecase struct {
	repo todoRepo
}

func NewtodoUsecase(repo todoRepo) *todoUsecase {
	return &todoUsecase{
		repo: repo,
	}
}

func (uc *todoUsecase) todoHelp() {

}

func (uc *todoUsecase) todoAdd(something string) {
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

func (uc *todoUsecase) todoDone(index string) {
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

func (uc *todoUsecase) todoList() {
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
