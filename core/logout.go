package core

// LogoutInput ...
type LogoutInput struct {
	AccessToken string
}

// LogoutOutput ...
type LogoutOutput struct{}

func (c *core) Logout(input *LogoutInput) (*LogoutOutput, error) {
	token, err := c.JWT.Parse(input.AccessToken)
	if err != nil {
		return nil, err
	}
	err = c.Database.DeleteAccessToken(token.ID)
	if err != nil {
		return nil, err
	}
	return &LogoutOutput{}, nil
}
