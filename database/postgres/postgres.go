package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/davidchristie/identity/config"
	"github.com/davidchristie/identity/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // file driver
	"github.com/google/uuid"
	_ "github.com/lib/pq" // postgres driver
)

// ErrNotImplemented the method has not been implemented yet.
var ErrNotImplemented = errors.New("not implemented")

type postgresDatabase struct {
	DB *sql.DB
}

// New creates a new Postgres database instance.
func New(c config.Database) database.Database {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", c.Username(), c.Password(), c.Host(), c.Name())
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	err = runMigrations(db)
	if err != nil {
		log.Fatal("Error running database migrations: ", err)
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

func (p *postgresDatabase) UpdateUser(input *database.UpdateUserInput) error {
	return ErrNotImplemented
}

func runMigrations(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		fmt.Println("Database migrations failed: ", err)
		fmt.Println("Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
		return runMigrations(db)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
	if err != nil {
		return err
	}
	m.Up()
	return nil
}
