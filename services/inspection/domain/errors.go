package domain

import "errors"

var (
	ErrNotFound   = errors.New("vehicle not found")
	ErrValidation = errors.New("validation failed")
)
