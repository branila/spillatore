package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/branila/spillatore/database"
	"github.com/branila/spillatore/types"
)

// Entry point for the HTTP handler
func Master(w http.ResponseWriter, r *http.Request) {
	update, err := parseRequest(r)
	if err != nil {
		log.Printf("Error parsing request: %s", err)
		return
	}

	receivedMsg := getMessageText(*update)

	if !strings.HasPrefix(receivedMsg, "@SpillatoreBot") {
		return
	}

	reply, err := processMessage(receivedMsg)
	if err != nil {
		log.Printf("Error processing message: %s", err)
	}

	fmt.Println(reply)

	sendReply(update.Message.Chat.Id, reply)
}

func getMessageText(update types.Update) string {
	if update.Message.Text != "" {
		return update.Message.Text // The message is a text message
	}

	return update.Message.Caption // The message is an image with a caption
}

func processMessage(msg string) (string, error) {
	rexp := regexp.MustCompile(`^@SpillatoreBot \+(\d+)$`)
	if !rexp.MatchString(msg) {
		return "", nil
	}

	num := extractNumber(msg)

	database.IncrementCounter(num)

	return "Carusi siamo a " + strconv.Itoa(database.GetCounter()) + " bire. 🍻", nil
}

func extractNumber(msg string) int {
	parts := strings.Split(msg, "+")

	num, _ := strconv.Atoi(parts[1])

	return num
}

func sendReply(chatID int, reply string) {
	responseBody, err := sendText(chatID, reply)
	if err != nil {
		log.Printf("Got error %s, response body is '%s'", err, responseBody)
		return
	}

	log.Printf("Message sent to chat %d", chatID)
}
