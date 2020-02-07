package core

import (
	"fmt"
)

// GetUserInput ...
type GetUserInput struct {
	AccessToken string
}

// GetUserOutput ...
type GetUserOutput struct {
	Email string
}

func (c *core) GetUser(input *GetUserInput) (*GetUserOutput, error) {
	fmt.Println("GetUser")

	tokenContent, err := c.Token.ParseAccessToken(input.AccessToken)
	if err != nil {
		return nil, err
	}

	session, err := c.Database.GetSessionByID(tokenContent.ID)
	if err != nil {
		return nil, err
	}

	user, err := c.Database.GetUserByID(session.UserID)
	if err != nil {
		return nil, err
	}

	return &GetUserOutput{
		Email: user.Email,
	}, nil
}
