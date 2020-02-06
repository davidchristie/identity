package core_test

import (
	"reflect"
	"testing"

	"github.com/davidchristie/identity/core"
	"github.com/davidchristie/identity/entity"
	"github.com/davidchristie/identity/mock"
	"github.com/golang/mock/gomock"
)

type logoutTestCase struct {
	DeleteSessionError error
	ExpectedError      error
	ExpectedOutput     *core.LogoutOutput
	Input              *core.LogoutInput
	ParseError         error
	ParseOutput        *entity.Session
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
		ParseError: nil,
		ParseOutput: &entity.Session{
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
				DeleteSession(testCase.ParseOutput.ID).
				Return(testCase.DeleteSessionError)
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
