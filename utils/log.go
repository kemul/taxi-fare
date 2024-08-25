package utils

import (
	"encoding/json"
	"log"
	"time"
)

func LogError(err error) error {
	logData := map[string]string{
		"error": err.Error(),
		"time":  time.Now().Format(time.RFC3339),
	}
	logJSON, _ := json.Marshal(logData)
	log.Println(string(logJSON))
	return err
}
