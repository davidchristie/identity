package handle

import (
	"encoding/json"
	"net/http"

	"github.com/davidchristie/identity/core"
)

// Login ...
func Login(c core.Core) func(http.ResponseWriter, *http.Request) {
	return handlePost(func(writer http.ResponseWriter, request *http.Request) {
		_, err := c.Login(&core.LoginInput{})
		if err != nil {
			writeError(err, writer)
			return
		}
		response := &struct{}{}
		blob, _ := json.Marshal(response)
		writer.Write(blob)
	})
}
