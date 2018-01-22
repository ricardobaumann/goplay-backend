package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	cardMsg := fmt.Sprintf(`
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
				"title": "%s
				"subtitle": "%s",
				"formattedText": "%s",
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
			}
		]
	}
		`, "title text", "subtitle text", "text with newlines and such")

	//respBytes, _ := json.Marshal(&botResponse)

	return events.APIGatewayProxyResponse{Body: cardMsg, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
