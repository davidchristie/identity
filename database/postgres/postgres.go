package postgres

import (
	"errors"

	"github.com/davidchristie/identity/database"
	_ "github.com/lib/pq" // postgres driver
)

const connStr = "postgres://identity:identity@postgres:5432/identity?sslmode=disable"

// ErrNotImplemented the method has not been implemented yet.
var ErrNotImplemented = errors.New("not implemented")

type postgresDatabase struct{}

// New creates a new Postgres database instance.
func New() database.Database {
	return &postgresDatabase{}
}

func (p *postgresDatabase) CreateAccessToken(*database.CreateAccessTokenInput) (*database.AccessToken, error) {
	return nil, ErrNotImplemented
}

func (p *postgresDatabase) DeleteAccessToken(id string) error {
	return ErrNotImplemented
}

func (p *postgresDatabase) DeleteUser(id string) error {
	return ErrNotImplemented
}

func (p *postgresDatabase) GetAccessTokenByID(id string) (*database.AccessToken, error) {
	return nil, ErrNotImplemented
}

func (p *postgresDatabase) GetUserByEmail(email string) (*database.User, error) {
	return nil, ErrNotImplemented
}

func (p *postgresDatabase) GetUserByID(id string) (*database.User, error) {
	return nil, ErrNotImplemented
}

func (p *postgresDatabase) UpdateUser(input *database.UpdateUserInput) error {
	return ErrNotImplemented
}
