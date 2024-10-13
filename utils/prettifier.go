package utils

import (
	"encoding/json"
)

func PrettifyObject(o any) string {
	data, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "Marshal Error: " + err.Error()
	}

	return string(data)
}
