package ledger

import "errors"

// define ledger-specific errors
var (
	ErrInvalidSeverity = errors.New("invalid severity level")
	ErrNoWriter        = errors.New("cannot use ledger without a writer")
)
