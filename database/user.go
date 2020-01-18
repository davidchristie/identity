package database

// User ...
type User struct {
	CreatedAt    string
	DeletedAt    *string
	Email        string
	ID           string
	PasswordHash []byte
	UpdatedAt    string
}
