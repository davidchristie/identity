//go:generate mockgen -destination ../mock/token.go -package mock github.com/davidchristie/identity/token Token

package token

import (
	"github.com/google/uuid"
)

type Content struct {
	ID uuid.UUID
}

// Token ...
type Token interface {
	NewAccessToken(*Content) (string, error)
	ParseAccessToken(string) (*Content, error)
}
