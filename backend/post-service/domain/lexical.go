package domain

import (
	"encoding/json"

	"log"
)


func CreateLexical(text string) map[string]interface{} {
	var l map[string]interface{}
	err := json.Unmarshal([]byte(text), &l)
	if err != nil {
		log.Fatalf("could not unmarshal: %v", err)
	}

	return l
}
