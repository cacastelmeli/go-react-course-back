package infrastructure_todo

import "github.com/gofiber/fiber/v2"

func SetupTodoRoutes(router fiber.Router) {
	handler := NewTodoHandler()

	router.Post("/todos", handler.Add)
	router.Get("/todos", handler.GetAll)
	router.Patch("/todos", handler.Update)
	router.Delete("/todos/:id", handler.Remove)
}
