package db

import (
	"fmt"
	"time"
)

type Todo struct {
	ID    int       `json:"id"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
	Time  time.Time `json:"time"`
}

var todos = []*Todo{
	&Todo{
		ID:    1,
		Title: "hehe",
		Body:  "TEST",
		Time:  time.Now().Add(2 * time.Hour),
	},
	&Todo{
		ID:    2,
		Title: "hehe2",
		Body:  "TEST2",
		Time:  time.Now().Add(4 * time.Hour),
	},
}

func Find() ([]*Todo, error) {
	return todos, nil
}

func FindByID(id int) (*Todo, error) {
	for _, todo := range todos {
		if id == todo.ID {
			return todo, nil
		}
	}
	return nil, fmt.Errorf("No todo exists with ID: %d", id)
}

func getIndexById(id int) int {
	for idx, todo := range todos {
		if id == todo.ID {
			return idx
		}
	}
	return -1
}

func AddTodo(t *Todo) (*Todo, error) {
	t.ID = len(todos) + 1
	todos = append(todos, t)
	return t, nil
}

func UpdateTodo(id int, t *Todo) error {
	todo, _ := FindByID(id)
	todo.Title = t.Title
	todo.Body = t.Body
	todo.Time = t.Time
	return nil
}

func DeleteTodo(id int) error {
	idx := getIndexById(id)
	if idx == -1 {
		return fmt.Errorf("Could not find a todo by ID: %d", id)
	}
	todos = append(todos[:idx], todos[idx+1:]...)
	for _, todo := range todos[idx:] {
		todo.ID = id
		id++
	}
	return nil
}
