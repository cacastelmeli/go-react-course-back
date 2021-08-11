package infrastructure_todo

import (
	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
)

type InMemoryTodoRepository struct {
	// Expose todos for tests
	Todos []*domain_todo.Todo
}

func NewInMemoryTodoRepository(todos []*domain_todo.Todo) *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		Todos: todos,
	}
}

func (r *InMemoryTodoRepository) Add(todo *domain_todo.Todo) {
	r.Todos = append(r.Todos, todo)
}

func (r *InMemoryTodoRepository) GetAll() []*domain_todo.Todo {
	return r.Todos
}

func (r *InMemoryTodoRepository) Remove(id domain_todo.TodoId) {
	foundIndex := r.findIndex(id)
	lastIndex := len(r.Todos) - 1

	if foundIndex == -1 {
		return
	}

	r.Todos[foundIndex] = r.Todos[lastIndex]
	r.Todos = r.Todos[:lastIndex]
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

	return r.Todos[foundIndex]
}

func (r *InMemoryTodoRepository) findIndex(id domain_todo.TodoId) int {
	for i, todo := range r.Todos {
		if todo.Id == id {
			return i
		}
	}

	return -1
}
