package infrastructure_todo

import (
	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
)

// Globally managed todos
var todos = []*domain_todo.Todo{}

type InMemoryTodoRepository struct {
}

func (r *InMemoryTodoRepository) Add(todo *domain_todo.Todo) {
	todos = append(todos, todo)
}

func (r *InMemoryTodoRepository) GetAll() []*domain_todo.Todo {
	return todos
}

func (r *InMemoryTodoRepository) Remove(id float64) {
	foundIndex := -1
	lastIndex := len(todos) - 1

	for i, todo := range todos {
		if todo.Id == id {
			foundIndex = i
			break
		}
	}

	todos[foundIndex] = todos[lastIndex]
	todos = todos[:lastIndex]
}

func (r *InMemoryTodoRepository) Update(todo *domain_todo.Todo) {
	var foundTodo *domain_todo.Todo

	for _, t := range todos {
		if t.Id == todo.Id {
			foundTodo = t
			break
		}
	}

	if foundTodo != nil {
		foundTodo.Done = todo.Done
		foundTodo.Text = todo.Text
	}
}
