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

// WITH USER ID AUTH
func GetTicketHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	userId := c.QueryInt("user_id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	ticket := repository.GetTicket(id, userId)
	if ticket.Title == "" {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.JSON(ticket)
}

func VerifyTicketHandler(c *fiber.Ctx) error {
	// Req exmp: https://t.me/botbot/event?startapp=ghfdljkh34tn87cogkhjgfd908kj   ->   extract: ghfdljkh34tn87cogkhjgfd908kj
	ticketId := c.Params("ticket_id")
	verifierId := c.Params("verifier_id")

	if err := repository.VerifyTicket(ticketId, verifierId); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	return c.SendStatus(fiber.StatusOK)
}
