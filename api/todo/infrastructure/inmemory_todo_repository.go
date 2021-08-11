package infrastructure_todo

import (
	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
)

// Globally managed todos
var todos = []*domain_todo.Todo{}

type InMemoryTodoRepository struct {
}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{}
}

func (r *InMemoryTodoRepository) Add(todo *domain_todo.Todo) {
	todos = append(todos, todo)
}

func (r *InMemoryTodoRepository) GetAll() []*domain_todo.Todo {
	return todos
}

func (r *InMemoryTodoRepository) Remove(id domain_todo.TodoId) {
	foundIndex := r.findIndex(id)
	lastIndex := len(todos) - 1

	if foundIndex == -1 {
		return
	}

	todos[foundIndex] = todos[lastIndex]
	todos = todos[:lastIndex]
}

func (r *InMemoryTodoRepository) Update(todo *domain_todo.Todo) {
	foundTodo := r.Find(todo.Id)

	if foundTodo != nil {
		foundTodo.Done = todo.Done
		foundTodo.Text = todo.Text
	}
}

func (r *InMemoryTodoRepository) Find(id domain_todo.TodoId) *domain_todo.Todo {
	foundIndex := r.findIndex(id)

	if foundIndex == -1 {
		return nil
	}

	return todos[foundIndex]
}

func (r *InMemoryTodoRepository) findIndex(id domain_todo.TodoId) int {
	for i, todo := range todos {
		if todo.Id == id {
			return i
		}
	}

	return -1
}
