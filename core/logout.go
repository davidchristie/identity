package core

import (
	"fmt"
)

// LogoutInput ...
type LogoutInput struct {
	AccessToken string
}

// LogoutOutput ...
type LogoutOutput struct{}

func (c *core) Logout(input *LogoutInput) (*LogoutOutput, error) {
	fmt.Println("Logout")
	return &LogoutOutput{}, nil
}
