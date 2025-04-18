package bot

import (
	"fmt"
	"net/http"
	"ticket-api/config"
)

func SendTicketInChat(userId int, ticketId string) error {
	// text := fmt.Sprintf("Билет успешно получен! | ")
	text := "Билет успешно получен!"

	url := fmt.Sprintf("%s?startapp=%s", config.Config.ClientUrl, ticketId)

	req := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s&reply_markup={\"inline_keyboard\":[[{\"text\":\"Открыть билет\",\"web_app\":{\"url\":\"%s\"}}]]}", config.Config.BotToken, userId, text, url)

	// fmt.Println(req)

	_, err := http.Get(req)
	if err != nil {
		return err
	}

	return nil
}
