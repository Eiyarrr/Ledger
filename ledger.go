/*
Ledger provides a simple logging library.
*/

package ledger

import (
	"fmt"
	"io"
	"time"
)

/*
Ledger contains the information of the writer and minimum severity that will be recorded.
*/
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

/*
New creates a Ledger that writes to writer
with the minimum severity set to Info by default.
It returns ErrNoWriter if the writer is nil.
*/
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

/*
SetMinimumSeverity sets the minimum severity a Ledger will record.
It returns ErrInvalidSeverity if the severity is invalid and
leaves the existing minimum severity unchanged
*/
func (ledger *Ledger) SetMinimumSeverity(severity Severity) error {
	if severity.IsValid() {
		ledger.minimumSeverity = severity
		return nil
	}
	return ErrInvalidSeverity
}

/*
Record writes a timestamped log at a specified severity level.
Messages below the minimum severity are ignored and no error is returned.
It returns ErrInvalidSeverity if the severity is invalid and returns
an error from the writer if writing has failed.
*/
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
		return errUnknownSeverity
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
