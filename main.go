package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type BotRequest struct {
	QueryResult QueryResult `json:"queryResult"`
}

type QueryResult struct {
	Parameters map[string]string `json:"parameters"`
}

type Message struct {
	Speech string `json:"speech"`
	Type   int    `json:"type"`
}

type FulfillmentMessage struct {
	//Card           Card             `json:"card"`
	Messages []Message `json:"messages"`
	//Text []string `json:"text"`
}

type BotResponse struct {
	//FulfillmentText     string               `json:"fulfillmentText"`
	//FulfillmentMessages []FulfillmentMessage `json:"fulfillmentMessages"`
	Messages []Message `json:"messages"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	var botRequest BotRequest
	err := json.Unmarshal([]byte(request.Body), &botRequest)
	if err != nil {
		panic(err)
	}
	botResponse := &BotResponse{
		Messages: []Message{
			Message{
				Speech: "speech",
				Type:   0,
			},
		},
	}

	respBytes, _ := json.Marshal(&botResponse)

	return events.APIGatewayProxyResponse{Body: string(respBytes), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
