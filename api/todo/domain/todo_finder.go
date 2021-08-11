package domain_todo

type TodoFinder struct {
	Repo TodoRepository
}

func (t *TodoFinder) Find(id TodoId) (*Todo, error) {
	todo := t.Repo.Find(id)

	if todo == nil {
		return nil, NewTodoNotFoundError(id)
	}

	return todo, nil
}
