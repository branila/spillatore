package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/branila/spillatore/types"
)

func parseRequest(r *http.Request) (*types.Update, error) {
	var update types.Update

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Printf("Error parsing request body: %s", err.Error())
		return nil, err
	}

	return &update, nil
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	update, err := parseRequest(r)
	if err != nil {
		log.Printf("Error parsing request: %s", err.Error())
		return
	}

	message := fmt.Sprintf(`Ma guardatelo, il coglione ha detto "%s"`, update.Message.Text)

	responseBody, err := sendText(update.Message.Chat.Id, message)
	if err != nil {
		log.Printf("Got error %s from telegram, response body is %s", err.Error(), responseBody)
		return
	}

	log.Printf("Message sent to chat %d", update.Message.Chat.Id)
}

func sendText(chatId int, text string) (string, error) {
	log.Printf("Sending %s to chat_id: %d", text, chatId)

	api := "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"

	fmt.Println(api)

	response, err := http.PostForm(
		api,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		},
	)

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}

	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}

	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}

func main() {
	http.HandleFunc("/telegram", handleUpdate)

	log.Printf("Server listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
