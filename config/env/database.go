package env

import (
	"log"
	"os"
)

type database struct{}

func (d *database) Host() string {
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		log.Fatal("$DATABASE_HOST must be set")
	}
	return host
}

func (d *database) Password() string {
	password := os.Getenv("DATABASE_PASSWORD")
	if password == "" {
		log.Fatal("$DATABASE_PASSWORD must be set")
	}
	return password
}

func (d *database) Username() string {
	username := os.Getenv("DATABASE_USERNAME")
	if username == "" {
		log.Fatal("$DATABASE_USERNAME must be set")
	}
	return username
}
