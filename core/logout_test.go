package core_test

import (
	"reflect"
	"testing"

	"github.com/davidchristie/identity/core"
	"github.com/davidchristie/identity/database"
	"github.com/davidchristie/identity/mock"
	"github.com/golang/mock/gomock"
)

type logoutTestCase struct {
	DeleteAccessTokenError error
	ExpectedError          error
	ExpectedOutput         *core.LogoutOutput
	Input                  *core.LogoutInput
	ParseError             error
	ParseOutput            *database.AccessToken
}

var logoutTestCases = []logoutTestCase{
	// Successful logout
	logoutTestCase{
		DeleteAccessTokenError: nil,
		ExpectedError:          nil,
		ExpectedOutput:         &core.LogoutOutput{},
		Input: &core.LogoutInput{
			AccessToken: jwt1,
		},
		ParseError: nil,
		ParseOutput: &database.AccessToken{
			ID: uuid1,
		},
	},
}

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	for _, testCase := range logoutTestCases {

		mockDatabase := mock.NewMockDatabase(ctrl)
		if testCase.ParseOutput != nil {
			mockDatabase.
				EXPECT().
				DeleteAccessToken(testCase.ParseOutput.ID).
				Return(testCase.DeleteAccessTokenError)
		}

		mockJWT := mock.NewMockJWT(ctrl)
		mockJWT.
			EXPECT().
			Parse(testCase.Input.AccessToken).
			Return(testCase.ParseOutput, testCase.ParseError)

		core := core.New(core.Options{
			Crypto:   mock.NewMockCrypto(ctrl),
			Database: mockDatabase,
			JWT:      mockJWT,
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
