//go:generate mockgen -destination ../mock/database.go -package mock github.com/davidchristie/identity/database Database

package database

import (
	"context"
)

// CreateAccessTokenInput ...
type CreateAccessTokenInput struct{}

// CreateUserInput ...
type CreateUserInput struct {
	Context      context.Context
	Email        string
	PasswordHash []byte
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
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id string) (*User, error)
	UpdateUser(*UpdateUserInput) error
}
