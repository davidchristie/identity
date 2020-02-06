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
	"github.com/davidchristie/identity/jwt"
	jwtGo "github.com/davidchristie/identity/jwt/jwt_go"
	"github.com/davidchristie/identity/server"
	"github.com/davidchristie/identity/server/http"
)

type options struct {
	Crypto   crypto.Crypto
	Database database.Database
	JWT      jwt.JWT
	Server   server.Server
}

func main() {
	config := env.New()
	err := startService(&options{
		Crypto:   bcrypt.New(),
		Database: postgres.New(config.Database()),
		JWT:      jwtGo.New(config.Token()),
		Server:   http.New(config.Server()),
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
		JWT:      options.JWT,
	})
	return options.Server.Serve(core)
}
