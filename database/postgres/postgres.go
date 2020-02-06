package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/davidchristie/identity/config"
	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/entity"
	"github.com/google/uuid"
	_ "github.com/lib/pq" // postgres driver
)

const connStr = "postgres://identity:identity@postgres:5432/identity?sslmode=disable"

// ErrNotImplemented the method has not been implemented yet.
var ErrNotImplemented = errors.New("not implemented")

type postgresDatabase struct {
	DB *sql.DB
}

// New creates a new Postgres database instance.
func New(c config.Database) database.Database {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/identity?sslmode=disable", c.Username(), c.Password(), c.Host())
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	return &postgresDatabase{
		DB: db,
	}
}

func (p *postgresDatabase) DeleteSession(id uuid.UUID) error {
	return ErrNotImplemented
}

func (p *postgresDatabase) DeleteUser(id uuid.UUID) error {
	return ErrNotImplemented
}

func (p *postgresDatabase) GetSessionByID(id uuid.UUID) (*entity.Session, error) {
	return nil, ErrNotImplemented
}

func (p *postgresDatabase) GetUserByID(id uuid.UUID) (*entity.User, error) {
	return nil, ErrNotImplemented
}

func (p *postgresDatabase) UpdateUser(input *database.UpdateUserInput) error {
	return ErrNotImplemented
}
