package core

import (
	"context"
	"fmt"

	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/token"
)

// LoginInput ...
type LoginInput struct {
	Context  context.Context
	Email    string
	Password string
}

// LoginOutput ...
type LoginOutput struct {
	AccessToken string
}

func (c *core) Login(input *LoginInput) (*LoginOutput, error) {
	fmt.Println("Login")

	user, err := c.Database.GetUserByEmail(&database.GetUserByEmailInput{
		Context: input.Context,
		Email:   input.Email,
	})
	if err != nil {
		return nil, err
	}

	if !c.Crypto.IsCorrectPassword(input.Password, user.PasswordHash) {
		return nil, ErrWrongPassword
	}

	session, err := c.Database.CreateSession(&database.CreateSessionInput{
		Context: input.Context,
		UserID:  user.ID,
	})
	if err != nil {
		return nil, err
	}

	tokenString, err := c.Token.NewAccessToken(&token.Content{
		ID: session.ID,
	})
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		AccessToken: tokenString,
	}, nil
}
