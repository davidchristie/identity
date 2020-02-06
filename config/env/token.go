package env

import (
	"log"
	"os"
)

type token struct{}

func (t *token) Secret() string {
	port := os.Getenv("TOKEN_SECRET")
	if port == "" {
		log.Fatal("$TOKEN_SECRET must be set")
	}
	return port
}
