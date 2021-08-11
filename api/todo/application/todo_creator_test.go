package application_todo_test

import (
	"testing"

	application_todo "github.com/cacastelmeli/go-todo-back/api/todo/application"
	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
	insfrastructure_todo "github.com/cacastelmeli/go-todo-back/api/todo/infrastructure"
	"github.com/stretchr/testify/assert"
)

var inMemoryRepoCreate = insfrastructure_todo.NewInMemoryTodoRepository([]*domain_todo.Todo{})
var todoCreator = application_todo.NewTodoCreator(inMemoryRepoCreate)

func TestTodoCreator_Create(t *testing.T) {
	todoCreated := &domain_todo.Todo{
		Id:   1,
		Text: "Testing",
	}

	// Assert initial state as empty
	assert.Empty(t, inMemoryRepoCreate.Todos)

	err := todoCreator.Create(todoCreated)

	// Assert non-error
	assert.Nil(t, err)

	// Assert correct insertion
	assert.ElementsMatch(t, inMemoryRepoCreate.Todos, []*domain_todo.Todo{todoCreated})

	// Subsequent creations should throw error
	err = todoCreator.Create(todoCreated)

	// Assert not nil error
	assert.EqualError(t, err, domain_todo.NewTodoAlreadyExistsError(1).Error())
}
