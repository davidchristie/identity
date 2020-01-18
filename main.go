package main

import (
	"fmt"
	"log"

	"github.com/davidchristie/identity/core"
	"github.com/davidchristie/identity/crypto"
	"github.com/davidchristie/identity/crypto/bcrypt"
	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/database/postgres"
	"github.com/davidchristie/identity/server"
	"github.com/davidchristie/identity/server/http"
)

type options struct {
	Crypto   crypto.Crypto
	Database database.Database
	Server   server.Server
}

func main() {
	err := startService(&options{
		Crypto:   bcrypt.New(),
		Database: postgres.New(),
		Server:   http.New(http.Options{}),
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
	})
	return options.Server.Serve(core)
}
