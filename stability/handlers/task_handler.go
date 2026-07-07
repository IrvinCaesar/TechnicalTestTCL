package handlers

import (
	"strconv"
	"strings"

	"stability-test-task-api/models"
	"stability-test-task-api/store"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	tasks := store.GetAllTasks()
	return c.Status(200).JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, _ := strconv.Atoi(idParam)

	task := store.GetTaskByID(id)

	if task == nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	return c.Status(200).JSON(task)
}

func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if strings.TrimSpace(task.Title) == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "title cannot be empty",
		})
	}

	task = store.AddTask(task)
	return c.Status(201).JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	var task models.Task

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid task id",
		})
	}

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if strings.TrimSpace(task.Title) == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "title cannot be empty",
		})
	}

	task.ID = id

	updatedTask := store.UpdateTask(id, task)

	if updatedTask == nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	return c.Status(200).JSON(updatedTask)
}

func DeleteTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, _ := strconv.Atoi(idParam)
	deletedTask := store.GetTaskByID(id)

	if deletedTask == nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	store.DeleteTask(id)

	return c.Status(200).JSON(fiber.Map{
		"message": "deleted",
	})
}
