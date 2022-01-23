package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// This handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &WebhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	// Check if the message contains the word "marco"
	// if not, return without doing aything
	if !strings.Contains(strings.ToLower(body.Message.Text), "hi") {
		return
	}

	// If the text contains marco, call the `sayPolo` function, which
	// is defined below
	if err := SayHi(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

//The below code deals with the process of sending a response message to the user

// sayHi takes a chatID and sends "nihao" to them

// FInally, the main funtion starts our server on port 3000
func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}
