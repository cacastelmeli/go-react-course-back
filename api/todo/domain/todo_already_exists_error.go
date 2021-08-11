package domain_todo

import "fmt"

type TodoAlreadyExistsError struct {
	Id TodoId
}

func NewTodoAlreadyExistsError(id TodoId) error {
	return TodoAlreadyExistsError{id}
}

func (t TodoAlreadyExistsError) Error() string {
	return fmt.Sprintf("todo with id `%f` already exists", t.Id)
}
