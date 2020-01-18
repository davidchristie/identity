package core

import (
	"testing"

	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/mock"
	"github.com/golang/mock/gomock"
)

func TestSignup(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	email := "user@email.com"
	password := "pa$$word123"
	passwordHash := []byte("$2a$10$gYXXJulMpoUalXFgmOpKbO6v.nigV2lWf/Z3EwgykLdGzekwGfAbW")

	mockCrypto := mock.NewMockCrypto(ctrl)
	mockCrypto.
		EXPECT().
		GeneratePasswordHash(password).
		Return(passwordHash, nil)

	mockDatabase := mock.NewMockDatabase(ctrl)
	mockDatabase.
		EXPECT().
		CreateUser(&database.CreateUserInput{
			Email:        email,
			PasswordHash: passwordHash,
		}).
		Return(&database.User{
			Email: email,
		}, nil)

	core := New(Options{
		Crypto:   mockCrypto,
		Database: mockDatabase,
	})
	output, err := core.Signup(&SignupInput{
		Email:    email,
		Password: password,
	})
	expectedOutput := &SignupOutput{}

	if err != nil {
		t.Error("Error: ", err)
	}
	if output != expectedOutput {
		t.Errorf("Invalid output - actual: %v, expected: %v", output, expectedOutput)
	}
}
