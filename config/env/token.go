package env

import (
	"log"
	"os"
)

type token struct{}

func (t *token) Secret() string {
	secret := os.Getenv("TOKEN_SECRET")
	if secret == "" {
		log.Fatal("$TOKEN_SECRET must be set")
	}
	return secret
}
