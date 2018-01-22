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

type Button struct {
	Text     string `json:"text"`
	Postback string `json:"postback"`
}

type Card struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	//ImageURI string   `json:"imageUri"`
	Buttons []Button `json:"buttons"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
	DisplayText  string `json:"displayText"`
}

type FulfillmentMessage struct {
	//Card           Card             `json:"card"`
	SimpleResponse []SimpleResponse `json:"simpleResponses"`
	//Text []string `json:"text"`
}

type BotResponse struct {
	//FulfillmentText     string               `json:"fulfillmentText"`
	FulfillmentMessages []FulfillmentMessage `json:"fulfillmentMessages"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	var botRequest BotRequest
	err := json.Unmarshal([]byte(request.Body), &botRequest)
	if err != nil {
		panic(err)
	}
	botResponse := &BotResponse{

		FulfillmentMessages: []FulfillmentMessage{
			FulfillmentMessage{
				SimpleResponse: []SimpleResponse{
					SimpleResponse{
						TextToSpeech: "text to speech",
						DisplayText:  "display text",
					},
				},
			}}}

	respBytes, _ := json.Marshal(&botResponse)

	return events.APIGatewayProxyResponse{Body: string(respBytes), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
