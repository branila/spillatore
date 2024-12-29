package config

import (
	"encoding/json"
	"log"
	"os"
)

var Token string
var Webhook string

type Config struct {
	// The Telegram bot token
	Token string `json:"token"`

	// The URL where Telegram will send updates
	Webhook string `json:"webhook"`
}

func Setup() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Error opening config file: %s", err)
	}

	var config Config

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding config file: %s", err)
	}

	Token = config.Token
	Webhook = config.Webhook
}
