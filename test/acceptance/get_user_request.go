package acceptance

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetUserResponseBody is the body of a get user response.
type GetUserResponseBody struct {
	Email string `json:"email"`
}

// SendGetUserRequest sends a get user request to the identity service.
func SendGetUserRequest(accessToken string) *http.Response {
	request, err := http.NewRequest("GET", "http://localhost:8080/user", bytes.NewBuffer([]byte("")))
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	return response
}

// UnmarshalGetUserResponseBody unmarshals the body of a get user response.
func UnmarshalGetUserResponseBody(response *http.Response) GetUserResponseBody {
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	body := GetUserResponseBody{}
	err = json.Unmarshal(blob, &body)
	if err != nil {
		panic(err)
	}
	return body
}
