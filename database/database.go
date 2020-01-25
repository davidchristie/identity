//go:generate mockgen -destination ../mock/database.go -package mock github.com/davidchristie/identity/database Database

package database

import (
	"context"

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
	CreateAccessToken(*CreateAccessTokenInput) (*AccessToken, error)
	CreateUser(*CreateUserInput) (*User, error)
	DeleteAccessToken(id string) error
	DeleteUser(id string) error
	GetAccessTokenByID(id string) (*AccessToken, error)
	GetUserByEmail(*GetUserByEmailInput) (*User, error)
	GetUserByID(id string) (*User, error)
	UpdateUser(*UpdateUserInput) error
}
