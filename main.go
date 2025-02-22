package main

import (
	"github.com/gofiber/fiber/v2"
	"todo-task-go/database"
	"todo-task-go/handlers"
)

func main() {

	app := fiber.New()

	database.InitDatabase("postgres://postgres:12345@localhost:5432/postgres")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Todo Task API!")
	})

	// Define a route for /tasks
	app.Get("/tasks", handlers.GetTasks)
	app.Post("/tasks", handlers.CreateTask)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)

	app.Listen(":3000")
}
