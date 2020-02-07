package acceptance

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// LoginResponseBody is the body of a login response.
type LoginResponseBody struct {
	AccessToken string `json:"access_token"`
	Message     string `json:"message"`
}

// LoginRequestBody is the body of a login request.
type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SendLoginRequest sends a login request to the identity service.
func SendLoginRequest(body LoginRequestBody) *http.Response {
	requestBodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	return response
}

// UnmarshalLoginResponseBody unmarshals the body of a login response.
func UnmarshalLoginResponseBody(response *http.Response) LoginResponseBody {
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	body := LoginResponseBody{}
	err = json.Unmarshal(blob, &body)
	if err != nil {
		panic(err)
	}
	return body
}
