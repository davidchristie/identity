package core_test

import (
	"reflect"
	"testing"

	"github.com/davidchristie/identity/core"
	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/entity"
	"github.com/davidchristie/identity/mock"
	"github.com/davidchristie/identity/token"
	"github.com/golang/mock/gomock"
)

type loginTestCase struct {
	CreateSessionError   error
	CreateSessionOutput  *entity.Session
	ExpectedOutput       *core.LoginOutput
	ExpectedError        error
	GetUserByEmailError  error
	GetUserByEmailOutput *entity.User
	Input                *core.LoginInput
	IsCorrectPassword    *bool
}

var loginTestCases = []loginTestCase{
	// Successful login
	loginTestCase{
		CreateSessionOutput: &entity.Session{
			ID: uuid2,
		},
		ExpectedError: nil,
		ExpectedOutput: &core.LoginOutput{
			AccessToken: jwt1,
		},
		GetUserByEmailError: nil,
		GetUserByEmailOutput: &entity.User{
			ID:           uuid1,
			PasswordHash: hash1,
		},
		Input: &core.LoginInput{
			Context:  context1,
			Email:    email1,
			Password: password1,
		},
		IsCorrectPassword: &[]bool{true}[0],
	},

	// Email not found
	loginTestCase{
		ExpectedOutput:       nil,
		ExpectedError:        core.ErrEmailNotFound,
		GetUserByEmailError:  database.ErrNotFound,
		GetUserByEmailOutput: nil,
		Input: &core.LoginInput{
			Context:  context1,
			Email:    email1,
			Password: password1,
		},
	},

	// Wrong password
	loginTestCase{
		ExpectedOutput:      nil,
		ExpectedError:       core.ErrWrongPassword,
		GetUserByEmailError: nil,
		GetUserByEmailOutput: &entity.User{
			ID:           uuid1,
			PasswordHash: hash1,
		},
		Input: &core.LoginInput{
			Context:  context1,
			Email:    email1,
			Password: password2,
		},
		IsCorrectPassword: &[]bool{false}[0],
	},
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	for _, testCase := range loginTestCases {
		mockCrypto := mock.NewMockCrypto(ctrl)
		if testCase.IsCorrectPassword != nil {
			mockCrypto.
				EXPECT().
				IsCorrectPassword(testCase.Input.Password, testCase.GetUserByEmailOutput.PasswordHash).
				Return(*testCase.IsCorrectPassword)
		}

		mockDatabase := mock.NewMockDatabase(ctrl)
		if testCase.GetUserByEmailOutput != nil || testCase.GetUserByEmailError != nil {
			mockDatabase.
				EXPECT().
				GetUserByEmail(&database.GetUserByEmailInput{
					Context: testCase.Input.Context,
					Email:   testCase.Input.Email,
				}).
				Return(testCase.GetUserByEmailOutput, testCase.GetUserByEmailError)
		}
		if testCase.CreateSessionOutput != nil || testCase.CreateSessionError != nil {
			mockDatabase.
				EXPECT().
				CreateSession(&database.CreateSessionInput{
					Context: testCase.Input.Context,
					UserID:  testCase.GetUserByEmailOutput.ID,
				}).
				Return(testCase.CreateSessionOutput, testCase.CreateSessionError)
		}

		mockToken := mock.NewMockToken(ctrl)
		if testCase.CreateSessionOutput != nil {
			mockToken.
				EXPECT().
				NewAccessToken(&token.Content{
					ID: testCase.CreateSessionOutput.ID,
				}).
				Return(jwt1, nil)
		}

		core := core.New(core.Options{
			Crypto:   mockCrypto,
			Database: mockDatabase,
			Token:    mockToken,
		})

		output, err := core.Login(testCase.Input)
		if err != testCase.ExpectedError {
			t.Errorf("Invalid error: actual=%v, expected=%v", err, testCase.ExpectedError)
		}
		if !reflect.DeepEqual(output, testCase.ExpectedOutput) {
			t.Errorf("Invalid output: actual=%v, expected=%v", output, testCase.ExpectedOutput)
		}
	}
}
