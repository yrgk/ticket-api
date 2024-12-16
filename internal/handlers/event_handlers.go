package handlers

import (
	"ticket-api/pkg/repository"

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

	return c.JSON(event)
}