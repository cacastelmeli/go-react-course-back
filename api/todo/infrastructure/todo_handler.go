package infrastructure_todo

import (
	"strconv"

	application_todo "github.com/cacastelmeli/go-todo-back/api/todo/application"
	domain_todo "github.com/cacastelmeli/go-todo-back/api/todo/domain"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	creator  *application_todo.TodoCreator
	searcher *application_todo.TodoSearcher
	remover  *application_todo.TodoRemover
	updater  *application_todo.TodoUpdater
}

func NewTodoHandler() *TodoHandler {
	repo := &InMemoryTodoRepository{}
	creator := &application_todo.TodoCreator{
		Repo: repo,
	}

	searcher := &application_todo.TodoSearcher{
		Repo: repo,
	}

	remover := &application_todo.TodoRemover{
		Repo: repo,
	}

	updater := &application_todo.TodoUpdater{
		Repo: repo,
	}

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
	t.creator.Create(todo)

	return c.JSON(fiber.Map{"success": true})
}

func (t *TodoHandler) Update(c *fiber.Ctx) error {
	todo := &domain_todo.Todo{}

	c.BodyParser(todo)
	t.updater.Update(todo)

	return c.JSON(fiber.Map{"success": true})
}

func (t *TodoHandler) Remove(c *fiber.Ctx) error {
	id := c.Params("id")
	iId, err := strconv.ParseFloat(id, 64)

	if err != nil {
		return err
	}

	t.remover.Remove(iId)

	return c.JSON(fiber.Map{"success": true})
}
