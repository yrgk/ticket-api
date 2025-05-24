package repository

import (
	"bytes"
	"context"
	"crypto/sha256"
	"time"

	"encoding/hex"
	// "encoding/json"
	"errors"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"math/rand"

	// "time"

	"image"
	// "image/color"
	// "image/draw"
	// "math/rand"

	"github.com/disintegration/gift"

	"ticket-api/config"
	"ticket-api/internal/models"
	"ticket-api/internal/utils"
	"ticket-api/pkg/postgres"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/fogleman/gg"
)

func TakeTicket(body models.Ticket) error {
	tx := postgres.DB.Begin()

	// Trying to increase particicpants count
	result := tx.Exec(`
		UPDATE forms
		SET participants_count = participants_count + 1
		WHERE id = ? AND participants_count < participants_limit
	`, body.FormId)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("limit reached for this form")
	}

	// Creating ticket
	if err := tx.Create(&body).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func GetTicket(id string, userId int) models.TicketResponse {
	var ticket models.TicketResponse
	postgres.DB.Raw("SELECT f.title, t.qr_code_url, t.cover_url, t.ticket_number, f.participants_limit, t.form_id, t.variety, t.is_activated FROM forms f, tickets t WHERE t.ticket_id = ? AND t.user_id = ? AND f.id = t.form_id;", id, userId).Scan(&ticket)

	var variety models.Variety
	postgres.DB.Raw("SELECT * FROM varieties WHERE form_id = ?", ticket.ID)

	ticket.Variety = variety

	return ticket
}

func GetTicketForChecking(id string, userId int) models.TicketCheckResponse {
	var ticket models.TicketCheckResponse
	postgres.DB.Raw("SELECT variety, is_activated, user_id FROM tickets WHERE ticket_id = ?", id).Scan(&ticket)

	return ticket
}

func CheckTicket(ticketId string, validatorId int) (models.TicketCheckResponse, error) {
	var ticket models.Ticket
	postgres.DB.Raw("SELECT form_id, is_activated, variety_id FROM tickets WHERE ticket_id = ?", ticketId).Scan(&ticket)

	var form models.Form
	postgres.DB.Raw("SELECT title, user_id FROM forms WHERE id = ?", ticket.FormId).Scan(&form)

	// Checking if validator is real
	if form.UserId != validatorId {
		return models.TicketCheckResponse{}, errors.New("forbidden")
	}

	response := models.TicketCheckResponse{
		Title:       form.Title,
		IsActivated: ticket.IsActivated,
		VarietyId:   ticket.VarietyId,
	}

	return response, nil
}

// ticketId is Public
func ValidateTicket(ticketId string, userId int) error {
	var ticket models.Ticket
	postgres.DB.Raw("SELECT form_id FROM tickets WHERE ticket_id = ?", ticketId).Scan(&ticket)

	var form models.Form
	postgres.DB.Raw("SELECT user_id FROM forms WHERE id = ?", ticket.FormId).Scan(&form)

	// Checking if validator id is valid
	if userId != form.UserId {
		return errors.New("incorrect validator id")
	}

	// Changing status of Ticket and deleting QR-code link
	postgres.DB.Exec("UPDATE tickets SET qr_code_url = NULL, is_activated = TRUE WHERE ticket_id = ?", ticketId)

	// Updating data in clickhouse

	// Deleting qr code from s3
	if err := utils.DeleteFromS3(ticketId); err != nil {
		log.Println("ERROR DELETING QR CODE FROM S3", ticketId, err)
	}

	return nil
}

func GetMyTickets(id int) []models.MyTicketResponse {
	var tickets []models.MyTicketResponse
	postgres.DB.Raw("SELECT t.variety_id, t.ticket_id, t.is_activated, t.cover_url, f.title FROM tickets t JOIN forms f ON t.form_id = f.id WHERE t.user_id = ?", id).Scan(&tickets)

	return tickets
}

func CheckValidator(eventId, validatorId int) bool {
	var validator models.Validator
	postgres.DB.Raw("SELECT * FROM validators WHERE event_id = ? AND validator_id = ?", eventId, validatorId).Scan(&validator)

	return validator.EventId != 0
}

// func UploadUserData(body models.Ticket, ticketBody models.TakeTicketRequest) error {
// 	formData, err := json.Marshal(ticketBody.FormData)
// 	if err != nil {
// 		return err
// 	}

// 	data := models.TicketMeta{
// 		UserId:   body.UserId,
// 		FormId:   body.FormId,
// 		TicketId: body.TicketId,
// 		UserData:   formData,
// 	}

// 	if err := postgres.DB.Create(&data).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

const (
	width     = 512
	height    = 512
	blobCount = 5
)

func hexToColor(hexStr string) color.Color {
	r, _ := hex.DecodeString(hexStr[0:2])
	g, _ := hex.DecodeString(hexStr[2:4])
	b, _ := hex.DecodeString(hexStr[4:6])
	return color.RGBA{r[0], g[0], b[0], 180}
}

func getColorsFromHash(hash string, count int) []color.Color {
	var colors []color.Color
	for i := 0; i < count; i++ {
		start := i * 6
		colors = append(colors, hexToColor(hash[start:start+6]))
	}
	return colors
}

func MakeCover(seed, ticketId string) (string, error) {
	// Генерация хэша и цветов
	hash := sha256.Sum256([]byte(seed))
	hashStr := hex.EncodeToString(hash[:])
	colors := getColorsFromHash(hashStr, blobCount)
	rng := rand.New(rand.NewSource(int64(hash[0])))

	// Создаем canvas
	dc := gg.NewContext(width, height)
	// dc.SetRGB(0.1, 0.1, 0.2) // фон
	dc.SetRGB(0.7, 0.7, 0.7) // фон
	dc.Clear()

	for _, c := range colors {
		r := rng.Float64()*300 + 200
		x := rng.Float64() * width
		y := rng.Float64() * height
		grad := gg.NewRadialGradient(x, y, 0, x, y, r)
		rgba := c.(color.RGBA)
		grad.AddColorStop(0, color.RGBA{rgba.R, rgba.G, rgba.B, 200})
		grad.AddColorStop(1, color.RGBA{rgba.R, rgba.G, rgba.B, 0})
		dc.SetFillStyle(grad)
		dc.DrawCircle(x, y, r)
		dc.Fill()
	}

	// Преобразуем в image.Image и применяем размытие
	src := dc.Image()
	dst := image.NewRGBA(src.Bounds())

	g := gift.New(
		gift.GaussianBlur(100.0), // Уровень размытия
	)
	g.Draw(dst, src)

	// Кодируем в PNG
	var buf bytes.Buffer
	err := png.Encode(&buf, dst)
	if err != nil {
		return "", fmt.Errorf("failed to encode image: %w", err)
	}

	// Получаем S3 клиент
	client, err := utils.GetS3Client()
	if err != nil {
		return "", fmt.Errorf("failed to get s3 client: %w", err)
	}

	s := time.Now()
	// Ключ и загрузка
	key := fmt.Sprintf("covers/%s.png", ticketId)
	_, err = client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(config.Config.BucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String("image/png"),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}

	link := fmt.Sprintf("%s/%s/%s", config.Config.S3ApiUrl, config.Config.BucketName, key)

	fmt.Println("UPLOADING", time.Since(s))
	return link, nil
}
