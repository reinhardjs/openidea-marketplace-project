package domain

import (
	"errors"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("user is not found")
	ErrConflict            = errors.New("username exist")
	ErrBadParamInput       = errors.New("given param is not valid")
)
