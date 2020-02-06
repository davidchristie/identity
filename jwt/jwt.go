//go:generate mockgen -destination ../mock/jwt.go -package mock github.com/davidchristie/identity/jwt JWT

package jwt

import (
	"github.com/davidchristie/identity/database"
	"github.com/google/uuid"
)

// SignedStringInput ...
type SignedStringInput struct {
	ID uuid.UUID
}

// JWT ...
type JWT interface {
	Parse(string) (*database.AccessToken, error)
	SignedString(*SignedStringInput) (string, error)
}
