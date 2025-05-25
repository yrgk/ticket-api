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

	isLimit := repository.CheckLimit(body.FormId, body.UserId)

	if isLimit {
		return c.Status(fiber.StatusConflict).SendString("limit for account is reached")
	}

	t := time.Now()
	// Making a public id for ticket
	ticketId := utils.GetMD5Hash(fmt.Sprintf("%d%d%d%v", body.UserId, body.FormId, body.VarietyId, time.Now()))

	// Making a QR-code url in S3 (content of qr code: https://t.me/botname/app?startapp=check=ticketId)
	qrCode, err := utils.CreateQrCode(body, ticketId)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	fmt.Println("QR CODE", time.Since(t))

	form := repository.GetFormById(body.FormId)

	seed := fmt.Sprintf("#%d/%d %s", form.ParticipantsCount+1, form.ParticipantsLimit, ticketId)

	cover, err := repository.MakeCover(seed, ticketId)
	if err != nil {
		return c.SendStatus(fiber.StatusConflict)
	}

	ticketBody := models.Ticket{
		QrCodeUrl:    qrCode,
		CoverUrl:     cover,
		UserId:       body.UserId,
		FormId:       body.FormId,
		TicketNumber: form.ParticipantsCount + 1,
		TicketId:     ticketId,
		VarietyId:    body.VarietyId,
	}

	if err := repository.TakeTicket(ticketBody); err != nil {
		return c.Status(fiber.StatusConflict).SendString("form is full")
	}

	// SEND MESSAGE FROM BOT ABOUT TAKING A TICKET
	if err := bot.SendTicketInChat(body.UserId, body.FormId, ticketId); err != nil {
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

	ticket, err := repository.CheckTicket(ticketId, validatorId)
	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.JSON(ticket)
}

func ValidateTicketHandler(c *fiber.Ctx) error {
	ticketId := c.Params("ticketId")
	validatorId := c.QueryInt("validator_id")

	if err := repository.ValidateTicket(ticketId, validatorId); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
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
