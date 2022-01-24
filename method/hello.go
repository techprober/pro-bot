package method

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/TechProber/pro-bot/model"
)

func Hello(chatID int64, message string) error {
	log.Println(message)
	reqBody := &model.MessageReqBody{
		ChatID: chatID,
		Text:   "Nihao!!",
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	res, err := http.Post(
		model.TelegramApi+model.TelegramApiSendMessagePath,
		"application/json",
		bytes.NewBuffer(reqBytes),
	)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
