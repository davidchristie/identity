package entity

import (
	"github.com/google/uuid"
)

// Session ...
type Session struct {
	CreatedAt string
	DeletedAt *string
	ID        uuid.UUID
	UserID    uuid.UUID
}
