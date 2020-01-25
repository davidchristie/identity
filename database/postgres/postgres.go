package postgres

import (
	"database/sql"
	"errors"
	"log"

	"github.com/davidchristie/identity/database"
	_ "github.com/lib/pq" // postgres driver
)

const connStr = "postgres://identity:identity@postgres:5432/identity?sslmode=disable"

// ErrNotImplemented the method has not been implemented yet.
var ErrNotImplemented = errors.New("not implemented")

type postgresDatabase struct {
	DB *sql.DB
}

// New creates a new Postgres database instance.
func New() database.Database {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	return &postgresDatabase{
		DB: db,
	}
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

func (p *postgresDatabase) GetUserByEmail(input *database.GetUserByEmailInput) (*database.User, error) {
	return nil, ErrNotImplemented
}

func (p *postgresDatabase) GetUserByID(id string) (*database.User, error) {
	return nil, ErrNotImplemented
}

func (p *postgresDatabase) UpdateUser(input *database.UpdateUserInput) error {
	return ErrNotImplemented
}
