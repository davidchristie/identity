package env

import (
	"github.com/davidchristie/identity/config"
)

type env struct {
	database *database
	server   *server
	token    *token
}

func New() config.Config {
	return &env{
		database: &database{},
		server:   &server{},
		token:    &token{},
	}
}

func (e *env) Database() config.Database {
	return e.database
}

func (e *env) Server() config.Server {
	return e.server
}

func (e *env) Token() config.Token {
	return e.token
}
