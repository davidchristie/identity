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
func New(conf config.Database) database.Database {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", conf.Username(), conf.Password(), conf.Host(), conf.Name())
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	err = runMigrations(db, conf)
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

func runMigrations(db *sql.DB, conf config.Database) error {
	err := db.Ping()
	if err != nil {
		fmt.Println("Database migrations failed: ", err)
		fmt.Println("Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
		return runMigrations(db, conf)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+conf.Migrations(), "postgres", driver)
	if err != nil {
		return err
	}
	m.Up()
	return nil
}
