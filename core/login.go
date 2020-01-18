package core

import (
	"fmt"
)

// LoginInput ...
type LoginInput struct {
	Email    string
	Password string
}

// LoginOutput ...
type LoginOutput struct {
	AccessToken string
}

func (c *core) Login(input *LoginInput) (*LoginOutput, error) {
	fmt.Println("Login")
	return &LoginOutput{}, nil
}
