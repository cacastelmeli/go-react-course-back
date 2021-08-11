package infrastructure_todo_test

import (
	"context"
	"log"
	"testing"

	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
	infrastructure_todo "github.com/cacastelmeli/go-todo-back/api/todo/infrastructure"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

var mongoTodoRepository = infrastructure_todo.NewMongoTodoRepository()
var mongoTodoCollection = infrastructure_todo.MongoTodoCollection

func TestTodoRepository(t *testing.T) {
	clearCollection(t)

	insertedTodo := &domain_todo.Todo{
		Id:   1.234,
		Text: "Inserted todo",
	}

	mongoTodoRepository.Add(insertedTodo)

	// Assert todos inserted
	assert.ElementsMatch(t, getTodosFromCollection(), []*domain_todo.Todo{insertedTodo})
	// assert.ElementsMatch(t, []*domain_todo.Todo{insertedTodo}, mongoTodoRepository.GetAll())

	updatedTodo := &domain_todo.Todo{
		Id:   1.234,
		Text: "Updated todo",
		Done: true,
	}

	mongoTodoRepository.Update(updatedTodo)

	// Assert todos updated
	assert.ElementsMatch(t, getTodosFromCollection(), []*domain_todo.Todo{updatedTodo})
	assert.ElementsMatch(t, []*domain_todo.Todo{updatedTodo}, mongoTodoRepository.GetAll())

	mongoTodoRepository.Remove(1.234)

	// Assert todos removed
	assert.Empty(t, getTodosFromCollection())
	assert.Empty(t, mongoTodoRepository.GetAll())
}

func clearCollection(t *testing.T) {
	err := mongoTodoCollection.Drop(context.TODO())

	assert.Nil(t, err)
}

func getTodosFromCollection() []*domain_todo.Todo {
	todos := []*domain_todo.Todo{}
	cursor, err := mongoTodoCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Println(err)
		return nil
	}

	cursor.All(context.Background(), &todos)
	return todos
}
