package handlers

import (
	"fmt"
	"strings"
	"ticket-api/config"
	"ticket-api/internal/models"
	"ticket-api/pkg/repository"
	"ticket-api/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TakeTicketHandler(c *fiber.Ctx) error {
	var body models.TakeTicketRequest

	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	ticketId := utils.GetMD5Hash(fmt.Sprintf("%d%d%s%v", body.UserId, body.EventId, body.Variety, time.Now()))

	objectName := utils.GetMD5Hash(fmt.Sprintf("%s%s", ticketId, config.Config.Password))

	qrCode, err := repository.CreateQrCode(body, ticketId, objectName)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	qrCodeUrl := strings.Trim(string(qrCode), "\"")

	ticketBody := models.Ticket{
		QrCodeUrl: qrCodeUrl,
		UserId:    body.UserId,
		EventId:   body.EventId,
		TicketId:  ticketId,
	}

	if err := repository.TakeTicket(ticketBody); err != nil {
		return c.Status(fiber.StatusConflict).SendString("event is full")
	}

	// ADD FORM DATA TO CLICKHOUSE
	if err := repository.UploadUserData(ticketBody, body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	// SEND MESSAGE FROM BOT ABOUT TAKING A TICKET
	// if err := bot.SendTicketInChat(body.UserId, ticketId); err != nil {
	// 	return c.Status(fiber.StatusConflict).SendString("bot does not take a message")
	// }

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
	// Req exmp: https://t.me/botbot/event?startapp=ghfdljkh34tn87cogkhjgfd908kj   ->   extract: ghfdljkh34tn87cogkhjgfd908kj
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
