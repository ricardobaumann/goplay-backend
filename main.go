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

type BotResponse struct {
	FulfillmentText string `json:"fulfillmentText"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	var botRequest BotRequest
	err := json.Unmarshal([]byte(request.Body), &botRequest)
	if err != nil {
		panic(err)
	}
	botResponse := BotResponse{FulfillmentText: botRequest.QueryResult.Parameters["color"]}
	respBytes, _ := json.Marshal(&botResponse)

	return events.APIGatewayProxyResponse{Body: string(respBytes), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
