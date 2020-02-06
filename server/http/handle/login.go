package handle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidchristie/identity/core"
)

type loginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponseBody struct {
	AccessToken string `json:"access_token"`
}

// Login ...
func Login(c core.Core) func(http.ResponseWriter, *http.Request) {
	return handlePost(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Login request received")

		body := &loginRequestBody{}

		err := json.NewDecoder(request.Body).Decode(&body)
		if err != nil {
			writeError(err, writer)
			return
		}

		output, err := c.Login(&core.LoginInput{
			Context:  request.Context(),
			Email:    body.Email,
			Password: body.Password,
		})
		if err != nil {
			writeError(err, writer)
			return
		}

		response := &loginResponseBody{
			AccessToken: output.AccessToken,
		}

		fmt.Println("Sending login response")

		blob, _ := json.Marshal(response)
		addContentTypeJSONHeader(writer)
		writer.Write(blob)
	})
}
