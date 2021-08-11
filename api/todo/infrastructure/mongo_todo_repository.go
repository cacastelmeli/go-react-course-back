package infrastructure_todo

import (
	"context"
	"log"

	infrastructure_shared "github.com/cacastelmeli/go-todo-back/api/shared/infrastructure"
	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTodoRepository struct {
	collection *mongo.Collection
}

func NewMongoTodoRepository() *MongoTodoRepository {
	collection := infrastructure_shared.MongoDatabase.Collection("todo")

	return &MongoTodoRepository{
		collection,
	}
}

func (m *MongoTodoRepository) Add(todo *domain_todo.Todo) {
	todoBson := bson.M{
		"id":   todo.Id,
		"text": todo.Text,
		"done": todo.Done,
	}

	_, err := m.collection.InsertOne(context.Background(), todoBson)

	if err != nil {
		log.Println("Error while saving todo", err)
	}
}

func (m *MongoTodoRepository) GetAll() []*domain_todo.Todo {
	cursor, err := m.collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Println("Error while getting all todos", err)
		return nil
	}

	allTodos := []*domain_todo.Todo{}
	err = cursor.All(context.Background(), &allTodos)

	if err != nil {
		log.Println("Error while getting all todos", err)
		return nil
	}

	return allTodos
}

func (m *MongoTodoRepository) Remove(id domain_todo.TodoId) {
	_, err := m.collection.DeleteOne(context.TODO(), bson.D{{"id", id}})

	if err != nil {
		log.Println("Error while removing todo", err)
	}
}

func (m *MongoTodoRepository) Update(todo *domain_todo.Todo) {
	filter := bson.D{{"id", todo.Id}}
	update := bson.D{
		{
			"$set",
			bson.D{
				{"text", todo.Text},
				{"done", todo.Done},
			},
		},
	}

	_, err := m.collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Println("Error while updating todo", err)
	}
}

func (m *MongoTodoRepository) Find(id domain_todo.TodoId) *domain_todo.Todo {
	filter := bson.D{{"id", id}}
	todo := &domain_todo.Todo{}

	err := m.collection.FindOne(context.TODO(), filter).Decode(todo)

	if err != nil {
		return nil
	}

	return todo
}
