//go:generate mockgen -destination ../mock/crypto.go -package mock github.com/davidchristie/identity/crypto Crypto

package crypto

// Crypto ...
type Crypto interface {
	GeneratePasswordHash(password string) ([]byte, error)
	IsCorrectPassword(password string, hash []byte) bool
}
