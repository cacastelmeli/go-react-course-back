package application_todo

import domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"

type TodoCreator struct {
	Repo domain_todo.TodoRepository
}

func (t *TodoCreator) Create(todo *domain_todo.Todo) {
	t.Repo.Add(todo)
}
