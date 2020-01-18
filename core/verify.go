package core

import (
	"fmt"
)

// VerifyInput ...
type VerifyInput struct {
	AccessToken string
}

// VerifyOutput ...
type VerifyOutput struct {
	Email string
	ID    string
}

func (c *core) Verify(input *VerifyInput) (*VerifyOutput, error) {
	fmt.Println("Verify")
	return &VerifyOutput{}, nil
}
