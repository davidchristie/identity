package handle

import (
	"encoding/json"
	"net/http"
)

func writeError(err error, writer http.ResponseWriter) {
	switch err {
	default:
		writeErrorMessage(500, "An unknown error has occured.", writer)
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
