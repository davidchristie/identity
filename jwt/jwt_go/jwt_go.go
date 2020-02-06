package jwt_go

import (
	"errors"
	"fmt"

	"github.com/davidchristie/identity/config"
	"github.com/davidchristie/identity/entity"
	"github.com/davidchristie/identity/jwt"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type adapter struct {
	secret []byte
}

// New creates a new jwt-go adapter.
func New(c config.Token) jwt.JWT {
	var adapter jwt.JWT = &adapter{
		secret: []byte(c.Secret()),
	}
	return adapter
}

func (a *adapter) Parse(tokenString string) (*entity.AccessToken, error) {
	token, err := jwtGo.Parse(tokenString, func(token *jwtGo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.secret, nil
	})
	if claims, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		if str, ok := claims["id"].(string); ok {
			id, err := uuid.Parse(str)
			if err != nil {
				return nil, err
			}
			return &entity.AccessToken{
				ID: id,
			}, nil
		}
		return nil, errors.New("No ID field in JWT")
	}
	return nil, err
}

func (a *adapter) SignedString(input *jwt.SignedStringInput) (string, error) {
	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, jwtGo.MapClaims{
		"id": input.ID.String(),
	})
	tokenString, err := token.SignedString(a.secret)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}
