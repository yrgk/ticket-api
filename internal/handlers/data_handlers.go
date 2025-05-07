package handlers

import (
	"fmt"
	"ticket-api/internal/repository"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ExportDataToExcelHandler(c *fiber.Ctx) error {
	// Get params
	userId := c.QueryInt("user_id")
	formId := c.Params("id")

	// Checking if params are invalid
	if formId == "" || userId == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get userdata
	userData := repository.GetUserdata(userId, formId)
	if len(userData) == 0 {
		return c.Status(fiber.StatusNotFound).SendString("user data is empty")
	}

	// Parse into excel file
	file, err := repository.ParseToExcel(userData)
	if err != nil {
		return c.SendStatus(fiber.StatusConflict)
	}

	// Setting a filename
	filename := fmt.Sprintf("vellem_%d.xlsx", time.Now().Unix())

	// Adding headers
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Return downloading of excel file
	return c.SendStream(file)
}