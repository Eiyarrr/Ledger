package ledger

import (
	"fmt"
	"io"
	"time"
)

// struct containing configuration and state
type Ledger struct {
	writer          io.Writer
	minimumSeverity Severity
}

// struct containing information that needs to be recorded
type Entry struct {
	Timestamp time.Time
	Severity  Severity
	Message   string
}

// returns an explicitly constructed object
func New(writer io.Writer) (*Ledger, error) {
	if writer == nil {
		return nil, ErrNoWriter
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
	return ErrInvalidSeverity
}

// writes the message, severity, and any metadata to a file
func (ledger *Ledger) Record(severity Severity, message string) error {
	if !severity.IsValid() {
		return ErrInvalidSeverity
	}

	// if severity is lower than the minimum severity, ignore
	if severity < ledger.minimumSeverity {
		return nil
	}

	// fill Entry struct with information on the event
	entry := Entry{
		Timestamp: time.Now(),
		Severity:  severity,
		Message:   message,
	}

	// convert severity to string
	severityString := entry.Severity.String()
	if severityString == "" {
		return ErrUnknownSeverity
	}

	// format timestamp
	timestamp := entry.Timestamp.Format(time.RFC3339)

	// convert entry into a usable line
	line := fmt.Sprintf("[%s] | [%s] | %s\n", timestamp, severityString, entry.Message)

	// attempt to write the line
	bytesWritten, err := ledger.writer.Write([]byte(line))
	if err != nil {
		return err
	}
	if bytesWritten != len(line) {
		return ErrShortWrite
	}

	return nil
}
