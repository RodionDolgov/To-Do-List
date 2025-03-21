package handlers

import (
	"context"
	"todo-api/database"
	"todo-api/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := database.DB.Exec(context.Background(), "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)", task.Title, task.Description, task.Status)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(task)
}

func GetTasks(c *fiber.Ctx) error {
	rows, err := database.DB.Query(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := database.DB.Exec(context.Background(), "UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4", task.Title, task.Description, task.Status, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DB.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(204).SendString("Task deleted")
}
