package core

import (
	"context"
	"reflect"
	"testing"

	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

type testCase struct {
	CreateUserOutput  *database.User
	CreateUserError   error
	ExpectedOutput    *SignupOutput
	ExpectedError     error
	Input             *SignupInput
	MockDatabase      func() database.Database
	PasswordHash      []byte
	PasswordHashError error
}

const email1 = "user@email.com"
const password1 = "pa$$word123"

var context1 = context.Background()
var uuid1, _ = uuid.Parse("625af883-21ee-40d3-bc40-a753cece2f60")
var hash1 = []byte("$2a$10$gYXXJulMpoUalXFgmOpKbO6v.nigV2lWf/Z3EwgykLdGzekwGfAbW")

var testCases = []testCase{
	// Successful signup
	testCase{
		CreateUserOutput: &database.User{
			ID:    uuid1,
			Email: email1,
		},
		ExpectedOutput: &SignupOutput{},
		Input: &SignupInput{
			Context:  context1,
			Email:    email1,
			Password: password1,
		},
		PasswordHash: hash1,
	},

	// Short password
	testCase{
		ExpectedError: ErrShortPassword,
		Input: &SignupInput{
			Context:  context1,
			Email:    email1,
			Password: "123",
		},
	},

	// Duplicate user email
	testCase{
		CreateUserError: database.ErrDuplicateUserEmail,
		ExpectedError:   ErrEmailAlreadyInUse,
		Input: &SignupInput{
			Context:  context1,
			Email:    email1,
			Password: password1,
		},
		PasswordHash: hash1,
	},
}

func TestSignup(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	for _, testCase := range testCases {
		mockCrypto := mock.NewMockCrypto(ctrl)
		if testCase.PasswordHash != nil || testCase.PasswordHashError != nil {
			mockCrypto.
				EXPECT().
				GeneratePasswordHash(testCase.Input.Password).
				Return(testCase.PasswordHash, testCase.PasswordHashError)
		}

		mockDatabase := mock.NewMockDatabase(ctrl)
		if testCase.CreateUserOutput != nil || testCase.CreateUserError != nil {
			mockDatabase.
				EXPECT().
				CreateUser(&database.CreateUserInput{
					Context:      testCase.Input.Context,
					Email:        testCase.Input.Email,
					PasswordHash: testCase.PasswordHash,
				}).
				Return(testCase.CreateUserOutput, testCase.CreateUserError)
		}

		core := New(Options{
			Crypto:   mockCrypto,
			Database: mockDatabase,
		})

		output, err := core.Signup(testCase.Input)
		if err != testCase.ExpectedError {
			t.Errorf("Invalid error: actual=%v, expected=%v", err, testCase.ExpectedError)
		}
		if !reflect.DeepEqual(output, testCase.ExpectedOutput) {
			t.Errorf("Invalid output: actual=%v, expected=%v", output, testCase.ExpectedOutput)
		}
	}
}
