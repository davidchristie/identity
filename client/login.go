package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// loginResponseBody is the body of a login response.
type loginResponseBody struct {
	AccessToken string `json:"access_token"`
	Message     string `json:"message"`
}

// loginRequestBody is the body of a login request.
type loginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login sends a login request to the identity service.
func (c *client) Login(email string, password string) (*string, error) {
	requestBodyBytes, err := json.Marshal(&loginRequestBody{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("invalid response status: " + response.Status)
	}
	if response.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("invalid response content type: " + response.Header.Get("Content-Type"))
	}
	body := unmarshalLoginResponseBody(response)
	return &body.AccessToken, nil
}

// unmarshalLoginResponseBody unmarshals the body of a login response.
func unmarshalLoginResponseBody(response *http.Response) loginResponseBody {
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	body := loginResponseBody{}
	err = json.Unmarshal(blob, &body)
	if err != nil {
		panic(err)
	}
	return body
}
