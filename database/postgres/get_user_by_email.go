package postgres

import (
	"database/sql"
	"errors"

	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/entity"
	"github.com/google/uuid"
)

// ErrNotFound is returned when a queried entity does not exist.
var ErrNotFound = errors.New("not found")

// GetUserByEmail attempts to find a user with the specified email.
// If the account does not exist a ErrNotFound error is returned.
func (p *postgresDatabase) GetUserByEmail(input *database.GetUserByEmailInput) (*entity.User, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	const query = `
		SELECT id, email, password_hash FROM users
		WHERE email = $1
	`

	row := db.QueryRow(query, input.Email)
	var rowID uuid.UUID
	var rowEmail string
	var rowPasswordHash []byte
	if err := row.Scan(&rowID, &rowEmail, &rowPasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &entity.User{
		Email:        rowEmail,
		ID:           rowID,
		PasswordHash: rowPasswordHash,
	}, nil
}
