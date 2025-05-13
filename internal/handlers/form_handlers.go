package handlers

import (
	"ticket-api/internal/models"
	"ticket-api/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func CreateFormHandler(c *fiber.Ctx) error {
	var body models.FormRequest
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// TODO: add layouts and varieties

	if err := repository.CreateForm(body); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetFormHandler(c *fiber.Ctx) error {
	// Public form id
	id := c.Params("id")

	form := repository.GetForm(id)
	if form.Title == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(form)
}

func GetMyProjectsHandler(c *fiber.Ctx) error {
	userId := c.Query("user_id")

	projects := repository.GetMyProjects(userId)
	if len(projects.Forms) == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(projects)
}
