package repository

import (
	"ticket-api/internal/models"
	"ticket-api/pkg/postgres"
)

func CreateForm(form models.Form) error {
	if err := postgres.DB.Create(&form).Error; err != nil {
		return err
	}

	return nil
}

func GetForm(id int) models.FormResponse {
	var fields []models.FieldResponse
	postgres.DB.Raw("SELECT name, type FROM fields WHERE event_id = ?", id).Scan(&fields)

	var title string
	postgres.DB.Raw("SELECT title FROM events WHERE id = ?", id).Scan(&title)

	response := models.FormResponse{
		Title:  title,
		Fields: fields,
	}

	return response
}
