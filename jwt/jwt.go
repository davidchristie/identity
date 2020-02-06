//go:generate mockgen -destination ../mock/jwt.go -package mock github.com/davidchristie/identity/jwt JWT

package jwt

import (
	"github.com/davidchristie/identity/entity"
	"github.com/google/uuid"
)

// SignedStringInput ...
type SignedStringInput struct {
	ID uuid.UUID
}

// JWT ...
type JWT interface {
	Parse(string) (*entity.Session, error)
	SignedString(*SignedStringInput) (string, error)
}
