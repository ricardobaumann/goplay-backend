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

	println(cardMsg)

	carouselMsg := `
	{
		"data": {
			"google": {
			  "expectUserResponse": false
			}
		  },
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
				"items": [
				  {
					"description": "Option One Description",
					"image": {
					  "url": "https://www.welt.de/img/kmpkt/mobile172212734/4592507717-ci102l-w120/woman-working-with-a-baby.jpg",
					  "accessibilityText": "acessibility text for option one"
					},
					"optionInfo": {
					  "key": "itemOne",
					  "synonyms": [
						"thing one",
						"object one"
					  ]
					},
					"title": "Option One Title"
				  },
				  {
					"description": "Option Two Description",
					"image": {
					  "url": "https://www.welt.de/img/wirtschaft/mobile164536178/7232506147-ci102l-w120/Iona-Bresser-Mutter-von-6-Kinder-8.jpg",
					  "accessibilityText": "acessibility text for option two"
					},
					"optionInfo": {
					  "key": "itemTwo",
					  "synonyms": [
						"thing two",
						"object two"
					  ]
					},
					"title": "Option Two Title"
				  }
				],
				"platform": "google",
				"type": "carousel_card"
			  }

		]
	}
		`
	//respBytes, _ := json.Marshal(&botResponse)

	return events.APIGatewayProxyResponse{Body: carouselMsg, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
