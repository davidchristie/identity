package acceptance_test

import (
	"reflect"
	"testing"

	"github.com/davidchristie/identity/test/acceptance"
	"github.com/google/uuid"
)

func TestRequest(t *testing.T) {
	email := "test.user+" + uuid.New().String() + "@email.com"
	password := uuid.New().String()

	// Signup
	signupResponse := acceptance.SendSignupRequest(acceptance.SignupRequestBody{
		Email:    email,
		Password: password,
	})

	t.Run("SignupResponseStatusCode", func(t *testing.T) {
		const expected = 200
		actual := signupResponse.StatusCode
		if actual != expected {
			t.Errorf("signup response status code = %d; expected %d", actual, expected)
		}
	})

	t.Run("SignupResponseContentTypeHeader", func(t *testing.T) {
		expected := []string{"application/json"}
		actual := signupResponse.Header["Content-Type"]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("signup response Content-Type header = %s; expected %s", actual, expected)
		}
	})

	t.Run("SignupResponseBody", func(t *testing.T) {
		expected := acceptance.SignupResponseBody{}
		actual := acceptance.UnmarshalSignupResponseBody(signupResponse)
		t.Logf("response body: %s", actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("signup response body = %s; expected %s", actual, expected)
		}
	})

	// Login
	loginResponse := acceptance.SendLoginRequest(acceptance.LoginRequestBody{
		Email:    email,
		Password: password,
	})

	t.Run("LoginResponseStatusCode", func(t *testing.T) {
		const expected = 200
		actual := loginResponse.StatusCode
		if actual != expected {
			t.Errorf("login response status code = %d; expected %d", actual, expected)
		}
	})

	t.Run("LoginResponseContentTypeHeader", func(t *testing.T) {
		expected := []string{"application/json"}
		actual := loginResponse.Header["Content-Type"]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("login response Content-Type header = %s; expected %s", actual, expected)
		}
	})

	body := acceptance.UnmarshalLoginResponseBody(loginResponse)
	t.Run("GetUserResponseBody", func(t *testing.T) {
		t.Logf("login response body: %s", body)
		if body.AccessToken == "" {
			t.Errorf("login response body does not contain access token")
		}
	})

	// Get user
	getUserResponse := acceptance.SendGetUserRequest(body.AccessToken)

	t.Run("GetUserResponseStatusCode", func(t *testing.T) {
		const expected = 200
		actual := getUserResponse.StatusCode
		if actual != expected {
			t.Errorf("get user response status code = %d; expected %d", actual, expected)
		}
	})

	t.Run("GetUserResponseContentTypeHeader", func(t *testing.T) {
		expected := []string{"application/json"}
		actual := getUserResponse.Header["Content-Type"]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("get user response Content-Type header = %s; expected %s", actual, expected)
		}
	})

	t.Run("GetUserResponseBody", func(t *testing.T) {
		expected := acceptance.GetUserResponseBody{
			Email: email,
		}
		actual := acceptance.UnmarshalGetUserResponseBody(getUserResponse)
		t.Logf("get user response body: %s", actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("get user response body = %s; expected %s", actual, expected)
		}
	})
}
