package application_todo

import domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"

type TodoSearcher struct {
	Repo domain_todo.TodoRepository
}

func (t *TodoSearcher) GetAll() []*domain_todo.Todo {
	return t.Repo.GetAll()
}
