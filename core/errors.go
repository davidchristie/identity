package core

import "errors"

// ErrEmailAlreadyInUse a user with this email already exists.
var ErrEmailAlreadyInUse = errors.New("email already in use")

// ErrShortPassword the password is too short.
var ErrShortPassword = errors.New("short password")
