package handlers

import (
	"strings"
	"ticket-api/internal/models"
	"ticket-api/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

func TakeTicketHandler(c *fiber.Ctx) error {
	var body models.TakeTicketRequest

	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	qrCode, err := repository.CreateQrCode(body)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	qrCodeUrl := strings.Trim(string(qrCode), "\"")

	ticketBody := models.Ticket{
		QrCodeUrl: string(qrCodeUrl),
		UserId:    body.UserId,
		EventId:   body.EventId,
	}

	if err := repository.TakeTicket(ticketBody); err != nil {
		return c.Status(fiber.StatusConflict).SendString("event is full")
	}

	return c.SendStatus(fiber.StatusOK)
}
