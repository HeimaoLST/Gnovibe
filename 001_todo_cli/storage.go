package main

import (
	"encoding/json"
	"fmt"
	"log"
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
type todoRepo interface {
	Save(string) error
	Done(int) error
	List() ([]todoItem, error)
}

type todoJson struct {
	file string
}

func NewtodoJson(str string) todoJson {
	return todoJson{
		file: str,
	}
}
func (uc *todoJson) Save(str string) error {

	data, err := os.ReadFile(uc.file)
	if err != nil {
		return err
	}
	var items []todoItem

	if err := json.Unmarshal(data, &items); err != nil {
		log.Fatal("Json unmarshal error: ", err.Error())
		return err
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

func (uc *todoJson) Done(index int) error {
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
func (uc *todoJson) List() ([]todoItem, error) {

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
