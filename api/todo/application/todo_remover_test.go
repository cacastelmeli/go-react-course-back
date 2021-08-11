package application_todo_test

import (
	"testing"

	application_todo "github.com/cacastelmeli/go-todo-back/api/todo/application"
	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
	insfrastructure_todo "github.com/cacastelmeli/go-todo-back/api/todo/infrastructure"
	"github.com/stretchr/testify/assert"
)

var fakeTodo = &domain_todo.Todo{
	Id:   0,
	Text: "fake todo",
}

var inMemoryRepoRemove = insfrastructure_todo.NewInMemoryTodoRepository(
	[]*domain_todo.Todo{fakeTodo},
)
var todoRemover = application_todo.NewTodoRemover(inMemoryRepoRemove)

func TestTodoRemover_Remove(t *testing.T) {
	// Assert initial state
	assert.ElementsMatch(t, inMemoryRepoRemove.Todos, []*domain_todo.Todo{fakeTodo})

	err := todoRemover.Remove(0)

	// Assert non-error removing
	assert.Nil(t, err)

	// Assert empty todos
	assert.Empty(t, inMemoryRepoRemove.Todos)

	// Subsequent removes should throw error
	err = todoRemover.Remove(0)

	// Assert not nil error
	assert.EqualError(t, err, domain_todo.NewTodoNotFoundError(0).Error())
}
