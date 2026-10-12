package ledger

import (
	"errors"
	"io"
)

// define ledger severity levels
type Severity int

const (
	Info  Severity = iota //0
	Warn                  //1
	Error                 //2
	Fatal                 //3
)

// struct containing configuration and state
type Ledger struct {
	writer           io.Writer
	minimumSeverity Severity
}

// returns an explicitly constructed object
func New(writer io.Writer) (*Ledger, error) {
	if writer == nil {
		return nil, errors.New("cannot initialize ledger without a writer")
	}

	ledger := Ledger{
		writer:           writer,
		minimumSeverity: Info,
	}
	return &ledger, nil
}

// allows changing of minimum severity within an existing Ledger
// if severity is below the minimum, it is not logged
func (ledger *Ledger) SetMinimumSeverity(severity Severity) error {
	if severity.IsValid() {
		ledger.minimumSeverity = severity
		return nil
	}
	return errors.New("cannot set severity to a level outside the range")
}

// checks if severity is within the allowed bounds
func (severity Severity) IsValid() bool {
	return severity >= Info && severity <= Fatal
}
