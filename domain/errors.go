package domain

import (
	"errors"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrBadParamInput       = errors.New("given param is not valid")
)

var (
	ErrUserNotFound          = errors.New("user is not found")
	ErrUserConflict          = errors.New("username exist")
	ErrUserWrongPassword     = errors.New("incorrect password")
	ErrUserDuplicateUsername = errors.New("user already exist")
)
