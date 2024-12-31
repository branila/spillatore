package config

import (
	"encoding/json"
	"fmt"
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
	const defaultConfigPath = "./config.json"

	// Check if the config file exists
	if _, err := os.Stat(defaultConfigPath); os.IsNotExist(err) {
		fmt.Printf("Config file not found. Creating default config file at %s\n", defaultConfigPath)

		defaultConfig := Config{
			Token:   "your-telegram-bot-token",
			Webhook: "your-webhook-url",
		}

		file, err := os.Create(defaultConfigPath)
		if err != nil {
			log.Fatalf("Error creating default config file: %s\n", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")

		if err := encoder.Encode(defaultConfig); err != nil {
			log.Fatalf("Error writing default config file: %s\n", err)
		}

		log.Println("Default config file created. Please update it with your settings and restart the application.")

		os.Exit(1)
	}

	// Open and parse the config file
	file, err := os.Open(defaultConfigPath)
	if err != nil {
		log.Fatalf("Error opening config file: %s", err)
	}
	defer file.Close()

	var config Config

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding config file: %s\n", err)
	}

	Token = config.Token
	Webhook = config.Webhook
}
