package errors

import (
	"errors"
	"fmt"
)

var (
	ErrInternalServerError = errors.New("internal server error")

	ErrShortURLNotFound = errors.New("short URL not found")
)

func NewCustomInternalErr(err error) error {
	return fmt.Errorf("%w: %w", ErrInternalServerError, err)
}
