package postgres

import (
	"context"
	"database/sql"

	"github.com/davidchristie/identity/database"
	"github.com/google/uuid"
)

func (p *postgresDatabase) CreateUser(input *database.CreateUserInput) (*database.User, error) {
	if input.Context == nil {
		return nil, database.ErrNoContext
	}

	user := &database.User{
		Email:        input.Email,
		ID:           uuid.New(),
		PasswordHash: input.PasswordHash,
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	const query = `
		INSERT INTO users (id, email, password_hash)
		VALUES ($1, $2, $3);
	`

	_, err = tx.Exec(query, user.ID, user.Email, user.PasswordHash)
	if err != nil {
		tx.Rollback()
		if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
			return nil, database.ErrDuplicateUserEmail
		}
		return nil, err
	}

	tx.Commit()

	return user, nil
}
