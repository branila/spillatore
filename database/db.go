package database

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/branila/spillatore/types"
)

const dbName = "spillatore.json"

var database = types.Database{}

func syncDatabase() {
	jsonData, err := json.MarshalIndent(database, "", "  ")
	if err != nil {
		log.Printf("Error marshalling database: %s", err)
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory: %s", err)
		return
	}

	dbPath := filepath.Join(wd, dbName)

	err = os.WriteFile(dbPath, jsonData, fs.FileMode(0644))
	if err != nil {
		log.Printf("Error writing database: %s", err)
		return
	}
}
