package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"todo-task-go/da
	"todo-task-go/database"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	database.InitDatabase()

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
