package acceptance

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// SignupResponseBody is the body of a signup response.
type SignupResponseBody struct{}

// SignupRequestBody is the body of a signup request.
type SignupRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SendSignupRequest sends a signup request to the identity service.
func SendSignupRequest(body SignupRequestBody) *http.Response {
	requestBodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", "http://localhost:8080/signup", bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	return response
}

// UnmarshalSignupResponseBody unmarshals the body of a signup response.
func UnmarshalSignupResponseBody(response *http.Response) SignupResponseBody {
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	body := SignupResponseBody{}
	err = json.Unmarshal(blob, &body)
	if err != nil {
		panic(err)
	}
	return body
}
