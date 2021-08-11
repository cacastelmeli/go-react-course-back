package application_todo

import domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"

type TodoRemover struct {
	Repo   domain_todo.TodoRepository
	finder *domain_todo.TodoFinder
}

func NewTodoRemover(repo domain_todo.TodoRepository) *TodoRemover {
	return &TodoRemover{
		Repo:   repo,
		finder: &domain_todo.TodoFinder{Repo: repo},
	}
}

func (t *TodoRemover) Remove(id domain_todo.TodoId) error {
	// Verify todo existance
	_, err := t.finder.Find(id)

	if err != nil {
		return err
	}

	t.Repo.Remove(id)
	return nil
}
