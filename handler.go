package main

import (
	// "encoding/json"
	// "log"

	// "strings"

	// "github.com/TechProber/pro-bot/method"
	// "github.com/TechProber/pro-bot/model"
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// func handler(res http.ResponseWriter, req *http.Request) {

// body := &model.WebhookReqBody{} // decode the JSON response body
// if err := json.NewDecoder(req.Body).Decode(body); err != nil {
// 	log.Println("could not decode request body", err)
// 	return
// }

// if !strings.Contains(strings.ToLower(body.Message.Text), "morning") {
// 	return
// }

// if err := method.Hello(body.Message.Chat.ID, body.Message.Text); err != nil {
// 	log.Println("error in sending reply:", err)
// 	return
// }

// log.Println("reply sent")
// }
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	params := request.QueryStringParameters
	fmt.Println(request.HTTPMethod, request.Path)

	name := "World"
	if params["name"] != "" {
		name = params["name"]
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hi " + name,
	}, nil

}

func main() {
	// http.ListenAndServe(":3000", http.HandlerFunc(handler))
	lambda.Start(handler)
}
