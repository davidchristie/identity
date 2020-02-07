package postgres

import (
	"database/sql"

	"github.com/davidchristie/identity/entity"
	"github.com/google/uuid"
)

// GetSessionByID attempts to find a session with the specified ID.
// If the account does not exist a ErrNotFound error is returned.
func (p *postgresDatabase) GetSessionByID(id uuid.UUID) (*entity.Session, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	const query = `
		SELECT id, user_id FROM sessions
		WHERE id = $1
	`

	row := db.QueryRow(query, id)
	var rowID uuid.UUID
	var rowUserID uuid.UUID
	if err := row.Scan(&rowID, &rowUserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &entity.Session{
		ID:     rowID,
		UserID: rowUserID,
	}, nil
}
