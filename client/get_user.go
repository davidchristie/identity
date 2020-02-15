package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// getUserResponseBody is the body of a get user response.
type getUserResponseBody struct {
	Email string `json:"email"`
}

// GetUser sends a get user request to the identity service.
func (c *client) GetUser(accessToken string) (User, error) {
	request, err := http.NewRequest("GET", "http://localhost:8080/user", bytes.NewBuffer([]byte("")))
	request.Header.Set("Authorization", "Bearer "+accessToken)
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
	body, err := unmarshalGetUserResponseBody(response)
	if err != nil {
		return nil, err
	}
	return &user{
		email: body.Email,
	}, nil
}

// unmarshalGetUserResponseBody unmarshals the body of a get user response.
func unmarshalGetUserResponseBody(response *http.Response) (*getUserResponseBody, error) {
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := &getUserResponseBody{}
	err = json.Unmarshal(blob, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
