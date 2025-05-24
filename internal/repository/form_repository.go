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
		AccountLimit:      body.AccountLimit,
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

	postgres.DB.Raw("SELECT id, participants_count, participants_limit, account_limit, title FROM forms WHERE public_id = ?", publicId).Scan(&form)

	var fields []models.FieldResponse
	postgres.DB.Raw("SELECT name, type FROM fields WHERE form_id = ?", form.ID).Scan(&fields)
	if len(fields) == 0 {
		fields = []models.FieldResponse{}
	}

	var varieties []models.Variety
	postgres.DB.Raw("SELECT id, title, cover_url, price FROM varieties WHERE form_id = ?", form.ID).Scan(&varieties)
	if len(varieties) == 0 {
		varieties = []models.Variety{}
	}

	var layout models.LayoutResponse
	postgres.DB.Raw("SELECT title, type, schema, zones FROM layouts WHERE form_id = ?", form.ID).Scan(&layout)

	response := models.FormResponse{
		ID:           form.ID,
		Title:        form.Title,
		IsFull:       false,
		AccountLimit: form.AccountLimit,
		Fields:       fields,
		Varieties:    varieties,
		Layout:       layout,
	}

	if form.ParticipantsCount >= form.ParticipantsLimit {
		response.IsFull = true
	}

	return response
}

func CheckLimit(formId int, userId int) bool {
	var accountLimit int
	var userTickets int

	// Получаем лимит из таблицы forms
	postgres.DB.Raw("SELECT account_limit FROM forms WHERE id = ?", formId).Scan(&accountLimit)

	// Считаем количество билетов пользователя на данную форму
	postgres.DB.Raw("SELECT COUNT(*) FROM tickets WHERE form_id = ? AND user_id = ?", formId, userId).Scan(&userTickets)

	// Проверяем, достигнут ли лимит
	return userTickets >= accountLimit
}

func GetFormById(id int) models.Form {
	var form models.Form
	postgres.DB.First(&form, id)

	return form
}

func GetMyProjects(userId string) models.MyProjectsResponse {
	var projects models.MyProjectsResponse
	postgres.DB.Raw("SELECT public_id, title, participants_count, participants_limit FROM forms WHERE user_id = ?", userId).Scan(&projects.Forms)

	return projects
}
