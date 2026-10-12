package ledger

import (
	"errors"
	"io"
)

// define ledger-specific errors
var (
	ErrInvalidSeverity = errors.New("invalid severity level")
	ErrNoWriter        = errors.New("cannot use ledger without a writer")
	ErrShortWrite      = io.ErrShortWrite
	ErrUnknownSeverity = errors.New("unknown severity when converting to string")
)
