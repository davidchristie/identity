package http

import (
	"fmt"
	"net/http"

	"github.com/davidchristie/identity/config"
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
func New(options config.Server) server.Server {
	fmt.Println("Creating HTTP server...")
	return &httpServer{
		Address: ":" + options.Port(),
	}
}

func (h *httpServer) Serve(core core.Core) error {
	fmt.Println("Starting HTTP server...")
	routes := http.NewServeMux()
	routes.HandleFunc("/login", handle.Login(core))
	routes.HandleFunc("/logout", handle.Logout(core))
	routes.HandleFunc("/signup", handle.Signup(core))
	routes.HandleFunc("/user", handle.User(core))
	return http.ListenAndServe(h.Address, routes)
}
