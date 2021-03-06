package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/davidchristie/identity/database"
)

const minPasswordLength = 8

// SignupInput ...
type SignupInput struct {
	Context  context.Context
	Email    string
	Password string
}

// SignupOutput ...
type SignupOutput struct{}

func (c *core) Signup(input *SignupInput) (*SignupOutput, error) {
	err := validateSignupInput(input)
	if err != nil {
		return nil, err
	}
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
		switch err {
		case database.ErrDuplicateUserEmail:
			return nil, ErrEmailAlreadyInUse

		default:
			fmt.Println("[Signup] Unknown error: ", err)
			return nil, errors.New("unknown error")
		}
	}
	fmt.Println("User created")
	return &SignupOutput{}, nil
}

func validateSignupInput(input *SignupInput) error {
	if len(input.Password) < minPasswordLength {
		return ErrShortPassword
	}
	return nil
}
