package repository

import (
	"fmt"
	"ticket-api/internal/models"
	"ticket-api/internal/utils"
	"ticket-api/pkg/postgres"
	"time"
)

func CreateForm(body models.FormRequest) error {
	// Making public id for linking
	publicId := utils.GetMD5Hash(fmt.Sprintf("%s%d%v", body.Title, body.UserId, time.Now()))

	form := models.Form{
		Title:             body.Title,
		PublicId:          publicId,
		ParticipantsLimit: body.ParticipantsLimit,
		UserId:            body.UserId,
	}

	// Creating a form record in DB
	if err := postgres.DB.Create(&form).Error; err != nil {
		return err
	}

	// Check if there is no fields
	if len(body.Fields) != 0 {
		// Adding form id to every field
		for i := range body.Fields {
			body.Fields[i].FormId = form.ID
		}

		// Creating field records in DB
		if err := postgres.DB.Create(&body.Fields).Error; err != nil {
			return err
		}
	}

	return nil
}

func GetForm(publicId string) models.FormResponse {
	var form models.Form

	postgres.DB.Raw("SELECT id, participants_count, participants_limit, title FROM forms WHERE public_id = ?", publicId).Scan(&form)

	var fields []models.FieldResponse
	postgres.DB.Raw("SELECT name, type FROM fields WHERE form_id = ?", form.ID).Scan(&fields)
	if len(fields) == 0 {
		fields = []models.FieldResponse{}
	}

	var varieties []models.Variety
	postgres.DB.Raw("SELECT title, cover_url, price FROM varieties WHERE form_id = ?", form.ID).Scan(&varieties)

	var hall models.Layout
	postgres.DB.Raw("SELECT * FROM layouts WHERE form_id = ?", form.ID).Scan(&hall)

	response := models.FormResponse{
		ID:        form.ID,
		Title:     form.Title,
		IsFull:    false,
		Fields:    fields,
		Varieties: varieties,
		Hall:      hall,
	}

	if form.ParticipantsCount >= form.ParticipantsLimit {
		response.IsFull = true
	}

	return response
}

func GetMyProjects(userId string) models.MyProjectsResponse {
	var projects models.MyProjectsResponse
	postgres.DB.Raw("SELECT public_id, title, participants_count, participants_limit FROM forms WHERE user_id = ?", userId).Scan(&projects.Forms)

	return projects
}
