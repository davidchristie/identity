package handle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidchristie/identity/core"
)

type signupRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Signup ...
func Signup(c core.Core) func(http.ResponseWriter, *http.Request) {
	return handlePost(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Signup request received")

		body := &signupRequestBody{}

		err := json.NewDecoder(request.Body).Decode(&body)
		if err != nil {
			writeError(err, writer)
			return
		}

		_, err = c.Signup(&core.SignupInput{
			Context:  request.Context(),
			Email:    body.Email,
			Password: body.Password,
		})
		if err != nil {
			writeError(err, writer)
			return
		}

		response := &struct{}{}
		blob, _ := json.Marshal(response)
		writer.Write(blob)
	})
}
