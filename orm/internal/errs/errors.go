package errs

import (
	"errors"
	"fmt"
)

var (
	ErrPointerOnly = errors.New("orm: only pointer is supported")
	ErrNoRows      = errors.New("orm: no rows in result set")
)

func NewErrUnsupportedExpression(expr any) error {
	return fmt.Errorf("orm: unsupported expression %v", expr)
}

func NewErrInvalidTagContent(content string) error {
	return fmt.Errorf("orm: invalid tag content %s ", content)
}

func NewErrUnknownField(field string) error {
	return fmt.Errorf("orm: unknown field %s", field)
}
