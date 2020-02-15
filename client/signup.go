package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// signupResponseBody is the body of a signup response.
type signupResponseBody struct{}

// signupRequestBody is the body of a signup request.
type signupRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Signup sends a signup request to the identity service.
func (c *client) Signup(email string, password string) error {
	requestBodyBytes, err := json.Marshal(&signupRequestBody{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", c.host+"/signup", bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New("invalid response status: " + response.Status)
	}
	if response.Header.Get("Content-Type") != "application/json" {
		return errors.New("invalid response content type: " + response.Header.Get("Content-Type"))
	}
	return nil
}

// unmarshalSignupResponseBody unmarshals the body of a signup response.
func unmarshalSignupResponseBody(response *http.Response) signupResponseBody {
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	body := signupResponseBody{}
	err = json.Unmarshal(blob, &body)
	if err != nil {
		panic(err)
	}
	return body
}
