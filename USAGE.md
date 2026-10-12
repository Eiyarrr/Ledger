## Basic Example

# CODE
```
package main

import (
	"fmt"
	"os"

	// import Ledger
	ledger "github.com/Eiyarrr/Ledger"
)

func main() {
	// Create a writer that implements io.Writer
	writer := os.Stdout

	// Initialize a new Ledger
	scribe, err := ledger.New(writer)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/*
		// set the minimum severity to whatever you'd like (optional)
		err = scribe.SetMinimumSeverity(ledger.Warn)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	*/

	// some example records
	err = scribe.Record(ledger.Info, "Ledger setup complete!")       // typical program steps
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	err = scribe.Record(ledger.Warn, "This is a warning.")           // strange state reached
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	err = scribe.Record(ledger.Error, "And this is an Error!")       // problem reached
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	err = scribe.Record(ledger.Fatal, "Finish Him! (MK mentioned?)") // unrecoverable state
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// note: the program can still run after a fatal log.
	// Fatal simply means that the program can no longer recover and function as intended
}
```

# OUTPUT
```
  ~/dev/test/ledger ❯ go run .
[2026-10-12T12:48:31-07:00] | [INFO] | Ledger setup complete!
[2026-10-12T12:48:31-07:00] | [WARN] | This is a warning.
[2026-10-12T12:48:31-07:00] | [ERROR] | And this is an Error!
[2026-10-12T12:48:31-07:00] | [FATAL] | Finish Him! (MK mentioned?)
```
