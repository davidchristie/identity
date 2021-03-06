package main

import (
	"fmt"
	"log"

	"github.com/davidchristie/identity/config/env"
	"github.com/davidchristie/identity/core"
	"github.com/davidchristie/identity/crypto"
	"github.com/davidchristie/identity/crypto/bcrypt"
	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/database/postgres"
	"github.com/davidchristie/identity/server"
	"github.com/davidchristie/identity/server/http"
	"github.com/davidchristie/identity/token"
	"github.com/davidchristie/identity/token/jwt"
)

type options struct {
	Crypto   crypto.Crypto
	Database database.Database
	Server   server.Server
	Token    token.Token
}

func main() {
	config := env.New()
	err := startService(&options{
		Crypto:   bcrypt.New(),
		Database: postgres.New(config.Database()),
		Server:   http.New(config.Server()),
		Token:    jwt.New(config.Token()),
	})
	if err != nil {
		log.Fatal("Fatal error in service: ", err)
	}
}

func startService(options *options) error {
	fmt.Println("Starting service...")
	core := core.New(core.Options{
		Crypto:   options.Crypto,
		Database: options.Database,
		Token:      options.Token,
	})
	return options.Server.Serve(core)
}
