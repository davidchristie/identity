package entity

import (
	"github.com/google/uuid"
)

// AccessToken ...
type AccessToken struct {
	CreatedAt string
	DeletedAt *string
	ID        uuid.UUID
	UserID    uuid.UUID
}
