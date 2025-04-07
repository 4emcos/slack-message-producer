package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type SlackMessage struct {
	Text      string `json:"text"`
	ChannelId string `json:"channel"`
}

func main() {
	slackAuthorizerKey := os.Getenv("SLACK_AUTHORIZER_KEY")

	slackMessage := SlackMessage{
		Text:      "Hello from Bot! in " + time.Now().Format("2006-01-02 15:04:05"),
		ChannelId: "C08M5M73Y2G",
	}

	slackMessage2, _ := json.Marshal(slackMessage)

	payload := bytes.NewBuffer(slackMessage2)

	req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/chat.postMessage", payload)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+slackAuthorizerKey)

	resp, err := http.DefaultClient.Do(req)

	log.Println(resp.StatusCode)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(string(responseBody))
}
