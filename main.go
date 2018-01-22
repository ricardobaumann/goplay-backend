package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//APIAIRequest : Incoming request format from APIAI
type APIAIRequest struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Result    struct {
		Parameters map[string]string `json:"parameters"`
		Contexts   []interface{}     `json:"contexts"`
		Metadata   struct {
			IntentID                  string `json:"intentId"`
			WebhookUsed               string `json:"webhookUsed"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
			IntentName                string `json:"intentName"`
		} `json:"metadata"`
		Score float32 `json:"score"`
	} `json:"result"`
	Status struct {
		Code      int    `json:"code"`
		ErrorType string `json:"errorType"`
	} `json:"status"`
	SessionID       string      `json:"sessionId"`
	OriginalRequest interface{} `json:"originalRequest"`
}

//APIAIMessage : Response Message Structure
type APIAIMessage struct {
	Speech      string `json:"speech"`
	DisplayText string `json:"displayText"`
	Source      string `json:"source"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	var t APIAIRequest
	json.Unmarshal([]byte(request.Body), &t)

	//msg := "{\"speech\": \"speech\", \"displayText\" : \"display\"}" //json.Marshal(APIAIMessage{Source: "Hotel Feedback System", Speech: "Thank you for the feedback", DisplayText: "Thank you for the feedback"})

	msg := `
	{
		
		"messages": [
			{
				"speech": "content to be read aloud",
				"type": 0
			},
	
	
			{
				"platform": "google",
				"type": "simple_response",
				"displayText": "top level text", 
				"textToSpeech": "voice speech to be read out loud"  
			},
			{
				"platform": "google",
				"type": "basic_card",
				"title": "title text",
				"subtitle": "subtitle text",
				"formattedText": "text with newlines and such",
				"image": {
					"url": "http://example.com/image.png",
					"accessibilityText": "image descrition for screen readers"  
				},
				"buttons": [
					{
						"title": "Link title",
						"openUrlAction": {
							"url": "https://example.com/linkout.html"
						}
					}
				]
			},
			{
				"platform": "google",
				"type": "suggestion_chips",
				"suggestions": [
					{
						"title": "Next"
					},
					{
						"title": "Previous"
					},
					{
						"title": "Return to Results"
					}
				]
			}
		]
	}
		`

	//respBytes, _ := json.Marshal(&botResponse)

	return events.APIGatewayProxyResponse{Body: msg, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
