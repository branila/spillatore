package handler

import (
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

	if !strings.Contains(receivedMsg, "@SpillatoreBot") {
		return
	}

	reply := processMessage(receivedMsg)

	sendReply(update.Message.Chat.Id, reply)
}

func getMessageText(update types.Update) string {
	if update.Message.Text != "" {
		return update.Message.Text // The message is a text message
	}

	return update.Message.Caption // The message is an image with a caption
}

func processAddition(msg string) string {
	num := extractNumber(msg)

	database.IncrementCounter(num)

	return "Carusi siamo a " + strconv.Itoa(database.GetCounter()) + " bire 🍻"
}

func processSubtraction(msg string) string {
	num := extractNumber(msg)

	database.DecrementCounter(num)

	return "Carusi siamo a " + strconv.Itoa(database.GetCounter()) + " bire 🍻"
}

func processSetting(msg string) string {
	num := extractNumber(msg)

	database.SetCounter(num)

	return "Picciotti ne abbiamo " + strconv.Itoa(database.GetCounter())
}

func processMessage(msg string) string {
	msg = strings.ToLower(msg)

	additionRegexp := regexp.MustCompile(`\+(\d+)`)
	subtractionRegexp := regexp.MustCompile(`-(\d+)`)
	setRegexp := regexp.MustCompile(`siamo a (\d+)`)
	getCounterRegexp1 := regexp.MustCompile(`a quante siamo`)
	getCounterRegexp2 := regexp.MustCompile(`quante ne abbiamo`)
	sanFaiRegexp := regexp.MustCompile(`san fai`)
	regoleRegexp := regexp.MustCompile(`regole`)

	switch {
	case additionRegexp.MatchString(msg):
		return processAddition(msg)

	case subtractionRegexp.MatchString(msg):
		return processSubtraction(msg)

	case setRegexp.MatchString(msg):
		return processSetting(msg)

	case getCounterRegexp1.MatchString(msg) || getCounterRegexp2.MatchString(msg):
		return "Siamo a " + strconv.Itoa(database.GetCounter()) + " bire 🍻"

	case sanFaiRegexp.MatchString(msg):
		return "Ma stai zitto"

	case regoleRegexp.MatchString(msg):
		return "Non rompere i coglioni"

	default:
		return "Lasciami in pace"
	}
}

func extractNumber(msg string) int {
	numRegexp := regexp.MustCompile(`\d+`)
	match := numRegexp.FindString(msg)
	num, _ := strconv.Atoi(match) // Non serve gestire l'errore, il numero è garantito
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
