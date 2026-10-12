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
	logger, err := ledger.New(buffer)
	if err != nil {
		t.Fatal("failed to create new ledger")
	}

	// record various entries
	err = logger.Record(ledger.Info, "test")
	if err != nil {
		t.Fatal("failed to record on info", err)
	}
	err = logger.Record(ledger.Warn, "test")
	if err != nil {
		t.Fatal("failed to record on warn ", err)
	}
	err = logger.Record(ledger.Error, "test")
	if err != nil {
		t.Fatal("failed to record on error ", err)
	}
	err = logger.Record(ledger.Fatal, "test")
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
	logger, err := ledger.New(buffer)
	if err != nil {
		t.Fatal("failed to create new ledger")
	}

	// set ledger minimum severity to Warn
	err = logger.SetMinimumSeverity(ledger.Warn)
	if err != nil {
		t.Fatal("failed to set minimum severity")
	}
	err = logger.Record(ledger.Warn, "test")
	if err != nil {
		t.Fatal("failed to record on warn ", err)
	}
	
	// attempt to record an event with a lower severity and one that should pass
	err = logger.Record(ledger.Info, "test")
	if err != nil {
		t.Fatal("failed to record on info", err)
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
	logger, err := ledger.New(buffer)
	if err != nil {
		t.Fatal("failed to create new ledger")
	}

	// attempt to set an invalid minimum severity
	err = logger.SetMinimumSeverity(ledger.Severity(99))
	if errors.Is(err, ledger.ErrInvalidSeverity) {
		t.Fatal("failed to error invalid severity level when setting minimum severity")
	}

	// attempt to record an event with an invalid severity
	err = logger.Record(ledger.Severity(99), "test")
	if errors.Is(err, ledger.ErrInvalidSeverity) {
		t.Fatal("failed to error invalid severity level when setting minimum severity")
	}
}
