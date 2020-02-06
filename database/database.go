//go:generate mockgen -destination ../mock/database.go -package mock github.com/davidchristie/identity/database Database

package database

import (
	"context"

	"github.com/davidchristie/identity/entity"
	"github.com/google/uuid"
)

// CreateSessionInput ...
type CreateSessionInput struct {
	Context context.Context
	UserID  uuid.UUID
}

// CreateUserInput ...
type CreateUserInput struct {
	Context      context.Context
	Email        string
	PasswordHash []byte
}

// GetUserByEmailInput ...
type GetUserByEmailInput struct {
	Context context.Context
	Email   string
}

// UpdateUserInput ...
type UpdateUserInput struct{}

// Database ...
type Database interface {
	CreateSession(*CreateSessionInput) (*entity.Session, error)
	CreateUser(*CreateUserInput) (*entity.User, error)
	DeleteSession(id uuid.UUID) error
	DeleteUser(id uuid.UUID) error
	GetSessionByID(id uuid.UUID) (*entity.Session, error)
	GetUserByEmail(*GetUserByEmailInput) (*entity.User, error)
	GetUserByID(id uuid.UUID) (*entity.User, error)
	UpdateUser(*UpdateUserInput) error
}
