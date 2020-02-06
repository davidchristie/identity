package postgres

import (
	"database/sql"

	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/entity"
	"github.com/google/uuid"
)

const createUserQuery = `INSERT INTO users (id, email, password_hash) VALUES ($1, $2, $3);`

func (p *postgresDatabase) CreateUser(input *database.CreateUserInput) (*entity.User, error) {
	if input.Context == nil {
		return nil, database.ErrNoContext
	}

	user := &entity.User{
		Email:        input.Email,
		ID:           uuid.New(),
		PasswordHash: input.PasswordHash,
	}

	tx, err := p.DB.BeginTx(input.Context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(createUserQuery, user.ID, user.Email, user.PasswordHash)
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
