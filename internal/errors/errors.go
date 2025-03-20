package errors

import (
	"errors"
	"fmt"
)

var (
	ErrInternalServerError = errors.New("internal server error")

	ErrShortURLNotFound      = errors.New("short URL not found")
	ErrShortURLAlreadyExists = errors.New("short URL already exists")
)

func NewCustomInternalErr(err error) error {
	return fmt.Errorf("%w: %w", ErrInternalServerError, err)
}
