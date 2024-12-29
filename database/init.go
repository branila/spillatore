package database

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func Init() {
	raw := getRawDatabase()

	err := json.Unmarshal(raw, &database)
	if err != nil {
		log.Printf("Error unmarshalling database: %s", err)
		return
	}
}

func getRawDatabase() []byte {
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory: %s", err)
		return nil
	}

	dbPath := filepath.Join(wd, dbName)

	prepareDatabase(dbPath)

	content, err := os.ReadFile(dbPath)
	if err != nil {
		log.Printf("Error reading database: %s", err)
		return nil
	}

	return content
}

func prepareDatabase(dbPath string) {
	if !dbExists(dbPath) {
		file := createDatabase(dbPath)
		defer file.Close()

		setDefault(file)
	}
}

func dbExists(dbPath string) bool {
	_, err := os.Stat(dbPath)
	return !os.IsNotExist(err)
}

func createDatabase(dbPath string) *os.File {
	file, err := os.Create(dbPath)
	if err != nil {
		log.Printf("Error creating database: %s", err)
		return nil
	}

	return file
}

func setDefault(file *os.File) {
	jsonData, err := json.MarshalIndent(database, "", "  ")
	if err != nil {
		log.Printf("Error marshalling database: %s", err)
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Printf("Error writing database: %s", err)
		return
	}
}
