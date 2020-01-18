package database

import "errors"

// ErrDuplicateUserEmail a user with email already exists.
var ErrDuplicateUserEmail = errors.New("duplicate user email")

// ErrNoContext no context was specified.
var ErrNoContext = errors.New("no context")
