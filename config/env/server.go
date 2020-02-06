package env

import (
	"log"
	"os"
)

type server struct{}

func (d *server) Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	return port
}
