package core_test

import (
	"reflect"
	"testing"

	"github.com/davidchristie/identity/core"
	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/jwt"
	"github.com/davidchristie/identity/mock"
	"github.com/golang/mock/gomock"
)

type loginTestCase struct {
	CreateAccessTokenError  error
	CreateAccessTokenOutput *database.AccessToken
	ExpectedOutput          *core.LoginOutput
	ExpectedError           error
	GetUserByEmailError     error
	GetUserByEmailOutput    *database.User
	Input                   *core.LoginInput
	IsCorrectPassword       bool
}

var loginTestCases = []loginTestCase{
	// Successful signup
	loginTestCase{
		CreateAccessTokenOutput: &database.AccessToken{
			ID: uuid2,
		},
		ExpectedOutput: &core.LoginOutput{
			AccessToken: jwt1,
		},
		GetUserByEmailError: nil,
		GetUserByEmailOutput: &database.User{
			ID:           uuid1,
			PasswordHash: hash1,
		},
		Input: &core.LoginInput{
			Context:  context1,
			Email:    email1,
			Password: password1,
		},
		IsCorrectPassword: true,
	},
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	for _, testCase := range loginTestCases {
		mockCrypto := mock.NewMockCrypto(ctrl)
		mockCrypto.
			EXPECT().
			IsCorrectPassword(testCase.Input.Password, testCase.GetUserByEmailOutput.PasswordHash).
			Return(testCase.IsCorrectPassword)

		mockDatabase := mock.NewMockDatabase(ctrl)
		mockDatabase.
			EXPECT().
			GetUserByEmail(&database.GetUserByEmailInput{
				Context: testCase.Input.Context,
				Email:   testCase.Input.Email,
			}).
			Return(testCase.GetUserByEmailOutput, testCase.GetUserByEmailError)
		mockDatabase.
			EXPECT().
			CreateAccessToken(&database.CreateAccessTokenInput{
				Context: testCase.Input.Context,
				UserID:  testCase.GetUserByEmailOutput.ID,
			}).
			Return(testCase.CreateAccessTokenOutput, testCase.CreateAccessTokenError)

		mockJWT := mock.NewMockJWT(ctrl)
		mockJWT.
			EXPECT().
			SignedString(&jwt.SignedStringInput{
				ID: testCase.CreateAccessTokenOutput.ID,
			}).
			Return(jwt1, nil)

		core := core.New(core.Options{
			Crypto:   mockCrypto,
			Database: mockDatabase,
			JWT:      mockJWT,
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
