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
	wrongPassword := "wrong_password"

	// Not signed up
	accessToken, err := identity.Login(email, password)
	if accessToken != nil {
		t.Errorf("access token returned when user is not signed up")
	}
	if err == nil {
		t.Errorf("no error when user is not signed up")
	}
	if err.Error() != "This email doesn't belong to an account." {
		t.Errorf("wrong error message: %s", err)
	}

	// Signup
	err = identity.Signup(email, password)
	if err != nil {
		t.Error(err)
	}

	// Wrong password
	accessToken, err = identity.Login(email, wrongPassword)
	if accessToken != nil {
		t.Errorf("access token returned when password is wrong")
	}
	if err == nil {
		t.Errorf("no error when password is wrong")
	}
	if err.Error() != "Incorrect password." {
		t.Errorf("wrong error message: %s", err)
	}

	// Login
	accessToken, err = identity.Login(email, password)
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
