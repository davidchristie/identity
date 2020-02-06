package bcrypt_test

import (
	"testing"

	"github.com/davidchristie/identity/crypto/bcrypt"
	"github.com/golang/mock/gomock"
)

type IsCorrectPasswordTestCase struct {
	ExpectedOutput bool
	Password       string
	PasswordHash   []byte
}

var IsCorrectPasswordTestCases = []IsCorrectPasswordTestCase{
	IsCorrectPasswordTestCase{
		ExpectedOutput: true,
		Password:       "385a1ec0-2016-413a-bfdf-181ce0a38a42",
		PasswordHash:   []byte("$2a$10$FeUt6nt8bXqXDbJoUGU9qeLTi6K6RvXgNMJcv6FoqgyoBi6KOVir22"),
	},
}

func TestIsCorrectPassword(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	crypto := bcrypt.New()

	for _, testCase := range IsCorrectPasswordTestCases {
		output := crypto.IsCorrectPassword(testCase.Password, testCase.PasswordHash)
		if output != testCase.ExpectedOutput {
			t.Errorf("Invalid output: actual=%v, expected=%v", output, testCase.ExpectedOutput)
		}
	}
}
