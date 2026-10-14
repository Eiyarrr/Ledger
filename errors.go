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
	// this error is functionally unreachable, it is simply here in case something breaks during future implementations
	errUnknownSeverity = errors.New("unknown severity when converting to string")
)
