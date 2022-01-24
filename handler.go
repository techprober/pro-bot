package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/TechProber/pro-bot/method"
	"github.com/TechProber/pro-bot/model"
)

// This handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &model.WebhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println("could not decode request body", err)
		return
	}

	if !strings.Contains(strings.ToLower(body.Message.Text), "hi") {
		return
	}

	if err := method.SayHi(body.Message.Chat.ID); err != nil {
		log.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	log.Println("reply sent")
}

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}
