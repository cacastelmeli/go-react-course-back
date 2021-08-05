package server

import (
	infrastructure_todo "github.com/cacastelmeli/go-todo-back/api/todo/infrastructure"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupApiRoutes(router *fiber.App) {
	rootGroup := router.Group("/api")

	infrastructure_todo.SetupTodoRoutes(rootGroup)
}

func setupServer() *fiber.App {
	server := fiber.New()

	server.Use(cors.New())

	setupApiRoutes(server)

	return server
}

func StartServer() {
	server := setupServer()
	server.Listen(":8080")
}
