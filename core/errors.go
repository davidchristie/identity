package core

import "errors"

// ErrEmailAlreadyInUse a user with this email already exists.
var ErrEmailAlreadyInUse = errors.New("email already in use")
