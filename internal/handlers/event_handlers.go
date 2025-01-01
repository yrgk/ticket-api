package handlers

import (
	"ticket-api/internal/models"
	"ticket-api/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

func GetEventHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON("Invalid id")
	}

	event := repository.GetEvent(id)
	if event.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(event)
}

func CreateEventHandler(c *fiber.Ctx) error {
	var body models.CreateEventRequest
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	validate := validator.New()

	// return c.JSON(body.FormData)
	err := validate.Struct(body)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	return c.JSON(body)

	if err := repository.CreateEvent(body); err != nil {
		return c.Status(fiber.StatusConflict).JSON(err)
	}

	return nil
}
