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
type entry struct {
	timestamp time.Time
	severity  Severity
	message   string
}

// New() creates a Ledger that writes to writer.
// It returns ErrNoWriter if the writer is nil.
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

// SetMinimumSeverity() changes the Ledger.minimumSeverity field.
// If the severity of a Record() call is below Ledger.minimumSeverity, it is ignored.
func (ledger *Ledger) SetMinimumSeverity(severity Severity) error {
	if severity.IsValid() {
		ledger.minimumSeverity = severity
		return nil
	}
	return ErrInvalidSeverity
}

// Record() writes into the Ledger.writer the timestamp, severity, and message.
func (ledger *Ledger) Record(severity Severity, message string) error {
	if !severity.IsValid() {
		return ErrInvalidSeverity
	}

	// if severity is lower than the minimum severity, ignore
	if severity < ledger.minimumSeverity {
		return nil
	}

	// fill Entry struct with information on the event
	entry := entry{
		timestamp: time.Now(),
		severity:  severity,
		message:   message,
	}

	// convert severity to string
	severityString := entry.severity.String()
	if severityString == "" {
		return ErrUnknownSeverity
	}

	// format timestamp
	timestamp := entry.timestamp.Format(time.RFC3339)

	// convert entry into a usable line
	line := fmt.Sprintf("[%s] | [%s] | %s\n", timestamp, severityString, entry.message)

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
