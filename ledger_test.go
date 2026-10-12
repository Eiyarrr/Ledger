package ledger_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	ledger "github.com/Eiyarrr/Ledger"
)

func TestRecordAllSeverities(t *testing.T) {
	// create byte buffer (implements io.Writer but into memory)
	buffer := bytes.NewBuffer([]byte(""))

	// create ledger
	scribe, err := ledger.New(buffer)
	if err != nil {
		t.Fatal("failed to create new ledger")
	}

	// record various entries
	err = scribe.Record(ledger.Info, "test")
	if err != nil {
		t.Fatal("failed to record on info", err)
	}
	err = scribe.Record(ledger.Warn, "test")
	if err != nil {
		t.Fatal("failed to record on warn ", err)
	}
	err = scribe.Record(ledger.Error, "test")
	if err != nil {
		t.Fatal("failed to record on error ", err)
	}
	err = scribe.Record(ledger.Fatal, "test")
	if err != nil {
		t.Fatal("failed to record on fatal ", err)
	}

	// test entries
	if !strings.Contains(buffer.String(), "[INFO] | test") {
		t.Fatal("record didn't properly write on info")
	}
	if !strings.Contains(buffer.String(), "[WARN] | test") {
		t.Fatal("record didn't properly write on warn")
	}
	if !strings.Contains(buffer.String(), "[ERROR] | test") {
		t.Fatal("record didn't properly write on error")
	}
	if !strings.Contains(buffer.String(), "[FATAL] | test") {
		t.Fatal("record didn't properly write on fatal")
	}
}

func TestSetMinimumSeverity(t *testing.T) {
	// create byte buffer (implements io.Writer but into memory)
	buffer := bytes.NewBuffer([]byte(""))

	// create ledger
	scribe, err := ledger.New(buffer)
	if err != nil {
		t.Fatal("failed to create new ledger")
	}

	// set ledger minimum severity to Warn
	err = scribe.SetMinimumSeverity(ledger.Warn)
	if err != nil {
		t.Fatal("failed to set minimum severity to warn")
	}

	// attempt to record an event with a lower severity and one that should pass
	err = scribe.Record(ledger.Info, "test")
	if err != nil {
		t.Fatal("failed to record on info", err)
	}
	err = scribe.Record(ledger.Warn, "test")
	if err != nil {
		t.Fatal("failed to record on warn ", err)
	}

	// check if the event was written regardless and that the others still write
	if strings.Contains(buffer.String(), "[INFO] | test") {
		t.Fatal("record didn't properly write on info")
	}
	if !strings.Contains(buffer.String(), "[WARN] | test") {
		t.Fatal("record didn't properly write on warn")
	}
}

func TestInvalidSeverity(t *testing.T) {
	// create byte buffer (implements io.Writer but into memory)
	buffer := bytes.NewBuffer([]byte(""))

	// create ledger
	scribe, err := ledger.New(buffer)
	if err != nil {
		t.Fatal("failed to create new ledger")
	}

	// attempt to set an invalid minimum severity
	err = scribe.SetMinimumSeverity(ledger.Severity(99))
	if !errors.Is(err, ledger.ErrInvalidSeverity) {
		t.Fatal("failed to error invalid severity level when setting minimum severity")
	}

	// attempt to record an event with an invalid severity
	err = scribe.Record(ledger.Severity(99), "test")
	if !errors.Is(err, ledger.ErrInvalidSeverity) {
		t.Fatal("failed to error invalid severity level when recording")
	}
}

func TestNoWriter(t *testing.T) {
	// attempt to create ledger with a nil writer
	_, err := ledger.New(nil)
	if !errors.Is(err, ledger.ErrNoWriter) {
		t.Fatal("failed to error nil writer")
	}
}

// create a writer that writes only a certain number of bytes
type ShortWriter struct {
	MaxBytes int
	Buffer   *bytes.Buffer
}

// implement the Write() func to the short writer
func (sw *ShortWriter) Write(message []byte) (int, error) {
	maxPos := min(sw.MaxBytes, len(message))

	return sw.Buffer.Write(message[:maxPos])
}

func TestShortWrite(t *testing.T) {
	// create byte buffer (implements io.Writer but into memory)
	buffer := bytes.NewBuffer([]byte(""))

	// create writer that writes for 3 bytes then silently fails
	writer := &ShortWriter{MaxBytes: 3, Buffer: buffer}

	// create ledger
	scribe, err := ledger.New(writer)
	if err != nil {
		t.Fatal("failed to create new ledger")
	}

	err = scribe.Record(ledger.Info, "test, this is over 3 byte")
	if !errors.Is(err, ledger.ErrShortWrite) {
		t.Fatal("failed to error on short write")
	}
}

// create a writer that doesn't write and only errors
type ErrorWriter struct {
	ErrorToWrite string
}

// implement Write() for the error ErrorWriter
func (ew *ErrorWriter) Write(message []byte) (int, error) {
	return 0, errors.New(ew.ErrorToWrite)
}

func TestErrorWrite(t *testing.T) {
	// create writer that just errors on write
	writer := &ErrorWriter{"testErr"}

	// create ledger
	scribe, err := ledger.New(writer)
	if err != nil {
		t.Fatal("failed to create new ledger")
	}

	err = scribe.Record(ledger.Info, "test")
	if err == nil || err.Error() != errors.New("testErr").Error() {
		t.Fatal("failed to error on errored write")
	}
}
