package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/branila/spillatore/config"
)

func Set() {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook", config.Token)

	data, err := json.Marshal(map[string]string{
		"url": config.Webhook,
	})
	if err != nil {
		log.Fatalf("Failed to marshal payload: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	var result struct {
		Ok          bool   `json:"ok"`
		Result      bool   `json:"result"`
		Description string `json:"description"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	if !result.Ok {
		log.Fatalf("Failed to set webhook: %s", result.Description)
	}
}
