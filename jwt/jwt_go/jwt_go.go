package jwt_go

import (
	"errors"
	"fmt"

	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/jwt"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// TODO: Get this secret from environment.
var hmacSampleSecret = []byte("YBgJ7-QtP8RNW7QkNd7o")

type jwtGoAdapter struct{}

// New creates a new jwt-go adapter.
func New() jwt.JWT {
	var adapter jwt.JWT = &jwtGoAdapter{}
	return adapter
}

func (j *jwtGoAdapter) Parse(tokenString string) (*database.AccessToken, error) {
	token, err := jwtGo.Parse(tokenString, func(token *jwtGo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if claims, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		if str, ok := claims["id"].(string); ok {
			id, err := uuid.Parse(str)
			if err != nil {
				return nil, err
			}
			return &database.AccessToken{
				ID: id,
			}, nil
		}
		return nil, errors.New("No ID field in JWT")
	}
	return nil, err
}

func (j *jwtGoAdapter) SignedString(input *jwt.SignedStringInput) (string, error) {
	fmt.Println("SignedString 1")

	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, jwtGo.MapClaims{
		"id": input.ID.String(),
	})

	fmt.Println("SignedString 2")

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("SignedString 3")

	return tokenString, nil
}
