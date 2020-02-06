package handle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidchristie/identity/core"
)

func writeError(err error, writer http.ResponseWriter) {
	fmt.Println(err)
	switch err {
	case core.ErrEmailAlreadyInUse:
		writeErrorMessage(http.StatusBadRequest, "This email is already in use.", writer)

	case core.ErrShortPassword:
		writeErrorMessage(http.StatusBadRequest, "This password is not long enough.", writer)

	default:
		writeErrorMessage(http.StatusInternalServerError, "An unknown error has occured.", writer)
	}
}

func writeErrorMessage(code int, message string, writer http.ResponseWriter) {
	data := struct {
		Message string `json:"message"`
	}{message}
	blob, _ := json.Marshal(data)
	addContentTypeJSONHeader(writer)
	writer.WriteHeader(code)
	writer.Write(blob)
}
