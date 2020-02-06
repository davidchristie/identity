package postgres

import (
	"database/sql"

	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/entity"
	"github.com/google/uuid"
)

const createAccessTokenQuery = `INSERT INTO access_tokens (id, user_id) VALUES ($1, $2);`

func (p *postgresDatabase) CreateAccessToken(input *database.CreateAccessTokenInput) (*entity.AccessToken, error) {
	if input.Context == nil {
		return nil, database.ErrNoContext
	}

	token := &entity.AccessToken{
		ID:     uuid.New(),
		UserID: input.UserID,
	}

	tx, err := p.DB.BeginTx(input.Context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(createAccessTokenQuery, token.ID, token.UserID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return token, nil
}
