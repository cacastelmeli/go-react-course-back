package application_todo

import domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"

type TodoUpdater struct {
	Repo domain_todo.TodoRepository
}

func (t *TodoUpdater) Update(todo *domain_todo.Todo) {
	t.Repo.Update(todo)
}
