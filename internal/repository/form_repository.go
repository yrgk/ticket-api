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

	// Adding form id to every field
	for i := range body.Fields {
		body.Fields[i].FormId = form.ID
	}

	// Creating field records in DB
	if err := postgres.DB.Create(&body.Fields).Error; err != nil {
		return err
	}

	return nil
}

func GetForm(publicId string) models.FormResponse {
	var data struct {
		Id    uint   `json:"id"`
		Title string `json:"title"`
	}

	postgres.DB.Raw("SELECT id, title FROM forms WHERE public_id = ?", publicId).Scan(&data)

	var fields []models.FieldResponse
	postgres.DB.Raw("SELECT name, type FROM fields WHERE form_id = ?", data.Id).Scan(&fields)

	response := models.FormResponse{
		ID: data.Id,
		Title:  data.Title,
		Fields: fields,
	}

	return response
}

func GetMyProjects(userId string) models.MyProjectsResponse {
	var projects models.MyProjectsResponse
	postgres.DB.Raw("SELECT public_id, title FROM forms WHERE user_id = ?", userId).Scan(&projects.Forms)

	return projects
}
