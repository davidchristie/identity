package core

import "errors"

// ErrEmailAlreadyInUse a user with this email already exists.
var ErrEmailAlreadyInUse = errors.New("email already in use")

// ErrEmailNotFound the email was not found.
var ErrEmailNotFound = errors.New("email not found")

// ErrWrongPassword the password is incorrect.
var ErrWrongPassword = errors.New("wrong password")

// ErrShortPassword the password is too short.
var ErrShortPassword = errors.New("short password")
