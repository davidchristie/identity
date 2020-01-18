package server

import "github.com/davidchristie/identity/core"

// Server ...
type Server interface {
	Serve(core.Core) error
}
