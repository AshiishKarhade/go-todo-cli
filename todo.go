package todo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

// An item represents a task
type item struct {
	Task        string    // task name
	Done        bool      // whether the task is completed or not
	CreatedAt   time.Time // timestamp of creation of task
	CompletedAt time.Time // timestamp of completion of task
}

// array of items
type Todos []item

// Add function accepts a task name (string) and adds it to our TODO list
func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

// Complete function accepts index number of the task and marks it completed
func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true
	return nil
}

// Delete function accepts index number of the removes the task from TODO list
func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

// function to load the file - .json file where our tasks will be stored on system
func (t *Todos) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

// function to write the tasks to our file - .json file where our tasks wil be stored on system
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}
