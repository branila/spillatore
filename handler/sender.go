package handler

import (
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/branila/spillatore/config"
)

func sendText(chatId int, text string) (string, error) {
	api := "https://api.telegram.org/bot" + config.Token + "/sendMessage"

	response, err := http.PostForm(
		api,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		},
	)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	bodyString := string(bodyBytes)

	return bodyString, nil
}
