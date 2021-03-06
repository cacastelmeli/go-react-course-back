package application_todo_test

import (
	"testing"

	application_todo "github.com/cacastelmeli/go-todo-back/api/todo/application"
	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
	insfrastructure_todo "github.com/cacastelmeli/go-todo-back/api/todo/infrastructure"
	"github.com/stretchr/testify/assert"
)

var inMemoryRepoSearch = insfrastructure_todo.NewInMemoryTodoRepository(
	[]*domain_todo.Todo{
		{
			Id:   1,
			Text: "Test 1",
		},
		{
			Id:   2,
			Text: "Test 2",
			Done: true,
		},
	},
)
var todoSearcher = application_todo.NewTodoSearcher(inMemoryRepoSearch)

func TestTodoSearcher_Search(t *testing.T) {
	// Assert all elements
	assert.ElementsMatch(t, inMemoryRepoSearch.Todos[:], todoSearcher.GetAll())
}
