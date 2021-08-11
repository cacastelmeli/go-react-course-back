package domain_todo

import "fmt"

type TodoNotFoundError struct {
	Id TodoId
}

func NewTodoNotFoundError(id TodoId) error {
	return TodoNotFoundError{id}
}

func (t TodoNotFoundError) Error() string {
	return fmt.Sprintf("todo not found with id `%f`", t.Id)
}
