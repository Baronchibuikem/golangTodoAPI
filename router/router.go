package router

import (
	"todoapp/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())

	// Todo
	todo := api.Group("/todos")
	todo.Post("/", handler.CreateTodo)
	todo.Get("/", handler.GetTodos)
	todo.Get("/:id", handler.GetTodo)
	todo.Put("/:id", handler.UpdateTodo)
	todo.Delete("/:id", handler.DeleteTodo)
}