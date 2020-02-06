package bcrypt

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/davidchristie/identity/crypto"
)

type bcryptCrypto struct{}

const cost = 10

// New ...
func New() crypto.Crypto {
	return &bcryptCrypto{}
}

func (c *bcryptCrypto) GeneratePasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), cost)
}

func (c *bcryptCrypto) IsCorrectPassword(password string, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, []byte(password)) == nil
}
