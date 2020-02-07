package handle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidchristie/identity/core"
)

type getUserResponseBody struct {
	Email string `json:"email"`
}

// User ...
func User(c core.Core) func(http.ResponseWriter, *http.Request) {
	return handleGet(func(writer http.ResponseWriter, request *http.Request) {
		authorization := request.Header.Get("Authorization")

		fmt.Println("AUTHORIZATION: " + authorization)

		if authorization == "" {
			writeErrorMessage(401, "No authorization header.", writer)
			return
		}

		authType := authorization[:6]

		if authType != "Bearer" {
			writeErrorMessage(401, "No bearer token.", writer)
			return
		}

		credentials := authorization[7:len(authorization)]

		request.Header.Get("Authorization")
		output, err := c.GetUser(&core.GetUserInput{
			AccessToken: credentials,
		})
		if err != nil {
			writeError(err, writer)
			return
		}

		response := &getUserResponseBody{
			Email: output.Email,
		}
		blob, _ := json.Marshal(response)
		addContentTypeJSONHeader(writer)
		writer.Write(blob)
	})
}
