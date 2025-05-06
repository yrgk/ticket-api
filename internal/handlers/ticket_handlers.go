package handlers

import (
	"fmt"
	"ticket-api/internal/bot"
	"ticket-api/internal/models"
	"ticket-api/internal/repository"
	"ticket-api/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TakeTicketHandler(c *fiber.Ctx) error {
	var body models.TakeTicketRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	// Making a public id for ticket
	ticketId := utils.GetMD5Hash(fmt.Sprintf("%d%d%s%v", body.UserId, body.FormId, body.Variety, time.Now()))

	// Making an object name (key) for S3
	// objectName := utils.GetMD5Hash(fmt.Sprintf("%s%s", ticketId, config.Config.Password))

	// Making a QR-code url in S3 (content of qr code: https://t.me/botname/app?startapp=check=ticketId)
	qrCode, err := utils.CreateQrCode(body, ticketId)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	ticketBody := models.Ticket{
		QrCodeUrl: qrCode,
		UserId:    body.UserId,
		FormId:    body.FormId,
		TicketId:  ticketId,
	}

	if err := repository.TakeTicket(ticketBody); err != nil {
		return c.Status(fiber.StatusConflict).SendString("form is full")
	}

	// ADD FORM DATA TO CLICKHOUSE
	if err := repository.UploadUserData(ticketBody, body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	// SEND MESSAGE FROM BOT ABOUT TAKING A TICKET
	if err := bot.SendTicketInChat(body.UserId, ticketId); err != nil {
		return c.Status(fiber.StatusConflict).SendString("bot does not sent a message")
	}

	return c.SendStatus(fiber.StatusOK)
}

// WITH USER ID AUTH
func GetTicketHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	userId := c.QueryInt("user_id")

	ticket := repository.GetTicket(id, userId)
	if ticket.Title == "" {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.JSON(ticket)
}

func CheckTicketHandler(c *fiber.Ctx) error {
	ticketId := c.Params("ticket_id")

	validatorId := c.QueryInt("validator_id")
	if validatorId == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("validator id have not found")
	}

	ticket := repository.CheckTicket(ticketId, validatorId)

	return c.JSON(ticket)
}

func ValidateTicketHandler(c *fiber.Ctx) error {
	ticketId := c.Params("ticket_id")
	validatorId := c.Params("validator_id")

	if err := repository.ValidateTicket(ticketId, validatorId); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetMyTicketsHandler(c *fiber.Ctx) error {
	userId := c.QueryInt("user_id")

	tickets := repository.GetMyTickets(userId)

	return c.JSON(tickets)
}

func PayTicketHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
