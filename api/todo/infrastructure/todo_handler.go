package infrastructure_todo

import (
	"net/http"
	"os"
	"strconv"

	application_todo "github.com/cacastelmeli/go-todo-back/api/todo/application"
	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
	"github.com/gofiber/fiber/v2"
)

var mockMode = os.Getenv("MOCK") == "true"

type TodoHandler struct {
	creator  *application_todo.TodoCreator
	searcher *application_todo.TodoSearcher
	remover  *application_todo.TodoRemover
	updater  *application_todo.TodoUpdater
}

func NewTodoHandler() *TodoHandler {
	var repo domain_todo.TodoRepository

	if mockMode {
		repo = NewInMemoryTodoRepository([]*domain_todo.Todo{})
	} else {
		repo = NewMongoTodoRepository()
	}

	creator := application_todo.NewTodoCreator(repo)
	searcher := application_todo.NewTodoSearcher(repo)
	remover := application_todo.NewTodoRemover(repo)
	updater := application_todo.NewTodoUpdater(repo)

	return &TodoHandler{
		creator:  creator,
		searcher: searcher,
		remover:  remover,
		updater:  updater,
	}
}

func (t *TodoHandler) GetAll(c *fiber.Ctx) error {
	return c.JSON(t.searcher.GetAll())
}

func (t *TodoHandler) Add(c *fiber.Ctx) error {
	todo := &domain_todo.Todo{}

	c.BodyParser(todo)
	err := t.creator.Create(todo)

	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

func (t *TodoHandler) Update(c *fiber.Ctx) error {
	todo := &domain_todo.Todo{}

	c.BodyParser(todo)
	err := t.updater.Update(todo)

	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

func (t *TodoHandler) Remove(c *fiber.Ctx) error {
	id := c.Params("id")
	iId, err := strconv.ParseFloat(id, 64)

	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	err = t.remover.Remove(domain_todo.TodoId(iId))

	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.
		Status(http.StatusOK).
		JSON(fiber.Map{"success": true})
}
