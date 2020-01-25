package http

import (
	"fmt"
	"net/http"

	"github.com/davidchristie/identity/core"
	"github.com/davidchristie/identity/server"
	"github.com/davidchristie/identity/server/http/handle"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
}

// Options ...
type Options struct {
	Address string
}

type httpServer struct {
	Address string
}

// New ...
func New(options Options) server.Server {
	fmt.Println("Creating HTTP server...")
	h := &httpServer{
		Address: options.Address,
	}
	if h.Address == "" {
		h.Address = ":8080"
	}
	return h
}

func (h *httpServer) Serve(core core.Core) error {
	fmt.Println("Starting HTTP server...")
	routes := http.NewServeMux()
	routes.HandleFunc("/login", handle.Login(core))
	routes.HandleFunc("/logout", handle.Logout(core))
	routes.HandleFunc("/verify", handle.Verify(core))
	routes.HandleFunc("/signup", handle.Signup(core))
	return http.ListenAndServe(h.Address, routes)
}