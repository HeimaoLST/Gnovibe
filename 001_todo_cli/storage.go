package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Status int

const (
	STATUS_DOING = iota
	STATUS_DONE
)

type todoItem struct {
	Name       string
	CreateTime time.Time
	ModifyTime time.Time
	Status     Status
}
type TodoRepo interface {
	Save(string) error
	Done(int) error
	List() ([]todoItem, error)
}

type TodoJson struct {
	file string
}

func NewTodoJson(str string) TodoJson {
	return TodoJson{
		file: str,
	}
}
func (uc *TodoJson) Save(str string) error {

	data, err := os.ReadFile(uc.file)
	if err != nil {
		return err
	}
	var items []todoItem

	if err := json.Unmarshal(data, &items); err != nil {
		//in fact the log.Fatal will casue the exit
		//the express after it will never execute
		// log.Fatal("Json unmarshal error: ", err.Error())

		return fmt.Errorf("JSON unmarshal error: %w", err)

	}
	items = append(items, todoItem{
		Name:       str,
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
		Status:     STATUS_DOING,
	})
	// transform the object to string
	data, err = json.Marshal(items)
	if err != nil {
		return err
	}
	return os.WriteFile(uc.file, data, 0644)

}

func (uc *TodoJson) Done(index int) error {
	if index < 0 {
		return fmt.Errorf("the index is invalid")
	}

	data, err := os.ReadFile(uc.file)
	if err != nil {
		return err
	}

	var items []todoItem
	if err = json.Unmarshal(data, &items); err != nil {
		return err
	}

	if index > len(items)-1 {
		return fmt.Errorf("the index is invalid")

	}

	// olditem := items[index]
	// olditem.status = STATUS_DONE
	// items[index] = olditem

	items[index].Status = STATUS_DONE

	data, err = json.Marshal(items)
	if err != nil {
		return err
	}

	return os.WriteFile(uc.file, data, 0644)

}
func (uc *TodoJson) List() ([]todoItem, error) {

	data, err := os.ReadFile(uc.file)

	if err != nil {
		return nil, err
	}
	// may 1KB is enough

	var items []todoItem
	if err = json.Unmarshal(data, &items); err != nil {
		return nil, err
	}

	// if len(items) == 0 {
	// 	return nil, fmt.Errorf("your todolist is empty, plz try again after adding the entry")
	// }

	return items, nil
}
