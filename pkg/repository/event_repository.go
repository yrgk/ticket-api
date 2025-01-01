package repository

import (
	"ticket-api/internal/models"
	"ticket-api/pkg/mongo"
	"ticket-api/pkg/postgres"
)

func GetEvent(id int) models.Event {
	var event models.Event
	postgres.DB.Raw("SELECT title, description, is_paid, cover_url, participants_count, is_double_verify, organizator_id FROM events WHERE id = ?", id).Scan(&event)

	return event
}

func CreateEvent(body models.CreateEventRequest) error {
	event := models.Event{
		Title: body.Title,
		Description: body.Description,
		IsPaid: body.IsPaid,
		CoverUrl: body.CoverUrl,
		Capacity: body.Capacity,
		IsDoubleVerify: body.IsDoubleVerify,
		OrganizatorId: body.OrganizatorId,
	}

	if err := postgres.DB.Create(&event).Error; err != nil {
		return err
	}

	return nil
}

func UploadForm(fields []interface{}) error {
	if _, err := mongo.Collection.InsertMany(mongo.Ctx, fields); err != nil {
		return err
	}

	return nil
}