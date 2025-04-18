package handlers

import (
	"fmt"
	"ticket-api/internal/models"
	"ticket-api/internal/repository"
	"ticket-api/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateFormHandler(c *fiber.Ctx) error {
	var body models.Form
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	body.FormId = utils.GetMD5Hash(fmt.Sprintf("%s%d%v", body.Title, body.UserId, time.Now()))

	if err := repository.CreateForm(body); err != nil {
		return c.SendStatus(fiber.StatusConflict)
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetFormHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON("Invalid id")
	}

	form := repository.GetForm(id)
	if form.Title == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(form)
}