package core

import "errors"

// ErrEmailAlreadyInUse a user with this email already exists.
var ErrEmailAlreadyInUse = errors.New("email already in use")

// ErrWrongPassword the password is incorrect.
var ErrWrongPassword = errors.New("wrong password")

// ErrShortPassword the password is too short.
var ErrShortPassword = errors.New("short password")
