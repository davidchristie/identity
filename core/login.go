package core

import (
	"context"
	"fmt"

	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/jwt"
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

	token, err := c.Database.CreateAccessToken(&database.CreateAccessTokenInput{
		Context: input.Context,
		UserID:  user.ID,
	})
	if err != nil {
		return nil, err
	}

	tokenString, err := c.JWT.SignedString(&jwt.SignedStringInput{
		ID: token.ID,
	})
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		AccessToken: tokenString,
	}, nil
}
