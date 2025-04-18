package repository

import (
	"encoding/json"
	"ticket-api/internal/models"
	"ticket-api/pkg/postgres"
	"time"
)

func GetEvent(id int) models.Event {
	var event models.Event
	postgres.DB.Raw("SELECT * FROM events WHERE id = ?", id).Scan(&event)

	return event
}

func CreateEvent(body models.CreateEventRequest) error {
	parsedDuration, err := time.ParseDuration(body.Duration)
	if err != nil {
		return err
	}

	pageData, _ := json.Marshal(body.PageData)

	event := models.Event{
		Title:         body.Title,
		Description:   body.Description,
		PageData:      pageData,
		IsPaid:        body.IsPaid,
		CoverUrl:      body.CoverUrl,
		BasePrice:     body.BasePrice,
		Capacity:      body.Capacity,
		StartTime:     body.StartTime,
		Duration:      parsedDuration,
		OrganizatorId: body.OrganizatorId,
	}

	if err := postgres.DB.Create(&event).Error; err != nil {
		return err
	}

	for i := range body.FormData {
		body.FormData[i].FormId = event.ID
	}

	if err := postgres.DB.Create(&body.FormData).Error; err != nil {
		return err
	}

	return nil
}