package handle

import (
	"net/http"
)

func handlePost(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handler(w, r)
		} else {
			w.WriteHeader(404)
			w.Write([]byte("404 page not found"))
		}
	}
}
