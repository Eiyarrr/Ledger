package ledger

import (
	"errors"
	"fmt"
	"io"
	"time"
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
	writer          io.Writer
	minimumSeverity Severity
}

// struct containing information that needs to be recorded
type Entry struct {
	timestamp time.Time
	severity  Severity
	message   string
}

// returns an explicitly constructed object
func New(writer io.Writer) (*Ledger, error) {
	if writer == nil {
		return nil, errors.New("cannot initialize ledger without a writer")
	}

	ledger := Ledger{
		writer:          writer,
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
	return errors.New("invalid severity level")
}

// checks if severity is within the allowed bounds
func (severity Severity) IsValid() bool {
	return severity >= Info && severity <= Fatal
}

// writes the message, severity, and any metadata to a file
func (ledger *Ledger) Record(severity Severity, message string) error {
	if !severity.IsValid() {
		return errors.New("invalid severity level")
	}

	// if severity is lower than the minimum severity, ignore
	if severity < ledger.minimumSeverity {
		return nil
	}

	// fill Entry struct with information on the event
	entry := Entry{
		timestamp: time.Now(),
		severity:  severity,
		message:   message,
	}

	// convert severity to string
	severityString := entry.severity.String()
	if severityString == "" {
		return errors.New("unknown error converting severity to string")
	}

	// convert entry into a usable line
	line := fmt.Sprintf("[%s] | [%s] | %s", entry.timestamp, severityString, entry.message)

	// attempt to write the line
	_, err := ledger.writer.Write([]byte(line))
	if err != nil {
		return err
	}

	return nil
}

// string conversion of severity
func (severity Severity) String() string {
	switch severity {
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}

	return ""
}
