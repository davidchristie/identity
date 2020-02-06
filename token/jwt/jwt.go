package jwt

import (
	"errors"
	"fmt"

	"github.com/davidchristie/identity/config"
	"github.com/davidchristie/identity/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type adapter struct {
	secret []byte
}

// New creates a new jwt-go adapter.
func New(c config.Token) token.Token {
	var adapter token.Token = &adapter{
		secret: []byte(c.Secret()),
	}
	return adapter
}

func (a *adapter) NewAccessToken(content *token.Content) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": content.ID.String(),
	})
	tokenString, err := token.SignedString(a.secret)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}

func (a *adapter) ParseAccessToken(tokenString string) (*token.Content, error) {
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.secret, nil
	})
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		if str, ok := claims["id"].(string); ok {
			id, err := uuid.Parse(str)
			if err != nil {
				return nil, err
			}
			return &token.Content{
				ID: id,
			}, nil
		}
		return nil, errors.New("No ID field in JWT")
	}
	return nil, err
}
