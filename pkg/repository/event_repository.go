package repository

import (
	"ticket-api/internal/models"
	"ticket-api/pkg/postgres"
)

func GetEvent(id int) models.Event {
	var event models.Event
	postgres.DB.Raw("SELECT * FROM events WHERE id = ?", id).Scan(&event)

	return event
}

func CreateEvent(body models.CreateEventRequest) error {
	event := models.Event{
		Title:       body.Title,
		Description: body.Description,
		PageData:    body.PageData,
		IsPaid:      body.IsPaid,
		CoverUrl:    body.CoverUrl,
		BasePrice:   body.BasePrice,
		Capacity:    body.Capacity,
		StartTime:   body.StartTime,
		Duration:    (body.Duration),
		OrganizatorId: body.OrganizatorId,
		// IsDoubleVerify: body.IsDoubleVerify,
	}

	if err := postgres.DB.Create(&event).Error; err != nil {
		return err
	}

	for i := range body.FormData {
		body.FormData[i].EventId = event.ID
	}

	if err := postgres.DB.Create(&body.FormData).Error; err != nil {
		return err
	}

	return nil
}

func GetForm(id int) []models.Field {
	var fields []models.Field
	postgres.DB.Raw("SELECT name, type FROM fields WHERE event_id = ?", id).Scan(&fields)

	return fields
}
