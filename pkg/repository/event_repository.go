package repository

import (
	"ticket-api/internal/models"
	"ticket-api/pkg/postgres"
)

func GetEvent(id int) models.Event {
	var event models.Event
	postgres.DB.Raw("SELECT title, description, is_paid, cover_url, participants_count, is_double_verify, organizator_id FROM events WHERE id = ?", id).Scan(&event)

	return event
}