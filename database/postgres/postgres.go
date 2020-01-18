package postgres

import "github.com/davidchristie/identity/database"

import "errors"

const connStr = "postgres://accounts:acc0unts_secret123@accounts-database:5432/accounts?sslmode=disable"

var errNotImplemented = errors.New("Not implemeneted")

type postgresDatabase struct{}

// New ...
func New() database.Database {
	return &postgresDatabase{}
}

func (p *postgresDatabase) CreateAccessToken(*database.CreateAccessTokenInput) (*database.AccessToken, error) {
	return nil, errNotImplemented
}

func (p *postgresDatabase) CreateUser(*database.CreateUserInput) (*database.User, error) {
	return nil, errNotImplemented
}

func (p *postgresDatabase) DeleteAccessToken(id string) error {
	return errNotImplemented
}

func (p *postgresDatabase) DeleteUser(id string) error {
	return errNotImplemented
}

func (p *postgresDatabase) GetAccessTokenByID(id string) (*database.AccessToken, error) {
	return nil, errNotImplemented
}

func (p *postgresDatabase) GetUserByEmail(email string) (*database.User, error) {
	return nil, errNotImplemented
}

func (p *postgresDatabase) GetUserByID(id string) (*database.User, error) {
	return nil, errNotImplemented
}

func (p *postgresDatabase) UpdateUser(input *database.UpdateUserInput) error {
	return errNotImplemented
}
