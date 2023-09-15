package domain

import (
	"encoding/json"

	"log"
)

type Lexical struct {
	Root interface{} `json:"root"`
}

func CreateLexical(text string) Lexical {
	l := Lexical{}

	err := json.Unmarshal([]byte(text), &l.Root)
	if err != nil {
		log.Fatalf("could not unmarshal: %v", err)
	}

	return l
}
