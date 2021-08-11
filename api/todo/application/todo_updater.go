package application_todo

import domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"

type TodoUpdater struct {
	Repo   domain_todo.TodoRepository
	finder *domain_todo.TodoFinder
}

func NewTodoUpdater(repo domain_todo.TodoRepository) *TodoUpdater {
	return &TodoUpdater{
		Repo:   repo,
		finder: &domain_todo.TodoFinder{Repo: repo},
	}
}

func (t *TodoUpdater) Update(todo *domain_todo.Todo) error {
	// Verify todo existance
	_, err := t.finder.Find(todo.Id)

	if err != nil {
		return err
	}

	t.Repo.Update(todo)
	return nil
}
