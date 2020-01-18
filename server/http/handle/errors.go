package handle

import (
	"encoding/json"
	"github.com/davidchristie/identity/core"
	"net/http"
)

func writeError(err error, writer http.ResponseWriter) {
	switch err {
	case core.ErrEmailAlreadyInUse:
		writeErrorMessage(http.StatusBadRequest, "This email is already in use.", writer)

	default:
		writeErrorMessage(http.StatusInternalServerError, "An unknown error has occured.", writer)
	}
}

func writeErrorMessage(code int, message string, writer http.ResponseWriter) {
	data := struct {
		Message string `json:"message"`
	}{message}
	blob, _ := json.Marshal(data)
	writer.WriteHeader(code)
	writer.Write(blob)
}
