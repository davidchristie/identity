package core

import "github.com/davidchristie/identity/database"

// SignupInput ...
type SignupInput struct {
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
	_, err = c.Database.CreateUser(&database.CreateUserInput{
		Email:        input.Email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}
	return &SignupOutput{}, nil
}
