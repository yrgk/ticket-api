package handlers

import "github.com/gofiber/fiber/v2"

func ExportDataToExcel(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	formId := c.Params("id")

	return c.JSON(userId, formId)
}