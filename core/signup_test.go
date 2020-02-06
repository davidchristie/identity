package core_test

import (
	"reflect"
	"testing"

	"github.com/davidchristie/identity/core"
	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/entity"
	"github.com/davidchristie/identity/mock"
	"github.com/golang/mock/gomock"
)

type signupTestCase struct {
	CreateUserError   error
	CreateUserOutput  *entity.User
	ExpectedError     error
	ExpectedOutput    *core.SignupOutput
	Input             *core.SignupInput
	PasswordHash      []byte
	PasswordHashError error
}

var signupTestCases = []signupTestCase{
	// Successful signup
	signupTestCase{
		CreateUserOutput: &entity.User{
			ID:    uuid1,
			Email: email1,
		},
		ExpectedOutput: &core.SignupOutput{},
		Input: &core.SignupInput{
			Context:  context1,
			Email:    email1,
			Password: password1,
		},
		PasswordHash: hash1,
	},

	// Short password
	signupTestCase{
		ExpectedError: core.ErrShortPassword,
		Input: &core.SignupInput{
			Context:  context1,
			Email:    email1,
			Password: "123",
		},
	},

	// Duplicate user email
	signupTestCase{
		CreateUserError: database.ErrDuplicateUserEmail,
		ExpectedError:   core.ErrEmailAlreadyInUse,
		Input: &core.SignupInput{
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

	for _, testCase := range signupTestCases {
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

		core := core.New(core.Options{
			Crypto:   mockCrypto,
			Database: mockDatabase,
			Token:    mock.NewMockToken(ctrl),
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
