package core

import (
	"context"
	"fmt"

	"github.com/davidchristie/identity/database"
)

// SignupInput ...
type SignupInput struct {
	Context  context.Context
	Email    string
	Password string
}

// SignupOutput ...
type SignupOutput struct{}

func (c *core) Signup(input *SignupInput) (*SignupOutput, error) {
	passwordHash, err := c.Crypto.GeneratePasswordHash(input.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println("Creating new user...")
	_, err = c.Database.CreateUser(&database.CreateUserInput{
		Context:      input.Context,
		Email:        input.Email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("User created")
	return &SignupOutput{}, nil
}
