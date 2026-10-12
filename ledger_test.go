package ledger_test

import (
	"bytes"
	"strings"
	"testing"

	ledger "github.com/Eiyarrr/Ledger"
)

func TestRecord(t *testing.T) {
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
