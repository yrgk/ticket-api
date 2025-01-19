package handlers

import (
	"ticket-api/internal/models"
	"ticket-api/pkg/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

	form := repository.GetForm(id)

	response := models.EventResponse{
		Event: event,
		Form:  form,
	}

	return c.JSON(response)
}

func CreateEventHandler(c *fiber.Ctx) error {
	var body models.CreateEventRequest
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	validate := validator.New()

	err := validate.Struct(body)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	if err := repository.CreateEvent(body); err != nil {
		return c.Status(fiber.StatusConflict).JSON(err)
	}

	return c.SendStatus(fiber.StatusOK)
}
