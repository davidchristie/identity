package config

type Config interface {
	Database() Database
	Server() Server
	Token() Token
}

type Database interface {
	Host() string
	Name() string
	Password() string
	Username() string
}

type Server interface {
	Port() string
}

type Token interface {
	Secret() string
}
