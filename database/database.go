//go:generate mockgen -destination ../mock/database.go -package mock github.com/davidchristie/identity/database Database

package database

import (
	"context"

	"github.com/davidchristie/identity/entity"
	"github.com/google/uuid"
)

// CreateAccessTokenInput ...
type CreateAccessTokenInput struct {
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
	CreateAccessToken(*CreateAccessTokenInput) (*entity.AccessToken, error)
	CreateUser(*CreateUserInput) (*entity.User, error)
	DeleteAccessToken(id uuid.UUID) error
	DeleteUser(id uuid.UUID) error
	GetAccessTokenByID(id uuid.UUID) (*entity.AccessToken, error)
	GetUserByEmail(*GetUserByEmailInput) (*entity.User, error)
	GetUserByID(id uuid.UUID) (*entity.User, error)
	UpdateUser(*UpdateUserInput) error
}
