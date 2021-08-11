package application_todo_test

import (
	"testing"

	application_todo "github.com/cacastelmeli/go-todo-back/api/todo/application"
	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
	insfrastructure_todo "github.com/cacastelmeli/go-todo-back/api/todo/infrastructure"
	"github.com/stretchr/testify/assert"
)

var initialTodo = &domain_todo.Todo{
	Id:   10,
	Text: "todo",
}

var inMemoryRepoUpdate = insfrastructure_todo.NewInMemoryTodoRepository(
	[]*domain_todo.Todo{initialTodo},
)
var todoUpdater = application_todo.NewTodoUpdater(inMemoryRepoUpdate)

func TestTodoUpdater_Update(t *testing.T) {
	// Assert initial todos
	assert.ElementsMatch(t, inMemoryRepoUpdate.Todos, []*domain_todo.Todo{initialTodo})

	err := todoUpdater.Update(&domain_todo.Todo{
		Id:   20,
		Text: "Unknown",
		Done: true,
	})

	// Should throw on unknown todo's `id`
	assert.EqualError(t, err, domain_todo.NewTodoNotFoundError(20).Error())

	todoUpdated := &domain_todo.Todo{
		Id:   10,
		Text: "todo updated",
		Done: true,
	}

	todoUpdater.Update(todoUpdated)

	// Assert updated todo
	assert.ElementsMatch(t, inMemoryRepoUpdate.Todos, []*domain_todo.Todo{todoUpdated})
}
