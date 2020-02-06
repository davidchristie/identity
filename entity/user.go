package entity

import (
	"github.com/google/uuid"
)

// User ...
type User struct {
	CreatedAt    string
	DeletedAt    *string
	Email        string
	ID           uuid.UUID
	PasswordHash []byte
	UpdatedAt    string
}
