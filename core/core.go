package core

import (
	"log"

	"github.com/davidchristie/identity/crypto"
	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/jwt"
)

// Core ...
type Core interface {
	Login(*LoginInput) (*LoginOutput, error)
	Logout(*LogoutInput) (*LogoutOutput, error)
	Signup(*SignupInput) (*SignupOutput, error)
	Verify(*VerifyInput) (*VerifyOutput, error)
}

type core struct {
	Crypto   crypto.Crypto
	Database database.Database
	JWT      jwt.JWT
}

// Options ...
type Options struct {
	Crypto   crypto.Crypto
	Database database.Database
	JWT      jwt.JWT
}

// New ...
func New(options Options) Core {
	if options.Crypto == nil {
		log.Fatal("Crypto implementation not specified")
	}
	if options.Database == nil {
		log.Fatal("Database implementation not specified")
	}
	if options.JWT == nil {
		log.Fatal("JWT implementation not specified")
	}
	return &core{
		Crypto:   options.Crypto,
		Database: options.Database,
		JWT:      options.JWT,
	}
}
