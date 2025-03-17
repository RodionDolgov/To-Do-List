package main

import (
	"todo-api/database"
	"todo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Инициализация базы данных
	if err := database.InitDB("postgres://postgres:postgres@localhost:5432/ToDoList"); err != nil {
		panic(err)
	}

	// Маршруты
	app.Post("/tasks", handlers.CreateTask)
	app.Get("/tasks", handlers.GetTasks)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)

	app.Listen(":3000")
}
