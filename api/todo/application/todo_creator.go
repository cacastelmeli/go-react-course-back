package application_todo

import domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"

type TodoCreator struct {
	Repo   domain_todo.TodoRepository
	finder *domain_todo.TodoFinder
}

func NewTodoCreator(repo domain_todo.TodoRepository) *TodoCreator {
	return &TodoCreator{
		Repo:   repo,
		finder: &domain_todo.TodoFinder{Repo: repo},
	}
}

func (t *TodoCreator) Create(todo *domain_todo.Todo) error {
	foundTodo, _ := t.finder.Find(todo.Id)

	if foundTodo != nil {
		return domain_todo.NewTodoAlreadyExistsError(todo.Id)
	}

	t.Repo.Add(todo)
	return nil
}
