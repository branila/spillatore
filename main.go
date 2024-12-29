package main

import (
	"log"
	"net/http"

	"github.com/branila/spillatore/config"
	"github.com/branila/spillatore/database"
	"github.com/branila/spillatore/handler"
	"github.com/branila/spillatore/webhook"
)

func main() {
	database.Init()
	config.Setup()
	webhook.Set()

	http.HandleFunc("/telegram", handler.Master)

	log.Printf("Server listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
