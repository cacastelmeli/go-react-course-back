package application_todo_test

import (
	"testing"

	application_todo "github.com/cacastelmeli/go-todo-back/api/todo/application"
	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
	insfrastructure_todo "github.com/cacastelmeli/go-todo-back/api/todo/infrastructure"
	"github.com/stretchr/testify/assert"
)

var fakeTodos = []*domain_todo.Todo{
	{
		Id:   1,
		Text: "Test 1",
	},
	{
		Id:   2,
		Text: "Test 2",
		Done: true,
	},
}

var inMemoryRepoSearch = insfrastructure_todo.NewInMemoryTodoRepository(fakeTodos)
var todoSearcher = application_todo.NewTodoSearcher(inMemoryRepoSearch)

func TestTodoSearcher_Search(t *testing.T) {
	assert.ElementsMatch(t, fakeTodos[:], todoSearcher.GetAll())
}
