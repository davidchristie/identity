package acceptance_test

import (
	"testing"

	"github.com/davidchristie/identity/client"
	"github.com/google/uuid"
)

func TestRequest(t *testing.T) {
	identity := client.New(&client.Options{
		Host: "http://localhost:8080",
	})

	email := "test.user+" + uuid.New().String() + "@email.com"
	password := uuid.New().String()

	// Signup
	err := identity.Signup(email, password)
	if err != nil {
		t.Error(err)
	}

	// Login
	accessToken, err := identity.Login(email, password)
	if err != nil {
		t.Error(err)
	}
	if *accessToken == "" {
		t.Errorf("login response does not contain access token")
	}

	// Get user
	user, err := identity.GetUser(*accessToken)
	if err != nil {
		t.Error(err)
	}
	if user.Email() != email {
		t.Errorf("get user email = %s; expected %s", user.Email(), email)
	}
}
