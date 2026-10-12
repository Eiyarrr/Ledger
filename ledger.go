package ledger

import (
	"errors"
	"io"
)

// define ledger severity levels
type Severity int

const (
	INFO  Severity = iota //0
	WARN                  //1
	ERROR                 //2
	FATAL                 //3
)

// struct containing configuration and state
type Ledger struct {
	writer           io.Writer
	minimum_severity Severity
}

// returns an explicitly constructed object
func New(writer io.Writer) (*Ledger, error) {
	if writer == nil {
		return nil, errors.New("cannot initialize ledger without a writer")
	}
	
	ledger := Ledger{
		writer: writer,
		minimum_severity: INFO,
	}
	return &ledger, nil
}

// allows changing of minimum severity within an existing Ledger
// if severity is below the minimum, it is not logged
func (ledger *Ledger) MinimumSeverity(severity Severity) {
	ledger.minimum_severity = severity
}
