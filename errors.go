package ledger

import (
	"errors"
	"io"
)

/*
Errors returned by Ledger operations
*/
var (
	// ErrInvalidSeverity indicates that a severity value is invalid.
	ErrInvalidSeverity = errors.New("invalid severity level")

	// ErrNoWriter indicates that New was called with a nil writer.
	ErrNoWriter = errors.New("cannot use ledger without a writer")

	// ErrShortWrite indicates that a writer wrote fewer bytes than requested.
	ErrShortWrite = io.ErrShortWrite

	// errUnknownSeverity indicates that there was an issue internally when converting severity to a string.
	// This error is currently unreachable through normal API usage.
	errUnknownSeverity = errors.New("unknown severity when converting to string")
)
