package postgres

import (
	"database/sql"

	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/entity"
	"github.com/google/uuid"
)

// GetUserByID attempts to find a user with the specified ID.
// If the account does not exist a ErrNotFound error is returned.
func (p *postgresDatabase) GetUserByID(id uuid.UUID) (*entity.User, error) {
	const query = `
		SELECT id, email, password_hash FROM users
		WHERE id = $1
	`

	row := p.DB.QueryRow(query, id)
	var rowID uuid.UUID
	var rowEmail string
	var rowPasswordHash []byte
	if err := row.Scan(&rowID, &rowEmail, &rowPasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, database.ErrNotFound
		}
		return nil, err
	}
	return &entity.User{
		Email:        rowEmail,
		ID:           rowID,
		PasswordHash: rowPasswordHash,
	}, nil
}
