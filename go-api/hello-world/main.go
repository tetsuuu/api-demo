package main

import (
	"encoding/json"
	"log"
	"strings"
	"net/url"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

type SlackMessage struct {
	Text string `json:"text"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var keyword string
	var ignore string
	for _, value := range strings.Split(request.Body, "&") {
		param := strings.Split(value, "=")
		if param[0] == "trigger_word" {
			keyword, _ = url.QueryUnescape(param[1])
		}
		if param[0] == "user_name" {
			ignore, _ = url.QueryUnescape(param[1])
		}
	}

	if ignore == "slackbot" {
		return events.APIGatewayProxyResponse {}, nil
	}

	var text string
	if keyword == "damn it" {
		text = "What's happen bro!?"
	} else if keyword == "weather" {
		text = "The mood is always sunny!"
	} else if keyword == "がんばる" || keyword == "頑張る" {
		text = "おう！気張りやにーちゃん！"
	} else if keyword == "つかれた" || keyword == "疲れた" {
		text = "たまには休んでええんやで？"
	} else if keyword == "おはよう" || keyword == "おはよー" {
		text = "やかましいわ！もう少し寝させんかい！"
	}

	j, err := json.Marshal(SlackMessage{Text: text})
	if err != nil {
		log.Print(err)
		return events.APIGatewayProxyResponse{Body: "Error"}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(j),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
