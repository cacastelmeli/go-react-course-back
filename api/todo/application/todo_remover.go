package application_todo

import domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"

type TodoRemover struct {
	Repo domain_todo.TodoRepository
}

func (t *TodoRemover) Remove(id float64) {
	t.Repo.Remove(id)
}
