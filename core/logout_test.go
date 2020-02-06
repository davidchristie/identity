package core_test

import (
	"reflect"
	"testing"

	"github.com/davidchristie/identity/core"
	"github.com/davidchristie/identity/mock"
	"github.com/davidchristie/identity/token"
	"github.com/golang/mock/gomock"
)

type logoutTestCase struct {
	DeleteSessionError     error
	ExpectedError          error
	ExpectedOutput         *core.LogoutOutput
	Input                  *core.LogoutInput
	ParseAccessTokenError  error
	ParseAccessTokenOutput *token.Content
}

var logoutTestCases = []logoutTestCase{
	// Successful logout
	logoutTestCase{
		DeleteSessionError: nil,
		ExpectedError:      nil,
		ExpectedOutput:     &core.LogoutOutput{},
		Input: &core.LogoutInput{
			AccessToken: jwt1,
		},
		ParseAccessTokenError: nil,
		ParseAccessTokenOutput: &token.Content{
			ID: uuid1,
		},
	},
}

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	for _, testCase := range logoutTestCases {

		mockDatabase := mock.NewMockDatabase(ctrl)
		if testCase.ParseAccessTokenOutput != nil {
			mockDatabase.
				EXPECT().
				DeleteSession(testCase.ParseAccessTokenOutput.ID).
				Return(testCase.DeleteSessionError)
		}

		mockToken := mock.NewMockToken(ctrl)
		mockToken.
			EXPECT().
			ParseAccessToken(testCase.Input.AccessToken).
			Return(testCase.ParseAccessTokenOutput, testCase.ParseAccessTokenError)

		core := core.New(core.Options{
			Crypto:   mock.NewMockCrypto(ctrl),
			Database: mockDatabase,
			Token:    mockToken,
		})

		output, err := core.Logout(testCase.Input)
		if err != testCase.ExpectedError {
			t.Errorf("Invalid error: actual=%v, expected=%v", err, testCase.ExpectedError)
		}
		if !reflect.DeepEqual(output, testCase.ExpectedOutput) {
			t.Errorf("Invalid output: actual=%v, expected=%v", output, testCase.ExpectedOutput)
		}
	}
}
